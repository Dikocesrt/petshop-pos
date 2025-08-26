package config

import (
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGin(viperConfig *viper.Viper, log *slog.Logger) *gin.Engine {
    allowOrigins := viperConfig.GetString("ALLOW_ORIGINS")
    allowMethods := viperConfig.GetString("ALLOW_METHODS")
    allowHeaders := viperConfig.GetString("ALLOW_HEADERS")

    origins := []string{}
    if allowOrigins != "" {
        origins = append(origins, splitAndTrim(allowOrigins)...)
    }
    methods := []string{}
    if allowMethods != "" {
        methods = append(methods, splitAndTrim(allowMethods)...)
    }
    headers := []string{}
    if allowHeaders != "" {
        headers = append(headers, splitAndTrim(allowHeaders)...)
    }

    log.Info("CORS Config", "origins", origins, "methods", methods, "headers", headers)

    gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Recovery())
    r.Use(cors.New(cors.Config{
        AllowOrigins: origins,
        AllowMethods: methods,
        AllowHeaders: headers,
    }))
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