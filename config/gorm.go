package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection(viperConfig *viper.Viper) (*gorm.DB, error) {
    user := viperConfig.GetString("DB_USER")
    pass := viperConfig.GetString("DB_PASSWORD")
    host := viperConfig.GetString("DB_HOST")
    port := viperConfig.GetString("DB_PORT")
    dbName := viperConfig.GetString("DB_NAME")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}