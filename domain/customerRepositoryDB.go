package domain

import (
	"database/sql"

	"github.com/PyMarcus/go_banking_api/errs"
	"github.com/PyMarcus/go_banking_api/logger"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	connStr := "postgresql://postgres:your_password@localhost/postgres?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Error("Fail to connect into database err: " + err.Error())
		db.Close()
		panic(err)
	}

	return CustomerRepositoryDb{db: db}
}

func (crdb CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	db := crdb.db
	customers := make([]Customer, 0)

	selectCustomers := "SELECT * FROM customers;"
	err := db.Select(&customers, selectCustomers)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		logger.Error("Error while querying customer table: " + selectCustomers + err.Error())
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	return customers, nil
}

func (crdb CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	db := crdb.db
	var customer Customer

	selectCustomerById := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = $1;"

	err := db.Get(&customer, selectCustomerById, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		logger.Error("Error while scanning customers: " + err.Error())
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}
	return &customer, nil
}

func (crdb CustomerRepositoryDb) ByActiveStatus() ([]Customer, *errs.AppError) {
	db := crdb.db
	customers := make([]Customer, 0)

	selectCustomerByStatus := "SELECT * FROM customers WHERE status = $1;"

	err := db.Select(&customers, selectCustomerByStatus, 1)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		logger.Error("Error while querying customer table: " + selectCustomerByStatus + err.Error())
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	return customers, nil
}

func (crdb CustomerRepositoryDb) ByInactiveStatus() ([]Customer, *errs.AppError) {
	db := crdb.db
	customers := make([]Customer, 0)

	selectCustomerByStatus := "SELECT * FROM customers WHERE status = $1;"

	err := db.Select(&customers, selectCustomerByStatus, 0)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		logger.Error("Error while querying customer table: " + selectCustomerByStatus + err.Error())
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	return customers, nil
}
