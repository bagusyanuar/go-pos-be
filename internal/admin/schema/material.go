package schema

import (
	"mime/multipart"

	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/google/uuid"
)

// Material Request Map
type MaterialRequest struct {
	CategoryID  *uuid.UUID `json:"category_id" validate:"uuid4"`
	Name        string     `json:"name" validate:"required"`
	Description *string    `json:"description"`
}

type MaterialUnitRequest struct {
	Type  constant.MaterialUnitActionType `json:"type" validate:"required,oneof=create append"`
	Units []MaterialUnit                  `json:"units" validate:"required,dive"`
}

type MaterialUnit struct {
	UnitID         uuid.UUID `json:"unit_id" validate:"required,uuid4"`
	ConversionRate float64   `json:"conversion_rate" validate:"gt=0"`
	IsDefault      *bool     `json:"is_default" validate:"required,boolean"`
}

// Material Query Map

type MaterialQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

// Material Response Map

type MaterialCreateResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type MaterialResponse struct {
	ID          string                            `json:"id"`
	Name        string                            `json:"name"`
	Description *string                           `json:"description"`
	CreatedAt   string                            `json:"created_at"`
	UpdatedAt   string                            `json:"updated_at"`
	Category    *MaterialMaterialCategoryResponse `json:"category"`
	Units       []MaterialUnitResponse            `json:"units"`
}

type MaterialMaterialCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MaterialUnitResponse struct {
	UnitID         string  `json:"unit_id"`
	Name           string  `json:"name"`
	ConversionRate float64 `json:"conversion_rate"`
	IsDefault      bool    `json:"is_default"`
}

type MaterialImageRequest struct {
	Image *multipart.FileHeader `form:"image"`
}
