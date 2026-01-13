package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type supplierServiceImpl struct {
	SupplierRepository domain.SupplierRepository
	Config             *config.AppConfig
}

// Create implements domain.SupplierService.
func (s *supplierServiceImpl) Create(ctx context.Context, schema *schema.SupplierRequest) (*schema.SupplierCreateResponse, error) {
	supplierEntity := entity.Supplier{
		Name: schema.Name,
	}

	supplier, err := s.SupplierRepository.Create(ctx, &supplierEntity)
	if err != nil {
		return nil, err
	}

	response := mapper.ToSupplierCreate(supplier)
	return response, nil
}

// Delete implements domain.SupplierService.
func (s *supplierServiceImpl) Delete(ctx context.Context, id string) error {
	_, err := s.SupplierRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = s.SupplierRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Find implements domain.SupplierService.
func (s *supplierServiceImpl) Find(ctx context.Context, queryParams *schema.SupplierQuery) ([]schema.SupplierResponse, *util.PaginationMeta, error) {
	data, pagination, err := s.SupplierRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.SupplierResponse{}, nil, err
	}

	response := mapper.ToSuppliers(data)
	return response, pagination, nil
}

// FindByID implements domain.SupplierService.
func (s *supplierServiceImpl) FindByID(ctx context.Context, id string) (*schema.SupplierResponse, error) {
	data, err := s.SupplierRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	response := mapper.ToSupplier(data)
	return response, nil
}

// Update implements domain.SupplierService.
func (s *supplierServiceImpl) Update(ctx context.Context, id string, schema *schema.SupplierRequest) error {
	supplier, err := s.SupplierRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	supplier.Name = schema.Name

	_, err = s.SupplierRepository.Update(ctx, supplier)
	if err != nil {
		return err
	}
	return nil
}

func NewSupplierService(
	supplierRepository domain.SupplierRepository,
	config *config.AppConfig,
) domain.SupplierService {
	return &supplierServiceImpl{
		SupplierRepository: supplierRepository,
		Config:             config,
	}
}
