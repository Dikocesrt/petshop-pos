package route

import (
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
    App               *gin.Engine
}

func (c *RouteConfig) Setup() {
    c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
    guest := c.App.Group("/api/v1")
    guest.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "status": "OK",
            "message": "Service is healthy",
        })
    })
}