package schema

import "github.com/bagusyanuar/go-pos-be/pkg/util"

type SupplierAddressRequest struct {
	Addresses []SupplierAddress `json:"addresses" validate:"required,dive"`
}

type SupplierAddress struct {
	Type  string `json:"type" validate:"required,address_type"`
	Value string `json:"value" validate:"required"`
}

type SupplierAddressQuery struct {
	SupplierID string `json:"supplier_id"`
	util.QueryPagination
	util.QuerySort
}

type SupplierAddressResponse struct {
	ID         string `json:"id"`
	SupplierID string `json:"supplier_id"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
