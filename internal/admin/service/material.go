package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type materialServiceImpl struct {
	MaterialRepository domain.MaterialRepository
	Config             *config.AppConfig
}

// Create implements domain.MaterialService.
func (m *materialServiceImpl) Create(ctx context.Context, schema *schema.MaterialRequest) error {
	e := entity.Material{
		MaterialCategoryID: schema.CategoryID,
		Name:               schema.Name,
		Description:        schema.Description,
		Image:              nil,
	}

	_, err := m.MaterialRepository.Create(ctx, &e)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.MaterialService.
func (m *materialServiceImpl) Delete(ctx context.Context, id string) error {
	_, err := m.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = m.MaterialRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Find implements domain.MaterialService.
func (m *materialServiceImpl) Find(ctx context.Context, queryParams *schema.MaterialQuery) ([]schema.MaterialResponse, *util.PaginationMeta, error) {
	data, pagination, err := m.MaterialRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.MaterialResponse{}, nil, err
	}

	res := schema.ToMaterials(data)
	return res, pagination, nil
}

// FindByID implements domain.MaterialService.
func (m *materialServiceImpl) FindByID(ctx context.Context, id string) (*schema.MaterialResponse, error) {
	panic("unimplemented")
}

// Update implements domain.MaterialService.
func (m *materialServiceImpl) Update(ctx context.Context, id string, schema *schema.MaterialRequest) error {
	panic("unimplemented")
}

func NewMaterialService(
	materialRepository domain.MaterialRepository,
	config *config.AppConfig,
) domain.MaterialService {
	return &materialServiceImpl{
		MaterialRepository: materialRepository,
		Config:             config,
	}
}
