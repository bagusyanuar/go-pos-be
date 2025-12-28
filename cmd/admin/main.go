package main

import (
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/admin"
	"github.com/bagusyanuar/go-pos-be/internal/shared/bootstrap"
)

func main() {
	cfg := bootstrap.NewAdminConfig()

	defer cfg.Logger.Sync()

	admin.Register(cfg)

	envPort := cfg.Viper.GetString("APP_ADMIN_PORT")
	if envPort == "" {
		envPort = "8080" // Default port jika config kosong
	}
	port := fmt.Sprintf(":%s", envPort)

	// 4. Jalankan Server
	fmt.Printf("🚀 Admin server is running on port %s\n", port)
	if err := cfg.App.Listen(port); err != nil {
		// Gunakan Logger untuk mencatat error sebelum panic
		cfg.Logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		panic(err)
	}
}
