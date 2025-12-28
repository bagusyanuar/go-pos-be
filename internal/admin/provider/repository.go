package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	ProductCategory repository.ProductCategoryRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductCategory: repository.NewProductCategoryRepository(db),
	}
}
