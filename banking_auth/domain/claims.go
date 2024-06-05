package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// createClaims with 24 hours token expiration
func createClaims(username string) jwt.MapClaims {
	now := time.Now()
	exp := now.Add(time.Hour * 24).Unix()
	return jwt.MapClaims{
		"sub":   username,
		"iss":   "withoutdomain",
		"uid":   username,
		"roles": []string{"user"},
		"exp": exp,
	}
}
