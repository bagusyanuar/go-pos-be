package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type (
	// 1. Repository Interface (Kontrak untuk DB)
	ProductCategoryRepository interface {
		Find(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]entity.ProductCategory, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*entity.ProductCategory, error)
		Create(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error)
		Update(ctx context.Context, e *entity.ProductCategory) (*entity.ProductCategory, error)
		Delete(ctx context.Context, id string) error
	}

	// 2. Service Interface (Kontrak untuk business logic / usecase)
	ProductCategoryService interface {
		Find(ctx context.Context, queryParams *schema.ProductCategoryQuery) ([]schema.ProductCategoryResponse, *util.PaginationMeta, error)
		FindByID(ctx context.Context, id string) (*schema.ProductCategoryResponse, error)
		Create(ctx context.Context, schema *schema.ProductCategoryRequest) error
		Update(ctx context.Context, id string, schema *schema.ProductCategoryRequest) error
		Delete(ctx context.Context, id string) error
	}
)
