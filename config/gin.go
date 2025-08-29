package config

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGin(viperConfig *viper.Viper, log *slog.Logger) *gin.Engine {
    allowOrigins := viperConfig.GetString("ALLOW_ORIGINS")
    allowMethods := viperConfig.GetString("ALLOW_METHODS")
    allowHeaders := viperConfig.GetString("ALLOW_HEADERS")

    log.Info("Raw CORS Config from env", "origins", allowOrigins, "methods", allowMethods, "headers", allowHeaders)

    gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Recovery())
    
    // Custom CORS middleware yang lebih eksplisit
    r.Use(func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        
        // Set CORS headers
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-CSRF-Token")
        c.Header("Access-Control-Expose-Headers", "Content-Length")
        c.Header("Access-Control-Allow-Credentials", "false")
        c.Header("Access-Control-Max-Age", "43200") // 12 hours
        
        // Handle preflight requests
        if c.Request.Method == "OPTIONS" {
            log.Info("OPTIONS request received", "origin", origin, "path", c.Request.URL.Path)
            c.AbortWithStatus(204)
            return
        }
        
        log.Info("Request received", "method", c.Request.Method, "origin", origin, "path", c.Request.URL.Path)
        c.Next()
    })

    return r
}

func splitAndTrim(s string) []string {
    parts := strings.Split(s, ",")
    result := make([]string, 0, len(parts))
    for _, part := range parts {
        trimmed := strings.TrimSpace(part)
        if trimmed != "" {
            result = append(result, trimmed)
        }
    }
    return result
}