package domain

import (
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type ITransactionRepository interface {
	Transfer(Account, AccountTransaction) (*Transaction, *errs.AppError)
	GetOriginAccount(string, string) (*Account, *errs.AppError)
}

func (t Transaction) ToDTO() *dto.TransactionResponse{
	return &dto.TransactionResponse{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		TransactionValue: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
