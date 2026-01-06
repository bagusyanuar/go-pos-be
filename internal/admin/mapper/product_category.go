package mapper

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

func ToProductCategory(productCategory *entity.ProductCategory) *schema.ProductCategoryResponse {
	if productCategory == nil {
		return nil
	}
	return &schema.ProductCategoryResponse{
		ID:        productCategory.ID.String(),
		Name:      productCategory.Name,
		CreatedAt: productCategory.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt: productCategory.UpdatedAt.Format(constant.BaseDateTimeLayout),
	}
}

func ToProductCategories(productCategories []entity.ProductCategory) []schema.ProductCategoryResponse {
	responses := make([]schema.ProductCategoryResponse, 0, len(productCategories))
	for _, v := range productCategories {
		res := ToProductCategory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
