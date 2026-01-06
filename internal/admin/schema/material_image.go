package schema

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type MaterialImageRequest struct {
	MaterialID uuid.UUID             `form:"material_id" validate:"required,uuid4"`
	Image      *multipart.FileHeader `form:"image" validate:"required"`
}
