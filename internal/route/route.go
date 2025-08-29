package route

import (
	"petshop-pos/internal/handler"
	"petshop-pos/internal/middleware"

	"petshop-pos/internal/service"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
    App               *gin.Engine
    JWTService        service.JWTService
    AuthHandler       *handler.AuthHandler
    BrandHandler      *handler.BrandHandler
    CategoryHandler   *handler.CategoryHandler
    ProductHandler    *handler.ProductHandler
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

    api.GET("/cors-test", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "CORS is working",
        "origin":  c.Request.Header.Get("Origin"),
        "method":  c.Request.Method,
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
        // Brand routes
        brand := protected.Group("/brands")
        {
            brand.POST("/", c.BrandHandler.CreateBrand)
            brand.GET("/:id", c.BrandHandler.GetByID)
            brand.GET("/", c.BrandHandler.GetAll)
            brand.PUT("/:id", c.BrandHandler.Update)
            brand.DELETE("/:id", c.BrandHandler.Delete)
        }

        // Category routes
        category := protected.Group("/categories")
        {
            category.POST("/", c.CategoryHandler.CreateCategory)
            category.GET("/:id", c.CategoryHandler.GetByID)
            category.GET("/", c.CategoryHandler.GetAll)
            category.PUT("/:id", c.CategoryHandler.Update)
            category.DELETE("/:id", c.CategoryHandler.Delete)
        }

        // Product routes
        product := protected.Group("/products")
        {
            product.POST("/", c.ProductHandler.Create)
            product.GET("/:id", c.ProductHandler.GetByID)
            product.GET("/", c.ProductHandler.GetAll)
            product.PUT("/:id", c.ProductHandler.Update)
            product.DELETE("/:id", c.ProductHandler.Delete)
        }
    }
}