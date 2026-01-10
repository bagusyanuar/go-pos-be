package mapper

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

func ToMaterial(material *entity.Material) *schema.MaterialResponse {
	if material == nil {
		return nil
	}

	category := new(schema.MaterialMaterialCategoryResponse)
	if material.MaterialCategory != nil {
		category.ID = material.MaterialCategory.ID.String()
		category.Name = material.MaterialCategory.Name
	}

	units := make([]schema.MaterialUnitResponse, 0, len(material.Units))
	for _, v := range material.Units {
		unit := schema.MaterialUnitResponse{
			UnitID:         v.UnitID.String(),
			Name:           v.Unit.Name,
			ConversionRate: v.ConversionRate,
			IsDefault:      v.IsDefault,
		}
		units = append(units, unit)
	}

	return &schema.MaterialResponse{
		ID:          material.ID.String(),
		Name:        material.Name,
		Description: material.Description,
		CreatedAt:   material.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt:   material.UpdatedAt.Format(constant.BaseDateTimeLayout),
		Category:    category,
		Units:       units,
	}
}

func ToMaterials(materials []entity.Material) []schema.MaterialResponse {
	responses := make([]schema.MaterialResponse, 0, len(materials))
	for _, v := range materials {
		res := ToMaterial(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}

func ToMaterialCreate(material *entity.Material) *schema.MaterialCreateResponse {
	if material == nil {
		return nil
	}

	return &schema.MaterialCreateResponse{
		ID:          material.ID.String(),
		Name:        material.Name,
		Description: material.Description,
	}
}
