package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/auth/domain"
	"github.com/bagusyanuar/go-pos-be/internal/auth/service"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Services struct {
	Auth domain.AuthService
}

func NewServices(
	repos *Repositories,
	config *config.AppConfig,
) *Services {
	return &Services{
		Auth: service.NewAuthService(repos.User, config),
	}
}
