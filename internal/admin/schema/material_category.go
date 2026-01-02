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

func ToMaterialCategories(productCategories []entity.ProductCategory) []ProductCategoryResponse {
	responses := make([]ProductCategoryResponse, 0, len(productCategories))
	for _, v := range productCategories {
		res := ToProductCategory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
