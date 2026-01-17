package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

type (
	SupplierAddressRepository interface {
		FindBySupplierID(ctx context.Context, supplierID string) ([]entity.SupplierAddress, error)
		SyncAddresses(ctx context.Context, supplierID string, addressEntities []entity.SupplierAddress) error
	}

	SupplierAddressService interface {
		FindBySupplierID(ctx context.Context, supplierID string) ([]schema.SupplierAddressResponse, error)
		SyncAddresses(ctx context.Context, supplierID string) error
	}
)
