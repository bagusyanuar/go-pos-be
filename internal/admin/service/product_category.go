package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type productCategoryServiceImpl struct {
	ProductCategoryRepository domain.ProductCategoryRepository
	Config                    *config.AppConfig
}

// Create implements ProductCategoryService.
func (p *productCategoryServiceImpl) Create(ctx context.Context, schema *schema.ProductCategoryRequest) error {
	e := entity.ProductCategory{
		Name: schema.Name,
	}

	_, err := p.ProductCategoryRepository.Create(ctx, &e)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements ProductCategoryService.
func (p *productCategoryServiceImpl) Delete(ctx context.Context, id string) error {

	// validate data is exists
	_, err := p.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = p.ProductCategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryService.
func (p *productCategoryServiceImpl) Find(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]schema.ProductCategoryResponse, *util.PaginationMeta, error) {
	p.Config.Logger.Info("run find all service")
	data, pagination, err := p.ProductCategoryRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.ProductCategoryResponse{}, nil, err
	}

	res := schema.ToProductCategories(data)
	return res, pagination, nil
}

// FindByID implements ProductCategoryService.
func (p *productCategoryServiceImpl) FindByID(ctx context.Context, id string) (*schema.ProductCategoryResponse, error) {
	data, err := p.ProductCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := schema.ToProductCategory(data)
	return res, nil
}

// Update implements ProductCategoryService.
func (p *productCategoryServiceImpl) Update(ctx context.Context, id string, schema *schema.ProductCategoryRequest) error {
	productCategory, err := p.ProductCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	productCategory.Name = schema.Name

	_, err = p.ProductCategoryRepository.Update(ctx, productCategory)
	if err != nil {
		return err
	}
	return nil
}

func NewProductCategoryService(
	productCategoryRepository domain.ProductCategoryRepository,
	config *config.AppConfig,
) domain.ProductCategoryService {
	return &productCategoryServiceImpl{
		ProductCategoryRepository: productCategoryRepository,
		Config:                    config,
	}
}
