package schema

import "github.com/bagusyanuar/go-pos-be/pkg/util"

// Supplier Request Schema
type SupplierRequest struct {
	Name string `json:"name" validate:"required"`
}

type SupplierContactRequest struct {
	Contacts []SupplierContact `json:"contacts" validate:"required,dive"`
}

type SupplierContact struct {
	Type  string `json:"type" validate:"required,contact_type"`
	Value string `json:"value" validate:"required"`
}

// Supplier Query Param Schema
type SupplierQuery struct {
	Param string `json:"param"`
	util.QueryPagination
	util.QuerySort
}

// Supplier Response Schema
type SupplierResponse struct {
	ID        string                    `json:"id"`
	Name      string                    `json:"name"`
	CreatedAt string                    `json:"created_at"`
	UpdatedAt string                    `json:"updated_at"`
	Contacts  []SupplierContactResponse `json:"contacts"`
}

type SupplierCreateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SupplierContactResponse struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
