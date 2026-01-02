package schema

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

/* === Request & Query Material Category === */

type MaterialCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type MaterialCategoryQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

/* === Response Material Category === */

type MaterialCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToMaterialCategory(materialCategory *entity.MaterialCategory) *MaterialCategoryResponse {
	if materialCategory == nil {
		return nil
	}
	return &MaterialCategoryResponse{
		ID:   materialCategory.ID.String(),
		Name: materialCategory.Name,
	}
}

func ToMaterialCategories(materialCategories []entity.MaterialCategory) []MaterialCategoryResponse {
	responses := make([]MaterialCategoryResponse, 0, len(materialCategories))
	for _, v := range materialCategories {
		res := ToMaterialCategory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
