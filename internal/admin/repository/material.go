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
	"gorm.io/gorm/clause"
)

type materialRepositoryImpl struct {
	DB *gorm.DB
}

// AppendUnit implements domain.MaterialRepository.
func (m *materialRepositoryImpl) AppendUnit(
	ctx context.Context,
	materialEntity *entity.Material,
	e []entity.MaterialUnit,
) error {
	tx := m.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if len(e) > 0 {
		for i := range e {
			e[i].MaterialID = materialEntity.ID
		}

		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

// Create implements domain.MateriaRepository.
func (m *materialRepositoryImpl) Create(ctx context.Context, e *entity.Material) (*entity.Material, error) {
	tx := m.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.
		Omit(clause.Associations).
		Create(&e).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return e, nil
}

// Delete implements domain.MateriaRepository.
func (m *materialRepositoryImpl) Delete(ctx context.Context, id string) error {
	tx := m.DB.WithContext(ctx)

	result := tx.Delete(&entity.Material{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exception.ErrRecordNotFound
	}

	return nil
}

// Find implements domain.MateriaRepository.
func (m *materialRepositoryImpl) Find(ctx context.Context, queryParams *schema.MaterialQuery) ([]entity.Material, *util.PaginationMeta, error) {
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

// FindByID implements domain.MateriaRepository.
func (m *materialRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Material, error) {
	tx := m.DB.WithContext(ctx)

	material := new(entity.Material)

	if err := tx.
		Preload("MaterialCategory").
		Preload("Units.Unit").
		Where("id = ?", id).
		First(material).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return material, nil
}

// Update implements domain.MateriaRepository.
func (m *materialRepositoryImpl) Update(ctx context.Context, e *entity.Material) (*entity.Material, error) {
	tx := m.DB.WithContext(ctx)

	if err := tx.Save(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// UploadImage implements domain.MaterialRepository.
func (m *materialRepositoryImpl) UploadImage(ctx context.Context, e []entity.MaterialImage) error {
	err := m.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&e).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (m *materialRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewMaterialRepository(db *gorm.DB) domain.MaterialRepository {
	return &materialRepositoryImpl{
		DB: db,
	}
}
