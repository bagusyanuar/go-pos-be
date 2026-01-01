package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	App       *fiber.App
	Viper     *viper.Viper
	DB        *gorm.DB
	Logger    *zap.Logger
	Validator *validator.Validate
	JWT       *JWTConfig
}
