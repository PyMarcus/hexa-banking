package service

import (
	"github.com/PyMarcus/banking_auth/domain"
	"github.com/PyMarcus/banking_auth/dto"
	"github.com/PyMarcus/banking_auth/errs"
)

type IAuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	ValidateToken(*string) bool
}

type AuthService struct{
	repo domain.IAuthRepository
}

func (as AuthService) Login(request dto.LoginRequest) (*dto.LoginResponse, *errs.AppError){
	user, err := as.repo.GetUser(request.Username, request.Password)
	if err != nil{
		return nil, err 
	}
	token, err := as.repo.GenerateToken(user)
	if err != nil{
		return nil, err 
	}
	response := &dto.LoginResponse{Token: *token}
	return response, nil
}

func (as AuthService) ValidateToken(token *string) bool{
	return as.repo.ValidateToken(token)
}

func NewAuthService(repo domain.IAuthRepository) IAuthService{
	return AuthService{
		repo: repo,
	}
}