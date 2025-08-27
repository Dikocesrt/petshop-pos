package service

import (
	"petshop-pos/internal/dto"
	"petshop-pos/internal/repository"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/xvalidator"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
    userRepo   repository.UserRepository
    jwtService JWTService
    validator  *xvalidator.Validator
}

func NewAuthService(userRepo repository.UserRepository, jwtService JWTService, validator *xvalidator.Validator) AuthService {
    return &AuthServiceImpl{
        userRepo:   userRepo,
        jwtService: jwtService,
        validator:  validator,
    }
}

func (s *AuthServiceImpl) Login(req dto.LoginRequest) (*dto.LoginResponse, *exception.Exception) {
    // Validate request
    if err := s.validator.Struct(req); err != nil {
        return nil, exception.InvalidArgument("Invalid request")
    }

    // Find user by username and tenant
    user, err := s.userRepo.FindByUsernameAndTenant(req.Username, req.TenantName)
    if err != nil {
        return nil, err
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return nil, exception.Unauthenticated("Invalid credentials")
    }

    // Generate JWT tokens
    tokenPair, err := s.jwtService.GenerateTokenPair(user.ID, req.TenantName, string(user.Role))
    if err != nil {
        return nil, err
    }

    return &dto.LoginResponse{
        AccessToken:  tokenPair.AccessToken,
        RefreshToken: tokenPair.RefreshToken,
    }, nil
}

func (s *AuthServiceImpl) RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, *exception.Exception) {
    // Validate request
    if err := s.validator.Struct(req); err != nil {
        return nil, exception.InvalidArgument("Invalid request")
    }

    // Refresh access token
    accessToken, err := s.jwtService.RefreshAccessToken(req.RefreshToken)
    if err != nil {
        return nil, exception.Unauthenticated("invalid refresh token")
    }

    return &dto.RefreshTokenResponse{
        AccessToken: accessToken,
    }, nil
}