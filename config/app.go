package config

import (
	"log/slog"
	"petshop-pos/internal/handler"
	"petshop-pos/internal/repository"
	"petshop-pos/internal/route"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/xvalidator"
	"time"

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
    jwtService := service.NewJWTService(config.Config.GetString("ACCESS_JWT_SECRET"), config.Config.GetString("REFRESH_JWT_SECRET"), time.Hour*24, time.Hour*24*7)

    userRepo := repository.NewUserRepository(config.DB)

    tenantRepo := repository.NewTenantRepository(config.DB)

    authService := service.NewAuthService(userRepo, jwtService, config.Validate)
    authHandler := handler.NewAuthHandler(authService)

    routeConfig := route.RouteConfig{
        App:        config.App,
        JWTService: jwtService,
        AuthHandler: authHandler,
        TenantRepository: tenantRepo, //temporary
    }

	routeConfig.Setup()
}