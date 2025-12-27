package schema

type ProductCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductCategoryQuery struct {
	Param string `json:"param" query:"param"`
}
