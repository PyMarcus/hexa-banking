package service

import (
	"fmt"
	"net/http"

	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
)

// port
type ITransactionService interface {
	DoTransaction(dto.TransactionRequest, string) (*dto.TransactionResponse, *errs.AppError)
}

// implementation
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (dts DefaultTransactionService) DoTransaction(request dto.TransactionRequest, customerId string) (*dto.TransactionResponse, *errs.AppError){
	err := request.Validate()
	if err != nil{
		return nil, err 
	}
	account, err := dts.repo.GetOriginAccount(customerId, request.AccountId)
	
	if err != nil{
		return nil, err 
	}
	
	if account.Amount < request.Amount{
		return nil, &errs.AppError{Message: fmt.Sprintf("Insufficient funds: $%f", account.Amount), Code: http.StatusBadRequest}
	}
	destiny := &domain.AccountTransaction{AccountId: request.DestinyAccountId, Type: "Credit", Value: request.Amount}

	t, err := dts.repo.Transfer(*account, *destiny)
	if err != nil{
		return nil, err 
	}
	
	tdto := t.ToDTO()
	return tdto, nil 
}


func NewTransactionService(repositoryInject domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		repo: repositoryInject,
	}
}