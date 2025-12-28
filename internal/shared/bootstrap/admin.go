package bootstrap

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func NewAdminConfig() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)
	logger := config.NewLogger(viper)
	validator := config.NewValidator()

	return &config.AppConfig{
		App:       app,
		Viper:     viper,
		DB:        db,
		Logger:    logger,
		Validator: validator,
	}
}
