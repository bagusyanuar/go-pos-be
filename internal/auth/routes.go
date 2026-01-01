package auth

import (
	"github.com/bagusyanuar/go-pos-be/internal/auth/provider"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func RegisterRoutes(
	config *config.AppConfig,
	handlers *provider.Handlers,
) {
	app := config.App
	app.Post("/login", handlers.Auth.Login)
}
