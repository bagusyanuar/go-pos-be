package schema

import (
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type MaterialInventoryQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

type MaterialInventoryResponse struct {
	ID          string                             `json:"id"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description"`
	Quantity    float64                            `json:"quantity"`
	CreatedAt   string                             `json:"created_at"`
	UpdatedAt   string                             `json:"updated_at"`
	Category    *MaterialInventoryMaterialCategory `json:"category"`
	Units       []MaterialInventoryUnit            `json:"units"`
}

type MaterialInventoryMaterialCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MaterialInventoryUnit struct {
	UnitID         string  `json:"unit_id"`
	Name           string  `json:"name"`
	ConversionRate float64 `json:"conversion_rate"`
	IsDefault      bool    `json:"is_default"`
	Quantity       float64 `json:"quantity"`
}
