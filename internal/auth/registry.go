package auth

import (
	"github.com/bagusyanuar/go-pos-be/internal/auth/provider"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func Register(
	config *config.AppConfig,
) {
	repositories := provider.NewRepositories(config.DB)
	services := provider.NewServices(repositories, config)
	handlers := provider.NewHandlers(services, config)

	RegisterRoutes(config, handlers)
}
