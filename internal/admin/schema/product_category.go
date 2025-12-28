package schema

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

/* === Request & Query Product Category === */

type ProductCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductCategoryQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

/* === Response Product Category === */

type ProductCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToProductCategory(productCategory *entity.ProductCategory) *ProductCategoryResponse {
	if productCategory == nil {
		return nil
	}
	return &ProductCategoryResponse{
		ID:   productCategory.ID.String(),
		Name: productCategory.Name,
	}
}

func ToProductCategories(productCategories []entity.ProductCategory) []ProductCategoryResponse {
	responses := make([]ProductCategoryResponse, 0, len(productCategories))
	for _, v := range productCategories {
		res := ToProductCategory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
