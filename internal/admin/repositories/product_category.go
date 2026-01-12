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

	result := tx.Delete(&entity.ProductCategory{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exception.ErrRecordNotFound
	}

	return nil
}

// FindAll implements ProductCategoryRepository.
func (p *productCategoryRepositoryImpl) Find(
	ctx context.Context,
	queryParams *schema.ProductCategoryQuery,
) (
	[]entity.ProductCategory,
	*util.PaginationMeta,
	error,
) {
	tx := p.DB.WithContext(ctx)

	var totalItems int64
	var data []entity.ProductCategory

	if err := tx.
		Model(&entity.ProductCategory{}).
		Count(&totalItems).
		Error; err != nil {
		return []entity.ProductCategory{}, nil, err
	}

	sortFieldMap := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}
	sort := util.GetSortField(queryParams.Sort, "created_at", sortFieldMap)
	order := util.GetOrder(queryParams.Order)

	if err := tx.
		Scopes(
			p.filterByParam(queryParams.Param),
			util.SortScope(sort, order),
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

func (p *productCategoryRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}

		searchQuery := fmt.Sprintf("%%%s%%", param)
		return tx.
			Where("name ILIKE ?", searchQuery)
	}
}

func NewProductCategoryRepository(db *gorm.DB) domain.ProductCategoryRepository {
	return &productCategoryRepositoryImpl{
		DB: db,
	}
}
