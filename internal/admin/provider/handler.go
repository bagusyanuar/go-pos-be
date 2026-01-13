package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/handler"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type Handlers struct {
	Supplier          handler.SupplierHandler
	ProductCategory   handler.ProductCategoryHandler
	MaterialCategory  handler.MaterialCategoryHandler
	Unit              handler.UnitHandler
	Material          handler.MaterialHandler
	MaterialInventory handler.MaterialInventoryHandler
}

func NewHandlers(
	services *Services,
	config *config.AppConfig,
) *Handlers {
	return &Handlers{
		Supplier:          handler.NewSupplierHandler(services.Supplier, config),
		ProductCategory:   handler.NewProductCategoryHandler(services.ProductCategory, config),
		MaterialCategory:  handler.NewMaterialCategoryHandler(services.MaterialCategory, config),
		Unit:              handler.NewUnitHandler(services.Unit, config),
		Material:          handler.NewMaterialHandler(services.Material, config),
		MaterialInventory: handler.NewMaterialInventoryHandler(services.MaterialInventory, config),
	}
}
