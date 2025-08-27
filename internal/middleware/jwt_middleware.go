package middleware

import (
	"net/http"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(jwtService service.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, response.NewBaseErrorResponse("Authorization header is required"))
            c.Abort()
            return
        }

        // Check if token starts with "Bearer "
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            c.JSON(http.StatusUnauthorized, response.NewBaseErrorResponse("Invalid authorization header format"))
            c.Abort()
            return
        }

        // Validate token
        claims, err := jwtService.ValidateAccessToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, response.NewBaseErrorResponse("Invalid token"))
            c.Abort()
            return
        }

        // Validate tenant from header matches token
        tenantName := c.GetHeader("x-tenant-name")
        if tenantName != claims.TenantName {
            c.JSON(http.StatusForbidden, response.NewBaseErrorResponse("Tenant mismatch"))
            c.Abort()
            return
        }

        // Set user info in context
        c.Set("user_id", claims.UserID)
        c.Set("tenant_name", claims.TenantName)
        c.Set("user_role", claims.UserRole)

        c.Next()
    }
}