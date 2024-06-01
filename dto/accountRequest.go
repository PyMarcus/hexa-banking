package dto

import (
	"github.com/PyMarcus/go_banking_api/strategy"
	"github.com/PyMarcus/go_banking_api/errs"
)

type AccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (ar AccountRequest) Validate() *errs.AppError {
	var s strategy.IValidationStrategy

	switch ar.AccountType {
    case "saving":
        s = strategy.SavingAccountStrategy{}
    case "checking":
        s = strategy.CheckingAccountStrategy{}
    default:
        return errs.NewValidationError("Invalid account type")
    }
	return s.Validate(ar.Amount)
}
