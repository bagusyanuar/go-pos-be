package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/service"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Services struct {
	ProductCategory   domain.ProductCategoryService
	MaterialCategory  domain.MaterialCategoryService
	Unit              domain.UnitService
	Material          domain.MaterialService
	MaterialInventory domain.MaterialInventoryService
}

func NewServices(
	repos *Repositories,
	config *config.AppConfig,
) *Services {
	return &Services{
		ProductCategory:   service.NewProductCategoryService(repos.ProductCategory, config),
		MaterialCategory:  service.NewMaterialCategoryService(repos.MaterialCategory, config),
		Unit:              service.NewUnitService(repos.Unit, config),
		Material:          service.NewMaterialService(repos.Material, config),
		MaterialInventory: service.NewMaterialInventoryService(repos.MaterialInventory, config),
	}
}
