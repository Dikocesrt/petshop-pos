package config

import (
	"io"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

func InitSlog(viperConfig *viper.Viper) *slog.Logger {
    logPath := viper.GetString("LOG_PATH")
    if logPath == "" {
        logPath = "./logs"
    }
    debug := viper.GetBool("DEBUG")

    if err := os.MkdirAll(logPath, 0755); err != nil {
        panic(err)
    }
    file, err := os.OpenFile(logPath+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        panic(err)
    }

    var logLevel slog.Level
    if debug {
        logLevel = slog.LevelDebug
    } else {
        logLevel = slog.LevelInfo
    }

    logJSONHandler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, file), &slog.HandlerOptions{
        AddSource: true,
        Level:     logLevel,
    })
    logger := slog.New(logJSONHandler)
    slog.SetDefault(logger)
    return logger
}