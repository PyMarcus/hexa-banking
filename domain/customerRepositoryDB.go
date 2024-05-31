package domain

import (
	"database/sql"
	"log"

	"github.com/PyMarcus/go_banking_api/errs"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	connStr := "postgresql://postgres:your_password@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic("Fail to connect into database err: ", err)
		db.Close()
		return CustomerRepositoryDb{}
	}

	return CustomerRepositoryDb{db: db}
}

func (crdb CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	db := crdb.db

	selectCustomers := "SELECT * FROM customers;"
	rows, err := db.Query(selectCustomers)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		log.Println("Error while querying customer table:", selectCustomers, err)
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.BirthdayDate,
			&customer.City,
			&customer.ZipCode,
			&customer.Status,
		); err != nil {
			log.Println("Error while scanning customers:", err)
			return nil, errs.UnexpectedDatabaseError("Unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (crdb CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	db := crdb.db
	selectCustomerById := "SELECT * FROM customers WHERE customer_id = $1;"

	row := db.QueryRow(selectCustomerById, id)

	var customer Customer
	if err := row.Scan(
		&customer.Id,
		&customer.Name,
		&customer.BirthdayDate,
		&customer.City,
		&customer.ZipCode,
		&customer.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		log.Println("Error while scanning customers:", err)
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}
	return &customer, nil
}

func (crdb CustomerRepositoryDb) ByActiveStatus() ([]Customer, *errs.AppError) {
	db := crdb.db
	selectCustomerByStatus := "SELECT * FROM customers WHERE status = $1;"

	rows, err := db.Query(selectCustomerByStatus, 1)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		log.Println("Error while querying customer table:", selectCustomerByStatus, err)
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.BirthdayDate,
			&customer.City,
			&customer.ZipCode,
			&customer.Status,
		); err != nil {
			log.Println("Error while scanning customers:", err)
			return nil, errs.UnexpectedDatabaseError("Unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (crdb CustomerRepositoryDb) ByInactiveStatus() ([]Customer, *errs.AppError) {
	db := crdb.db
	selectCustomerByStatus := "SELECT * FROM customers WHERE status = $1;"

	rows, err := db.Query(selectCustomerByStatus, 0)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customers not found")
		}
		log.Println("Error while querying customer table:", selectCustomerByStatus, err)
		return nil, errs.UnexpectedDatabaseError("Unexpected database error")
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.BirthdayDate,
			&customer.City,
			&customer.ZipCode,
			&customer.Status,
		); err != nil {
			log.Println("Error while scanning customers:", err)
			return nil, errs.UnexpectedDatabaseError("Unexpected database error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}