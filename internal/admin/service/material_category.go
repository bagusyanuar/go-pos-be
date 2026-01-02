package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type materialCategoryServiceImpl struct {
	MaterialCategoryRepository domain.MaterialCategoryRepository
	Config                     *config.AppConfig
}

// Create implements domain.MaterialCategoryService.
func (m *materialCategoryServiceImpl) Create(ctx context.Context, schema *schema.MaterialCategoryRequest) error {
	e := entity.MaterialCategory{
		Name: schema.Name,
	}

	_, err := m.MaterialCategoryRepository.Create(ctx, &e)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.MaterialCategoryService.
func (m *materialCategoryServiceImpl) Delete(ctx context.Context, id string) error {
	// validate data is exists
	_, err := m.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = m.MaterialCategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Find implements domain.MaterialCategoryService.
func (m *materialCategoryServiceImpl) Find(ctx context.Context, queryParams *schema.MaterialCategoryQuery) ([]schema.MaterialCategoryResponse, *util.PaginationMeta, error) {
	data, pagination, err := m.MaterialCategoryRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.MaterialCategoryResponse{}, nil, err
	}

	res := schema.ToMaterialCategories(data)
	return res, pagination, nil
}

// FindByID implements domain.MaterialCategoryService.
func (m *materialCategoryServiceImpl) FindByID(ctx context.Context, id string) (*schema.MaterialCategoryResponse, error) {
	data, err := m.MaterialCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := schema.ToMaterialCategory(data)
	return res, nil
}

// Update implements domain.MaterialCategoryService.
func (m *materialCategoryServiceImpl) Update(ctx context.Context, id string, schema *schema.MaterialCategoryRequest) error {
	materialCategory, err := m.MaterialCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	materialCategory.Name = schema.Name

	_, err = m.MaterialCategoryRepository.Update(ctx, materialCategory)
	if err != nil {
		return err
	}
	return nil
}

func NewMaterialCategoryService(
	materialCategoryRepository domain.MaterialCategoryRepository,
	config *config.AppConfig,
) domain.MaterialCategoryService {
	return &materialCategoryServiceImpl{
		MaterialCategoryRepository: materialCategoryRepository,
		Config:                     config,
	}
}
