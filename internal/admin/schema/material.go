package schema

import (
	"mime/multipart"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/google/uuid"
)

type MaterialRequest struct {
	CategoryID  *uuid.UUID            `form:"category_id" validate:"uuid4"`
	Name        string                `form:"name" validate:"required"`
	Description *string               `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
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

	return &MaterialResponse{
		ID:          material.ID.String(),
		Name:        material.Name,
		Description: material.Description,
		Image:       material.Image,
		Category:    category,
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
