package repositories

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"gorm.io/gorm"
)

type supplierAddressRepositoryImpl struct {
	DB *gorm.DB
}

// FindBySupplierID implements domain.SupplierAddressRepository.
func (s *supplierAddressRepositoryImpl) FindBySupplierID(ctx context.Context, supplierID string) ([]entity.SupplierAddress, error) {
	panic("unimplemented")
}

// SyncAddresses implements domain.SupplierAddressRepository.
func (s *supplierAddressRepositoryImpl) SyncAddresses(ctx context.Context, supplierID string, addressEntities []entity.SupplierAddress) error {
	panic("unimplemented")
}

func NewSupplierAddressRepository(db *gorm.DB) domain.SupplierAddressRepository {
	return &supplierAddressRepositoryImpl{
		DB: db,
	}
}
