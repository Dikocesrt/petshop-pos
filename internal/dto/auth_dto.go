package dto

type LoginRequest struct {
    Username   string `json:"username" validate:"required"`
    Password   string `json:"password" validate:"required"`
    TenantName string `json:"-"` // Will be set from header
}

type LoginResponse struct {
    AccessToken  string       `json:"accessToken"`
    RefreshToken string       `json:"refreshToken"`
}

type RefreshTokenRequest struct {
    RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenResponse struct {
    AccessToken string `json:"accessToken"`
}