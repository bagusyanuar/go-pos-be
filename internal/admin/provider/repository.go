package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	ProductCategory   domain.ProductCategoryRepository
	MaterialCategory  domain.MaterialCategoryRepository
	Unit              domain.UnitRepository
	Material          domain.MaterialRepository
	MaterialInventory domain.MaterialInventoryRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductCategory:   repository.NewProductCategoryRepository(db),
		MaterialCategory:  repository.NewMaterialCategoryRepository(db),
		Unit:              repository.NewUnitRepository(db),
		Material:          repository.NewMaterialRepository(db),
		MaterialInventory: repository.NewMaterialInventoryRepository(db),
	}
}
