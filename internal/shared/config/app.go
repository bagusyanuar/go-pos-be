package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	App   *fiber.App
	Viper *viper.Viper
	DB    *gorm.DB
}
