package domain

import (
	"github.com/PyMarcus/go_banking_api/dto"
	"github.com/PyMarcus/go_banking_api/errs"
)

type Customer struct {
	Id           string `db:"customer_id"`
	Name         string `db:"name"`
	City         string `db:"city"`
	ZipCode      string `db:"zipcode"`
	BirthdayDate string `db:"date_of_birth"`
	Status       string `db:"status"`
}

func (customer Customer) ToDTO() *dto.CustomerResponse{
	customer.statusAsTxt()
	return &dto.CustomerResponse{
		Id: customer.Id,
		Name: customer.Name,
		ZipCode: customer.ZipCode,
		BirthdayDate: customer.BirthdayDate,
		Status: customer.Status,
		City: customer.City,
	}
}

func (customer *Customer) statusAsTxt(){
	if customer.Status == "false" {
		customer.Status = "inactive"
	}else{
		customer.Status = "active"
	}
}

// interface[port]
type ICustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByActiveStatus() ([]Customer, *errs.AppError)
	ByInactiveStatus() ([]Customer, *errs.AppError)
}
