package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type materialInventoryServiceImpl struct {
	MaterialInventoryRepository domain.MaterialInventoryRepository
	Config                      *config.AppConfig
}

// Find implements domain.MaterialInventoryService.
func (m *materialInventoryServiceImpl) Find(ctx context.Context, queryParams *schema.MaterialInventoryQuery) ([]schema.MaterialInventoryResponse, *util.PaginationMeta, error) {
	data, pagination, err := m.MaterialInventoryRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.MaterialInventoryResponse{}, nil, err
	}

	res := mapper.ToMaterialInventories(data)
	return res, pagination, nil
}

// FindByID implements domain.MaterialInventoryService.
func (m *materialInventoryServiceImpl) FindByID(ctx context.Context, id string) (*schema.MaterialInventoryResponse, error) {
	data, err := m.MaterialInventoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := mapper.ToMaterialInventory(data)
	return res, nil
}

func NewMaterialInventoryService(
	materialInventoryRepository domain.MaterialInventoryRepository,
	config *config.AppConfig,
) domain.MaterialInventoryService {
	return &materialInventoryServiceImpl{
		MaterialInventoryRepository: materialInventoryRepository,
		Config:                      config,
	}
}
