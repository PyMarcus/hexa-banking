package errs

import (
	"net/http"
)

type AppError struct {
	Code    int 	`json:"code"`
	Message string	`json:"message"`
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func UnexpectedDatabaseError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
