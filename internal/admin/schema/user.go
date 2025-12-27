package schema

type UserRequest struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserQuery struct {
	Param string `json:"param" query:"param"`
}
