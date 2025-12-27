package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func initializeFinanceApp() *config.AppConfig {
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

func StartFinanceApp() {
	cfg := initializeFinanceApp()

	envPort := cfg.Viper.GetString("APP_FINANCE_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Finance server running on", port)
	if err := server.Listen(port); err != nil {
		panic("failed to start finance server")
	}
}
