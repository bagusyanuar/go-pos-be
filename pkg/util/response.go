package util

import "github.com/gofiber/fiber/v2"

type APIResponse[T any] struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    T               `json:"data,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
	Errors  any             `json:"errors,omitempty"`
}

func SuccessResponse[T any](ctx *fiber.Ctx, message string, data T) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": message,
		"data":    data,
	})
}
