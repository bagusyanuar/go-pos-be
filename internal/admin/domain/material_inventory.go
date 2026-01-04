package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type (
	MaterialInventoryRepository interface {
		Find(ctx context.Context, queryParams *schema.MaterialInventoryQuery) ([]entity.Material, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Material, error)
	}

	MaterialInventoryService interface {
		Find(ctx context.Context, queryParams *schema.MaterialInventoryQuery) ([]schema.MaterialInventoryResponse, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*schema.MaterialInventoryResponse, error)
	}
)
