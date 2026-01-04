package schema

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
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
	Image       *string                   `json:"image"`
	Category    *MaterialCategoryResponse `json:"category"`
	Units       []MaterialUnitResponse    `json:"units"`
}

type MaterialUnitRequest struct {
	UnitID         uuid.UUID `json:"unit_id" validate:"required,uuid4"`
	ConversionRate float64   `json:"conversion_rate" validate:"required"`
	IsDefault      bool      `json:"is_default" validate:"required"`
}

type MaterialUnitResponse struct {
	UnitID         string  `json:"unit_id"`
	Name           string  `json:"name"`
	ConversionRate float64 `json:"conversion_rate"`
	IsDefault      bool    `json:"is_default"`
	Quantity       float64 `json:"quantity"`
}

func ToMaterial(material *entity.Material) *MaterialResponse {
	if material == nil {
		return nil
	}

	category := new(MaterialCategoryResponse)
	if material.MaterialCategory != nil {
		category.ID = material.MaterialCategory.ID.String()
		category.Name = material.MaterialCategory.Name
	}

	units := make([]MaterialUnitResponse, 0, len(material.Units))
	for _, v := range material.Units {
		unit := MaterialUnitResponse{
			UnitID:         v.UnitID.String(),
			Name:           v.Unit.Name,
			ConversionRate: v.ConversionRate,
			IsDefault:      v.IsDefault,
		}
		units = append(units, unit)
	}

	return &MaterialResponse{
		ID:          material.ID.String(),
		Name:        material.Name,
		Description: material.Description,
		Image:       material.Image,
		Category:    category,
		Units:       units,
	}
}

func ToMaterials(materials []entity.Material) []MaterialResponse {
	responses := make([]MaterialResponse, 0, len(materials))
	for _, v := range materials {
		res := ToMaterial(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
