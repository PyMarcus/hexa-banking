package domain

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/PyMarcus/banking_auth/config"
	"github.com/PyMarcus/banking_auth/errs"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken(user *User) (*string, *errs.AppError) {
	claims := createClaims(user.Username)
	tokenHeaderAndClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenHeaderAndClaims.SignedString([]byte(config.GlobalConfig.JWT_SECRET_KEY))

	if err != nil {
		return nil, &errs.AppError{Code: http.StatusInternalServerError, Message: "Internal server errror"}
	}
	return &token, nil
}

func validateToken(token *string) bool {
	secret, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Fail")
		}
		return []byte(config.GlobalConfig.JWT_SECRET_KEY), nil
	})
	if err != nil || !secret.Valid {
		return false
	}
	claims, ok := secret.Claims.(jwt.MapClaims)
	log.Println(claims)
	if !ok {
		return false
	}
	// expire token?
	expToken := int64(claims["exp"].(float64))
	now := time.Now().Unix()
	return now <= expToken
}
