package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type (
	MaterialRepository interface {
		Find(ctx context.Context, queryParams *schema.MaterialQuery) ([]entity.Material, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Material, error)
		Create(ctx context.Context, e *entity.Material) (*entity.Material, error)
		Update(ctx context.Context, e *entity.Material) (*entity.Material, error)
		Delete(ctx context.Context, id string) error
	}

	MaterialService interface {
		Find(ctx context.Context, queryParams *schema.MaterialQuery) ([]schema.MaterialResponse, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*schema.MaterialResponse, error)
		Create(ctx context.Context, schema *schema.MaterialRequest) error
		Update(ctx context.Context, id string, schema *schema.MaterialRequest) error
		Delete(ctx context.Context, id string) error
	}
)
