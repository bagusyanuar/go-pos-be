package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	Supplier          domain.SupplierRepository
	ProductCategory   domain.ProductCategoryRepository
	MaterialCategory  domain.MaterialCategoryRepository
	Unit              domain.UnitRepository
	Material          domain.MaterialRepository
	MaterialInventory domain.MaterialInventoryRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Supplier:          repositories.NewSupplierRepository(db),
		ProductCategory:   repositories.NewProductCategoryRepository(db),
		MaterialCategory:  repositories.NewMaterialCategoryRepository(db),
		Unit:              repositories.NewUnitRepository(db),
		Material:          repositories.NewMaterialRepository(db),
		MaterialInventory: repositories.NewMaterialInventoryRepository(db),
	}
}
