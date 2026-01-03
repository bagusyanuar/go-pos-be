package schema

import (
	"mime/multipart"

	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type MaterialRequest struct {
	CategoryID  string                `form:"category_id" validate:"required,uuid4"`
	Name        string                `form:"name" validate:"required"`
	Description string                `form:"description"`
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
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Category    *MaterialCategoryResponse `json:"category"`
}
