package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

type supplierServiceImpl struct {
	SupplierRepository domain.SupplierRepository
	Config             *config.AppConfig
}

// DeleteContact implements domain.SupplierService.
func (s *supplierServiceImpl) DeleteContact(ctx context.Context, supplierID string, contactID string) error {
	_, err := s.SupplierRepository.FindByID(ctx, supplierID)

	if err != nil {
		return err
	}

	err = s.SupplierRepository.DeleteContact(ctx, contactID)
	if err != nil {
		return err
	}

	return nil
}

// AddContacts implements domain.SupplierService.
func (s *supplierServiceImpl) AddContacts(ctx context.Context, id string, schema *schema.SupplierContactRequest) error {
	supplier, err := s.SupplierRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	newContacts := make([]entity.SupplierContact, 0, len(schema.Contacts))

	for _, contact := range schema.Contacts {
		newContacts = append(newContacts, entity.SupplierContact{
			SupplierID: &supplier.ID,
			Type:       contact.Type,
			Value:      contact.Value,
		})
	}

	return s.SupplierRepository.AddContacts(ctx, supplier, newContacts)
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
func (s *supplierServiceImpl) Find(ctx context.Context, queryParams *schema.SupplierQuery) ([]entity.Supplier, int64, error) {
	return s.SupplierRepository.Find(ctx, queryParams)
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
