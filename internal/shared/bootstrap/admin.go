package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func initialize() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)

	return &config.AppConfig{
		App:   app,
		Viper: viper,
		DB:    db,
	}
}

func StartAdminApp() {
	cfg := initialize()

	envPort := cfg.Viper.GetString("APP_ADMIN_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Admin server running on", port)
	if err := server.Listen(port); err != nil {
		panic("failed to start admin server")
	}
}
