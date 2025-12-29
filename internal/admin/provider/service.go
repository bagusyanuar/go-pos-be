package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/service"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Services struct {
	ProductCategory domain.ProductCategoryService
}

func NewServices(
	repos *Repositories,
	config *config.AppConfig,
) *Services {
	return &Services{
		ProductCategory: service.NewProductCategoryService(repos.ProductCategory, config),
	}
}
