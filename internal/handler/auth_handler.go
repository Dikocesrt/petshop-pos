package handler

import (
	"net/http"
	"petshop-pos/internal/dto"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req dto.LoginRequest

    // bind request body
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
        return
    }

    // Get tenant name from header
    tenantName := c.GetHeader("x-tenant-name")
    if tenantName == "" {
        c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Tenant name is required"))
        return
    }
    req.TenantName = tenantName

    // Call auth service
    resp, err := h.authService.Login(req)
    if err != nil {
        code, resp := response.MapExceptionToHTTP(err)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Login successful", resp))
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var req dto.RefreshTokenRequest
    // Bind request body
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
        return
    }

    // Call auth service
    resp, err := h.authService.RefreshToken(req)
    if err != nil {
        code, resp := response.MapExceptionToHTTP(err)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Token refreshed successfully", resp))
}