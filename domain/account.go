package domain

import "github.com/PyMarcus/go_banking_api/errs"

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string `db:"status"`
}

type AccountTransaction struct {
	AccountId   string
	Type 		string
	Value 		float64
}

type IAccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
