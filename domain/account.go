package domain

import "github.com/PyMarcus/go_banking_api/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type IAccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
