package domain

import (
	"strconv"

	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	client *sqlx.DB
}

// Save an account to the customer
func (ar *AccountRepository) Save(acc Account) (*Account, *errs.AppError) {
    sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING account_id"

    // Start transaction
    tctx, err := ar.client.Begin()
    
    if err != nil {
        return nil, &errs.AppError{Message: "Failed to start transaction", Code: 500}
    }
    defer func() {
        if err != nil {
            tctx.Rollback()
        }
    }()

    var id int64
    err = tctx.QueryRow(sqlInsert, acc.CustomerId, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status).Scan(&id)
    if err != nil {
        return nil, &errs.AppError{Message: "Failed to insert account ", Code: 500}
    }
    acc.AccountId = strconv.FormatInt(id, 10)

    // Commit transaction
    if err := tctx.Commit(); err != nil {
        return nil, &errs.AppError{Message: "Failed to commit transaction", Code: 500}
    }
    
    return &acc, nil
}

func NewAccountRepository(dbClient *sqlx.DB) IAccountRepository{
	return &AccountRepository{
		client: dbClient,
	}
}
