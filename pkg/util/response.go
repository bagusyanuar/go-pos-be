package util

type APIResponse[T any] struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    T               `json:"data"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
	Errors  any             `json:"errors,omitempty"`
}
