package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

type (
	// 1. Repository Interface for Supplier
	SupplierRepository interface {
		Find(ctx context.Context, queryParams *schema.SupplierQuery) ([]entity.Supplier, int64, error)
		FindByID(ctx context.Context, id string) (*entity.Supplier, error)
		Create(ctx context.Context, supplierEntity *entity.Supplier) (*entity.Supplier, error)
		Update(ctx context.Context, supplierEnityty *entity.Supplier) (*entity.Supplier, error)
		Delete(ctx context.Context, id string) error
	}

	// 2. Service Interface for Supplier
	SupplierService interface {
		Find(ctx context.Context, queryParams *schema.SupplierQuery) ([]entity.Supplier, int64, error)
		FindByID(ctx context.Context, id string) (*schema.SupplierResponse, error)
		Create(ctx context.Context, schema *schema.SupplierRequest) (*schema.SupplierCreateResponse, error)
		Update(ctx context.Context, id string, schema *schema.SupplierRequest) error
		Delete(ctx context.Context, id string) error
	}
)
