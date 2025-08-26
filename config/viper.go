package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
    config := viper.New()
    config.SetConfigName(".env")
    config.SetConfigType("env")
    config.AddConfigPath("./")
    config.AddConfigPath("../")

    err := config.ReadInConfig()
    if err != nil {
        slog.Error("No .env file found, using environment variables", slog.Any("error", err))
    }
    config.AutomaticEnv()
    return config
}