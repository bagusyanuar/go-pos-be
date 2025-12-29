package repository

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"gorm.io/gorm"
)

type productCategoryRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error) {
	tx := p.DB.WithContext(ctx)
	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// Delete implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	tx := p.DB.WithContext(ctx)
	productCategory, err := p.getCategoryByID(tx, id)

	if err != nil {
		return err
	}

	if err := tx.Delete(productCategory).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]entity.ProductCategory, *util.PaginationMeta, error) {
	tx := p.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.ProductCategory

	if err := tx.
		Model(&entity.ProductCategory{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.ProductCategory{}, nil, err
	}

	if err := tx.
		Scopes(
			util.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).
		Error; err != nil {
		return []entity.ProductCategory{}, nil, err
	}

	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)
	return data, &pagination, nil
}

// FindByID implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.ProductCategory, error) {
	tx := p.DB.WithContext(ctx)

	productCategory := new(entity.ProductCategory)

	if err := tx.Where("id = ?", id).
		First(productCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return productCategory, nil
}

// Update implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Update(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error) {
	tx := p.DB.WithContext(ctx)

	if err := tx.Save(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (p *productCategoryRepositoryImpl) getCategoryByID(tx *gorm.DB, id string) (*entity.ProductCategory, error) {
	productCategory := new(entity.ProductCategory)

	if err := tx.Where("id = ?", id).
		First(productCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return productCategory, nil
}

func NewProductCategoryRepository(db *gorm.DB) domain.ProductCategoryRepository {
	return &productCategoryRepositoryImpl{
		DB: db,
	}
}
