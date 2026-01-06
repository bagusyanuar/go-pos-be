package mapper

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

func ToMaterialCategory(materialCategory *entity.MaterialCategory) *schema.MaterialCategoryResponse {
	if materialCategory == nil {
		return nil
	}

	return &schema.MaterialCategoryResponse{
		ID:        materialCategory.ID.String(),
		Name:      materialCategory.Name,
		CreatedAt: materialCategory.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt: materialCategory.UpdatedAt.Format(constant.BaseDateTimeLayout),
	}
}

func ToMaterialCategories(materialCategories []entity.MaterialCategory) []schema.MaterialCategoryResponse {
	responses := make([]schema.MaterialCategoryResponse, 0, len(materialCategories))
	for _, v := range materialCategories {
		res := ToMaterialCategory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
