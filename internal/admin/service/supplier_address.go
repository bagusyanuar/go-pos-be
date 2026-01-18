package service

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

type supplierAddressServiceImpl struct {
	SupplierAddressRepository domain.SupplierAddressRepository
	SupplierRepository        domain.SupplierRepository
	Config                    *config.AppConfig
}
