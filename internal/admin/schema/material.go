package schema

import (
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/google/uuid"
)

type MaterialRequest struct {
	CategoryID  *uuid.UUID            `json:"category_id" validate:"uuid4"`
	Name        string                `json:"name" validate:"required"`
	Description *string               `json:"description"`
	Units       []MaterialUnitRequest `json:"units" validate:"required"`
}

type MaterialQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

type MaterialResponse struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Description *string                   `json:"description"`
	CreatedAt   string                    `json:"created_at"`
	UpdatedAt   string                    `json:"updated_at"`
	Category    *MaterialMaterialCategory `json:"category"`
	Units       []MaterialUnit            `json:"units"`
}

type MaterialUnitRequest struct {
	UnitID         uuid.UUID `json:"unit_id" validate:"required,uuid4"`
	ConversionRate float64   `json:"conversion_rate" validate:"required"`
	IsDefault      bool      `json:"is_default" validate:"required"`
}

type MaterialMaterialCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MaterialUnit struct {
	UnitID         string  `json:"unit_id"`
	Name           string  `json:"name"`
	ConversionRate float64 `json:"conversion_rate"`
	IsDefault      bool    `json:"is_default"`
}
