package service

import (
	"petshop-pos/internal/dto"
	"petshop-pos/pkg/exception"
)

type AuthService interface {
    Login(req dto.LoginRequest) (*dto.LoginResponse, *exception.Exception)
    RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, *exception.Exception)
}