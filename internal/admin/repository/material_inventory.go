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

type materialInventoryRepositoryImpl struct {
	DB *gorm.DB
}

// Find implements domain.MaterialInventoryRepository.
func (m *materialInventoryRepositoryImpl) Find(ctx context.Context, queryParams *schema.MaterialInventoryQuery) ([]entity.Material, *util.PaginationMeta, error) {
	tx := m.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.Material

	if err := tx.
		Model(&entity.Material{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.Material{}, nil, err
	}

	if err := tx.
		Preload("MaterialCategory").
		Preload("Units.Unit").
		Preload("Inventory").
		Scopes(
			m.filterByParam(queryParams.Param),
			util.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).
		Error; err != nil {
		return []entity.Material{}, nil, err
	}

	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)
	return data, &pagination, nil
}

// FindByID implements domain.MaterialInventoryRepository.
func (m *materialInventoryRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Material, error) {
	tx := m.DB.WithContext(ctx)

	material := new(entity.Material)

	if err := tx.
		Preload("MaterialCategory").
		Preload("Units.Unit").
		Preload("Inventory").
		Where("id = ?", id).
		First(material).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return material, nil
}

func (m *materialInventoryRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewMaterialInventoryRepository(db *gorm.DB) domain.MaterialInventoryRepository {
	return &materialInventoryRepositoryImpl{
		DB: db,
	}
}
