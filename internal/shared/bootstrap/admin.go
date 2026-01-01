package bootstrap

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"go.uber.org/zap"
)

func NewAdminConfig() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)
	baseLogger := config.NewLogger(viper)
	validator := config.NewValidator()
	jwt := config.NewJWTManager(viper)

	logger := baseLogger.With(zap.String("app", "admin-app"))
	return &config.AppConfig{
		App:       app,
		Viper:     viper,
		DB:        db,
		Logger:    logger,
		Validator: validator,
		JWT:       jwt,
	}
}
