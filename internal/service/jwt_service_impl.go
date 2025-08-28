package service

import (
	"errors"
	"petshop-pos/pkg/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceImpl struct {
    accessSecret  string
    refreshSecret string
    accessTTL     time.Duration
    refreshTTL    time.Duration
}

type TokenPair struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

type Claims struct {
    UserID     string `json:"user_id"`
    TenantName string `json:"tenant_name"`
    UserRole   string `json:"user_role"`
    jwt.RegisteredClaims
}

func NewJWTService(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) JWTService {
    return &JWTServiceImpl{
        accessSecret:  accessSecret,
        refreshSecret: refreshSecret,
        accessTTL:     accessTTL,
        refreshTTL:    refreshTTL,
    }
}

func (j *JWTServiceImpl) GenerateTokenPair(userID string, tenantName string, userRole string) (*TokenPair, *exception.Exception) {
    // Generate Access Token Claims
    accessClaims := &Claims{
        UserID:     userID,
        TenantName: tenantName,
        UserRole:   userRole,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTTL)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    // Sign the Access Token
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessTokenString, err := accessToken.SignedString([]byte(j.accessSecret))
    if err != nil {
        return nil, exception.Internal("Failed to generate access token", err)
    }

    // Generate Refresh Token Claims
    refreshClaims := &Claims{
        UserID:     userID,
        TenantName: tenantName,
        UserRole:   userRole,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTTL)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    // Sign the Refresh Token
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshTokenString, err := refreshToken.SignedString([]byte(j.refreshSecret))
    if err != nil {
        return nil, exception.Internal("Failed to generate refresh token", err)
    }

    return &TokenPair{
        AccessToken:  accessTokenString,
        RefreshToken: refreshTokenString,
    }, nil
}

func (j *JWTServiceImpl) ValidateAccessToken(tokenString string) (*Claims, *exception.Exception) {
    claims, err := j.validateToken(tokenString, j.accessSecret)
    if err != nil {
        return nil, exception.Unauthenticated("Invalid access token")
    }
    return claims, nil
}

func (j *JWTServiceImpl) ValidateRefreshToken(tokenString string) (*Claims, *exception.Exception) {
    claims, err := j.validateToken(tokenString, j.refreshSecret)
    if err != nil {
        return nil, exception.Unauthenticated("Invalid refresh token")
    }
    return claims, nil
}

func (j *JWTServiceImpl) validateToken(tokenString, secret string) (*Claims, error) {
    // Parse the JWT token
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    // Validate the token claims
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}

func (j *JWTServiceImpl) RefreshAccessToken(refreshToken string) (string, *exception.Exception) {
    // Validate the refresh token
    claims, err := j.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", exception.Unauthenticated("Invalid refresh token")
    }

    // Generate Access Token Claims
    accessClaims := &Claims{
        UserID:     claims.UserID,
        TenantName: claims.TenantName,
        UserRole:   claims.UserRole,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTTL)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    // Sign the Access Token
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessTokenString, signErr := accessToken.SignedString([]byte(j.accessSecret))
    if signErr != nil {
        return "", exception.Internal("Failed to generate access token", signErr)
    }

    return accessTokenString, nil
}