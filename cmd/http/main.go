package main

import (
	"fmt"
	"petshop-pos/config"
	"petshop-pos/pkg/xvalidator"
)

func main() {
    viperConfig := config.InitViper()
    log := config.InitSlog(viperConfig)
    db, err := config.NewGormConnection(viperConfig)
    if err != nil {
        log.Error("Failed to connect to DB", "error", err)
        return
    }
    validate := xvalidator.NewValidator()
    app := config.InitGin(viperConfig, log)

    config.Bootstrap(&config.BootstrapConfig{
        DB:       db,
        App:      app,
        Log:      log,
        Validate: validate,
        Config:   viperConfig,
    })

    port := viperConfig.GetInt("PORT")
    if port == 0 {
        port = 8000
    }
    addr := fmt.Sprintf(":%d", port)
    log.Info("Starting server", "addr", addr)
    if err := app.Run(addr); err != nil {
        log.Error("Failed to start server", "error", err)
    }
}