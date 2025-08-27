package route

import (
	"petshop-pos/internal/handler"
	"petshop-pos/internal/middleware"
	"petshop-pos/internal/repository"

	"petshop-pos/internal/service"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
    App               *gin.Engine
    JWTService        service.JWTService
    AuthHandler       *handler.AuthHandler
    TenantRepository  repository.TenantRepository
}

func (c *RouteConfig) Setup() {
    c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
    api := c.App.Group("/api/v1")

    api.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "status": "OK",
            "message": "Service is healthy",
        })
    })

    // Auth routes
    auth := api.Group("/auth")
    {
        auth.POST("/login", c.AuthHandler.Login)
        auth.POST("/refresh", c.AuthHandler.RefreshToken)
    }

    // Protected routes
    protected := api.Group("")
    protected.Use(middleware.JWTMiddleware(c.JWTService))
    {
        protected.GET("/auth-health", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{
                "status": "OK",
                "message": "Service is healthy",
            })
        })
    }
}