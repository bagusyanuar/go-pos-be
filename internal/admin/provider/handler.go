package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/handler"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Handlers struct {
	ProductCategory handler.ProductCategoryHandler
}

func NewHandlers(
	services *Services,
	config *config.AppConfig,
) *Handlers {
	return &Handlers{
		ProductCategory: *handler.NewProductCategoryHandler(services.ProductCategory, config),
	}
}
