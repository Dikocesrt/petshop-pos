package config

import (
	"log/slog"
	"petshop-pos/internal/route"
	"petshop-pos/pkg/xvalidator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
    DB       *gorm.DB
    App      *gin.Engine
    Log      *slog.Logger
    Validate *xvalidator.Validator
    Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
    routeConfig := route.RouteConfig{
        App: config.App,
    }

	routeConfig.Setup()
}