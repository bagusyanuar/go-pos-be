package mapper

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

func ToSupplier(supplier *entity.Supplier) *schema.SupplierResponse {
	if supplier == nil {
		return nil
	}

	return &schema.SupplierResponse{
		ID:        supplier.ID.String(),
		Name:      supplier.Name,
		CreatedAt: supplier.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt: supplier.UpdatedAt.Format(constant.BaseDateTimeLayout),
	}
}

func ToSupplierCreate(supplier *entity.Supplier) *schema.SupplierCreateResponse {
	if supplier == nil {
		return nil
	}

	return &schema.SupplierCreateResponse{
		ID:   supplier.ID.String(),
		Name: supplier.Name,
	}
}

func ToSuppliers(suppliers []entity.Supplier) []schema.SupplierResponse {
	responses := make([]schema.SupplierResponse, 0, len(suppliers))
	for _, supplier := range suppliers {
		response := ToSupplier(&supplier)
		if response != nil {
			responses = append(responses, *response)
		}
	}
	return responses
}
