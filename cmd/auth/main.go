package main

import (
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/auth"
	"github.com/bagusyanuar/go-pos-be/internal/shared/bootstrap"
)

func main() {
	cfg := bootstrap.NewAuthConfig()

	defer cfg.Logger.Sync()

	auth.Register(cfg)

	envPort := cfg.Viper.GetString("APP_AUTH_PORT")
	if envPort == "" {
		envPort = "8000"
	}
	port := fmt.Sprintf(":%s", envPort)
	fmt.Printf("🚀 Auth server is running on port %s\n", port)
	if err := cfg.App.Listen(port); err != nil {
		cfg.Logger.Error(fmt.Sprintf("Failed to start auth server: %v", err))
		panic(err)
	}
}
