package schema

import (
	"mime/multipart"
)

type MaterialImageRequest struct {
	Image *multipart.FileHeader `form:"image"`
}
