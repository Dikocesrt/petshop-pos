package service

import (
	"petshop-pos/pkg/exception"
)

type JWTService interface {
    GenerateTokenPair(userID string, tenantName string, userRole string) (*TokenPair, *exception.Exception)
    ValidateAccessToken(tokenString string) (*Claims, *exception.Exception)
    ValidateRefreshToken(tokenString string) (*Claims, *exception.Exception)
    RefreshAccessToken(refreshToken string) (string, *exception.Exception)
}