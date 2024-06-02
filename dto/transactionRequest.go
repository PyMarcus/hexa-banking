package dto

import (
	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/PyMarcus/go_banking_api/strategy"
)

type TransactionRequest struct {
	AccountId		  string  `json:"account_id"`
	DestinyAccountId  string  `json:"destiny_account_id"`
	TransactionType   string  `json:"transaction_type"`
	Amount            float64 `json:"amount"`
}

func (tr TransactionRequest) Validate() *errs.AppError {
	s := strategy.TransactionStrategy{}
	return s.Validate(tr.Amount)
}