package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/PyMarcus/go_banking_api/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	client *sqlx.DB
}

func (tr TransactionRepository) Transfer(origin Account, destiny AccountTransaction) (*Transaction, *errs.AppError) {
	transaction, err := tr.makeTransaction(origin, destiny)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (tr *TransactionRepository) GetOriginAccount(customerId string, accountId string) (*Account, *errs.AppError) {
	selectSql := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = $1;"
	origin := Account{}
	log.Println(accountId, customerId)
	err := tr.client.Get(&origin, selectSql, accountId)
	if err != nil {
		logger.Error(err.Error())
		return nil, &errs.AppError{Message: "Unexpected database error", Code: 500}
	}
	log.Println(origin)
	return &origin, nil
}

func (tr TransactionRepository) makeTransaction(origin Account, destiny AccountTransaction) (*Transaction, *errs.AppError) {
	tctx, err := tr.client.Begin()
	if err != nil {
		return nil, &errs.AppError{Message: "Failed to start transaction", Code: 500}
	}
	defer tctx.Rollback()

	creditFromOriginSql := "UPDATE accounts SET amount = amount - $1 WHERE account_id = $2;"
	_, err = tctx.Exec(creditFromOriginSql, destiny.Value, origin.AccountId)
	if err != nil {
		return nil, &errs.AppError{Message: "1)Failed to insert into account ", Code: 500}
	}

	depositToDestinySql := "UPDATE accounts SET amount = amount + $1 WHERE account_id = $2"
	_, err = tctx.Exec(depositToDestinySql, destiny.Value, destiny.AccountId)
	if err != nil {
		return nil, &errs.AppError{Message: "2)Failed to insert into account ", Code: 500}
	}

	transaction, e := tr.doTransaction(&origin, &destiny, tctx)
	if e != nil {
		return nil, &errs.AppError{Message: e.Error(), Code: 500}
	}
	
	if err := tctx.Commit(); err != nil {
		return nil, &errs.AppError{Message: "3)Failed to commit transaction", Code: 500}
	}

	return transaction, nil
}

func (tr *TransactionRepository) doTransaction(origin *Account, destiny *AccountTransaction, tctx *sql.Tx) (*Transaction, error) {
	insertTransactionsql := "INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4) RETURNING transaction_id;"
	t := &Transaction{
		AccountId:       origin.AccountId,
		Amount:          destiny.Value,
		TransactionType: destiny.Type,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	var id string
	err := tctx.QueryRow(insertTransactionsql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&id)
	if err != nil {
		return nil, err
	}
	t.TransactionId = id

	return t, nil
}

func NewTransactionRepository(dbClient *sqlx.DB) TransactionRepository {
	return TransactionRepository{
		client: dbClient,
	}
}
