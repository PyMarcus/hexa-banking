package service

import (
	"github.com/PyMarcus/go_banking_api/domain"
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
)

type ICustomerService interface {
	GetAllCustomers([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string)(*dto.CustomerResponse, *errs.AppError)
	GetActiveCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetInactiveCustomers() ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (d *DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	var customerDTO []dto.CustomerResponse
	customer, err := d.repo.FindAll()
	if err != nil{
		return nil, err 
	}
	for _, c := range customer{
		customerDTO = append(customerDTO, *c.ToDTO())
	}
	return customerDTO, nil 
}

func (d *DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := d.repo.ById(id)
	if err != nil{
		return nil, err 
	}
	return customer.ToDTO(), nil
}

func (d *DefaultCustomerService) GetActiveCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	var customerDTOActive []dto.CustomerResponse
	customer, err := d.repo.ByActiveStatus()
	if err != nil{
		return nil, err 
	}
	for _, c := range customer{
		customerDTOActive = append(customerDTOActive, *c.ToDTO())
	}
	return customerDTOActive, nil 
}

func (d *DefaultCustomerService) GetInactiveCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	var customerDTOInactive []dto.CustomerResponse
	customer, err := d.repo.ByInactiveStatus()
	if err != nil{
		return nil, err 
	}
	for _, c := range customer{
		customerDTOInactive = append(customerDTOInactive, *c.ToDTO())
	}
	return customerDTOInactive, nil 
}

func NewCustomerService(repositoryInject domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repositoryInject,
	}
}
