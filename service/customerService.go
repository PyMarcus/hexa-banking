package service

import (
	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/errs"
)

type ICustomerService interface {
	GetAllCustomers([]domain.Customer, *errs.AppError)
	GetCustomer(string)(*domain.Customer, *errs.AppError)
	GetActiveCustomers() ([]domain.Customer, *errs.AppError)
	GetInactiveCustomers() ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (d *DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll()
}

func (d *DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.ById(id)
}

func (d *DefaultCustomerService) GetActiveCustomers() ([]domain.Customer, *errs.AppError) {
	return d.repo.ByActiveStatus()
}

func (d *DefaultCustomerService) GetInactiveCustomers() ([]domain.Customer, *errs.AppError) {
	return d.repo.ByInactiveStatus()
}

func NewCustomerService(repositoryInject domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repositoryInject,
	}
}
