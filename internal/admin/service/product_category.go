package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/repository"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

type (
	ProductCategoryService interface {
		FindAll(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]schema.ProductCategoryResponse, error)
		FindByID(ctx context.Context, id string) (*schema.ProductCategoryResponse, error)
		Create(ctx context.Context, schema *schema.ProductCategoryRequest) error
		Update(ctx context.Context, id string, schema *schema.ProductCategoryRequest) error
		Delete(ctx context.Context, id string) error
	}

	productCategoryServiceImpl struct {
		ProductCategoryRepository repository.ProductCategoryRepository
		Config                    *config.AppConfig
	}
)

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
	err := p.ProductCategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryService.
func (p *productCategoryServiceImpl) FindAll(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]schema.ProductCategoryResponse, error) {
	data, err := p.ProductCategoryRepository.FindAll(ctx)
	if err != nil {
		return []schema.ProductCategoryResponse{}, err
	}

	res := schema.ToProductCategories(data)
	return res, nil
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
	entry := map[string]any{
		"name": schema.Name,
	}

	_, err := p.ProductCategoryRepository.Update(ctx, id, entry)
	if err != nil {
		return err
	}

	return nil
}

func NewProductCategoryService(
	productCategoryRepository repository.ProductCategoryRepository,
	config *config.AppConfig,
) ProductCategoryService {
	return &productCategoryServiceImpl{
		ProductCategoryRepository: productCategoryRepository,
		Config:                    config,
	}
}
