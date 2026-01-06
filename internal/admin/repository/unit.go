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

type unitRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements domain.UnitRepository.
func (u *unitRepositoryImpl) Create(ctx context.Context, e *entity.Unit) (*entity.Unit, error) {
	tx := u.DB.WithContext(ctx)
	if err := tx.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// Delete implements domain.UnitRepository.
func (u *unitRepositoryImpl) Delete(ctx context.Context, id string) error {
	tx := u.DB.WithContext(ctx)

	result := tx.Delete(&entity.Unit{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exception.ErrRecordNotFound
	}

	return nil
}

// Find implements domain.UnitRepository.
func (u *unitRepositoryImpl) Find(
	ctx context.Context,
	queryParams *schema.UnitQuery,
) (
	[]entity.Unit,
	*util.PaginationMeta,
	error,
) {
	tx := u.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.Unit

	if err := tx.
		Model(&entity.Unit{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.Unit{}, nil, err
	}

	sortFieldMap := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}
	sort := util.GetSortField(queryParams.Sort, "created_at", sortFieldMap)
	order := util.GetOrder(queryParams.Order)

	if err := tx.
		Scopes(
			u.filterByParam(queryParams.Param),
			util.SortScope(sort, order),
			util.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).
		Error; err != nil {
		return []entity.Unit{}, nil, err
	}

	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)
	return data, &pagination, nil
}

// FindByID implements domain.UnitRepository.
func (u *unitRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Unit, error) {
	tx := u.DB.WithContext(ctx)

	unit := new(entity.Unit)

	if err := tx.Where("id = ?", id).
		First(unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return unit, nil
}

// Update implements domain.UnitRepository.
func (u *unitRepositoryImpl) Update(ctx context.Context, e *entity.Unit) (*entity.Unit, error) {
	tx := u.DB.WithContext(ctx)

	if err := tx.Save(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (m *unitRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewUnitRepository(db *gorm.DB) domain.UnitRepository {
	return &unitRepositoryImpl{
		DB: db,
	}
}
