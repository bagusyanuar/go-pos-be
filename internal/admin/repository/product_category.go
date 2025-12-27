package repository

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"gorm.io/gorm"
)

type (
	ProductCategoryRepository interface {
		FindAll(ctx context.Context) ([]entity.ProductCategory, error)
		FindByID(ctx context.Context, id string) (*entity.ProductCategory, error)
		Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error)
		Update(ctx context.Context, id string, entry map[string]any) (*entity.ProductCategory, error)
		Delete(ctx context.Context, id string) error
	}

	productCategoryRepositoryImpl struct {
		DB *gorm.DB
	}
)

// Create implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error) {
	panic("unimplemented")
}

// Delete implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) FindAll(ctx context.Context) ([]entity.ProductCategory, error) {
	tx := p.DB.WithContext(ctx)

	var data []entity.ProductCategory
	if err := tx.Find(&data).Error; err != nil {
		return []entity.ProductCategory{}, err
	}
	return data, nil
}

// FindByID implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.ProductCategory, error) {
	tx := p.DB.WithContext(ctx)

	productCategory := new(entity.ProductCategory)

	if err := tx.Where("id = ?", id).
		First(productCategory).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}
	return productCategory, nil
}

// Update implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Update(ctx context.Context, id string, entry map[string]any) (*entity.ProductCategory, error) {
	panic("unimplemented")
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepositoryImpl{
		DB: db,
	}
}
