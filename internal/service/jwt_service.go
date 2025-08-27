package service

import (
	"petshop-pos/pkg/exception"

	"github.com/google/uuid"
)

type JWTService interface {
    GenerateTokenPair(userID uuid.UUID, tenantName string, userRole string) (*TokenPair, *exception.Exception)
    ValidateAccessToken(tokenString string) (*Claims, *exception.Exception)
    ValidateRefreshToken(tokenString string) (*Claims, *exception.Exception)
    RefreshAccessToken(refreshToken string) (string, *exception.Exception)
}