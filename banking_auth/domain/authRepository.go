package domain

import (
	"net/http"
	"database/sql"
	"github.com/PyMarcus/banking_auth/errs"
	"github.com/jmoiron/sqlx"
)

type IAuthRepository interface {
	GetUser(string, string) (*User, *errs.AppError)
	GenerateToken(*User) (*string, *errs.AppError)
	ValidateToken(*string) bool
}

type AuthRepository struct{
	conn *sqlx.DB
}

func (ar AuthRepository) GetUser(username string, password string) (*User, *errs.AppError){
	tx, err := ar.conn.Begin()
	if err != nil {
		return nil, &errs.AppError{Code: http.StatusInternalServerError, Message: "Unexpected server error"}
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var user User
	row := tx.QueryRow("SELECT username, password FROM users WHERE username = $1 and password = $2 limit 1;", username, password)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, &errs.AppError{Code: http.StatusNotFound, Message: "User not found"}
		}
		return nil, &errs.AppError{Code: http.StatusInternalServerError, Message: "Unexpected server error"}
	}

	return &user, nil
}

func (ar AuthRepository) GenerateToken(user *User) (*string, *errs.AppError){
	return generateToken(user)
}

func (ar AuthRepository) ValidateToken(token *string) bool{
	return validateToken(token)
}

func NewAuthRepository(db *sqlx.DB) IAuthRepository{
	return &AuthRepository{
		conn:db,
	}
}