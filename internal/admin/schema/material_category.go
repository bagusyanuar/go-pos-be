package schema

import (
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type MaterialCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type MaterialCategoryQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

type MaterialCategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
