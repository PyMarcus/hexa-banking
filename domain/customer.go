package domain

import "github.com/PyMarcus/go_banking_api/errs"

type Customer struct {
	Id           string
	Name         string
	City         string
	ZipCode      string
	BirthdayDate string
	Status       string
}

// interface[port]
type ICustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByActiveStatus() ([]Customer, *errs.AppError)
	ByInactiveStatus() ([]Customer, *errs.AppError)
}
