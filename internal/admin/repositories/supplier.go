package repositories

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

type supplierRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements domain.SupplierRepository.
func (s *supplierRepositoryImpl) Create(ctx context.Context, supplierEntity *entity.Supplier) (*entity.Supplier, error) {
	tx := s.DB.WithContext(ctx)
	if err := tx.Create(&supplierEntity).Error; err != nil {
		return nil, err
	}

	return supplierEntity, nil
}

// Delete implements domain.SupplierRepository.
func (s *supplierRepositoryImpl) Delete(ctx context.Context, id string) error {
	tx := s.DB.WithContext(ctx)

	result := tx.Delete(&entity.Supplier{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exception.ErrRecordNotFound
	}

	return nil
}

// Find implements domain.SupplierRepository.
func (s *supplierRepositoryImpl) Find(
	ctx context.Context,
	queryParams *schema.SupplierQuery,
) (
	[]entity.Supplier,
	int64,
	error,
) {
	tx := s.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.Supplier

	if err := tx.
		Model(&entity.Supplier{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.Supplier{}, 0, err
	}

	sortFieldMap := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}
	sort := util.GetSortField(
		queryParams.Sort,
		"created_at",
		sortFieldMap,
	)
	order := util.GetOrder(queryParams.Order)

	if err := tx.
		Preload("Contacts").
		Scopes(
			s.filterByParam(queryParams.Param),
			util.SortScope(sort, order),
			util.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).
		Error; err != nil {
		return []entity.Supplier{}, 0, err
	}

	// pagination := util.MakePagination(
	// 	queryParams.Page,
	// 	queryParams.PageSize,
	// 	totalItems,
	// )
	return data, totalItems, nil
}

// FindByID implements domain.SupplierRepository.
func (s *supplierRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Supplier, error) {
	tx := s.DB.WithContext(ctx)

	supplier := new(entity.Supplier)

	if err := tx.
		Preload("Contacts").
		Where("id = ?", id).
		First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrRecordNotFound
		}
		return nil, err
	}

	return supplier, nil
}

// Update implements domain.SupplierRepository.
func (s *supplierRepositoryImpl) Update(ctx context.Context, supplierEnityty *entity.Supplier) (*entity.Supplier, error) {
	tx := s.DB.WithContext(ctx)

	if err := tx.Save(supplierEnityty).Error; err != nil {
		return nil, err
	}

	return supplierEnityty, nil
}

func (s *supplierRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewSupplierRepository(db *gorm.DB) domain.SupplierRepository {
	return &supplierRepositoryImpl{
		DB: db,
	}
}
