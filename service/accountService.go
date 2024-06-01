package service

import (
	"time"

	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
)

// port
type IAccountService interface {
	Save(domain.Account) (*dto.AccountResponse, *errs.AppError)
}

// implementation
type DefaultAccountService struct {
	repo domain.IAccountRepository
}

func (das DefaultAccountService) NewAccount(request dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	if err := request.Validate() ; err != nil {
		return nil, err
	}
	acc := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}

	created, err := das.repo.Save(acc)
	if err != nil {
		return nil, err
	}
	return &dto.AccountResponse{AccountId: created.AccountId}, nil
}

func NewAccountService(repositoryInject domain.IAccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repositoryInject,
	}
}
