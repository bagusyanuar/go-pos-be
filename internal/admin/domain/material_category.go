package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type (
	// 1. Repository Interface (Kontrak untuk DB)
	MaterialCategoryRepository interface {
		Find(ctx context.Context, queryParams *schema.MaterialCategoryQuery) ([]entity.MaterialCategory, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.MaterialCategory, error)
		Create(ctx context.Context, e *entity.MaterialCategory) (*entity.MaterialCategory, error)
		Update(ctx context.Context, e *entity.MaterialCategory) (*entity.MaterialCategory, error)
		Delete(ctx context.Context, id string) error
	}

	// 2. Service Interface (Kontrak untuk business logic / usecase)
	MaterialCategoryService interface {
		Find(ctx context.Context, queryParams *schema.MaterialCategoryQuery) ([]schema.MaterialCategoryResponse, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*schema.MaterialCategoryResponse, error)
		Create(ctx context.Context, schema *schema.MaterialCategoryRequest) error
		Update(ctx context.Context, id string, schema *schema.MaterialCategoryRequest) error
		Delete(ctx context.Context, id string) error
	}
)
