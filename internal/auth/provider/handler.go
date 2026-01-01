package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/auth/handler"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Handlers struct {
	Auth handler.AuthHandler
}

func NewHandlers(
	services *Services,
	config *config.AppConfig,
) *Handlers {
	return &Handlers{
		Auth: handler.NewAuthHandler(services.Auth, config),
	}
}
