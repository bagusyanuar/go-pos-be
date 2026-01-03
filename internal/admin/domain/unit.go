package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type (
	UnitRepository interface {
		Find(ctx context.Context, queryParams *schema.UnitQuery) ([]entity.Unit, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.Unit, error)
		Create(ctx context.Context, e *entity.Unit) (*entity.Unit, error)
		Update(ctx context.Context, e *entity.Unit) (*entity.Unit, error)
		Delete(ctx context.Context, id string) error
	}

	UnitService interface {
		Find(ctx context.Context, queryParams *schema.UnitQuery) ([]schema.UnitResponse, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*schema.UnitResponse, error)
		Create(ctx context.Context, schema *schema.UnitRequest) error
		Update(ctx context.Context, id string, schema *schema.UnitRequest) error
		Delete(ctx context.Context, id string) error
	}
)
