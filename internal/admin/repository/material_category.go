package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"gorm.io/gorm"
)

type materialCategoryRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements domain.MaterialCategoryRepository.
func (m *materialCategoryRepositoryImpl) Create(ctx context.Context, e *entity.MaterialCategory) (*entity.MaterialCategory, error) {
	tx := m.DB.WithContext(ctx)
	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// Delete implements domain.MaterialCategoryRepository.
func (m *materialCategoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	tx := m.DB.WithContext(ctx)

	result := tx.Delete(&entity.MaterialCategory{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exception.ErrRecordNotFound
	}

	return nil
}

// Find implements domain.MaterialCategoryRepository.
func (m *materialCategoryRepositoryImpl) Find(ctx context.Context, queryParams *schema.MaterialCategoryQuery) ([]entity.MaterialCategory, *util.PaginationMeta, error) {
	tx := m.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.MaterialCategory

	if err := tx.
		Model(&entity.MaterialCategory{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.MaterialCategory{}, nil, err
	}

	if err := tx.
		Scopes(
			m.filterByParam(queryParams.Param),
			util.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).
		Error; err != nil {
		return []entity.MaterialCategory{}, nil, err
	}

	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)
	return data, &pagination, nil
}

// FindByID implements domain.MaterialCategoryRepository.
func (m *materialCategoryRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.MaterialCategory, error) {
	tx := m.DB.WithContext(ctx)

	materialCategory := new(entity.MaterialCategory)

	if err := tx.Where("id = ?", id).
		First(materialCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return materialCategory, nil
}

// Update implements domain.MaterialCategoryRepository.
func (m *materialCategoryRepositoryImpl) Update(ctx context.Context, e *entity.MaterialCategory) (*entity.MaterialCategory, error) {
	tx := m.DB.WithContext(ctx)

	if err := tx.Save(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (m *materialCategoryRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewMaterialCategoryRepository(db *gorm.DB) domain.MaterialCategoryRepository {
	return &materialCategoryRepositoryImpl{
		DB: db,
	}
}
