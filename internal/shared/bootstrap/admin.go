package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func initializeAdminApp() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)
	logger := config.NewLogger(viper)
	defer logger.Sync()
	validator := config.NewValidator()

	return &config.AppConfig{
		App:       app,
		Viper:     viper,
		DB:        db,
		Logger:    logger,
		Validator: validator,
	}
}

func StartAdminApp() {
	cfg := initializeAdminApp()

	envPort := cfg.Viper.GetString("APP_ADMIN_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Admin server running on", port)
	if err := server.Listen(port); err != nil {
		panic("failed to start admin server")
	}
}
