package handler

import (
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/gofiber/fiber/v2"
)

type (
	MaterialInventoryHandler interface {
		Find(ctx *fiber.Ctx) error
		FindByID(ctx *fiber.Ctx) error
	}

	materialInventoryHandlerImpl struct {
		MaterialInventoryService domain.MaterialInventoryService
		Config                   *config.AppConfig
	}
)

// Find implements MaterialInventoryHandler.
func (m *materialInventoryHandlerImpl) Find(ctx *fiber.Ctx) error {
	queryParams := new(schema.MaterialInventoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidQueryParameters.Error(),
			"error":   err.Error(),
		})
	}

	data, pagination, err := m.MaterialInventoryService.Find(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get material inventories",
		"data":    data,
		"meta":    pagination,
	})
}

// FindByID implements MaterialInventoryHandler.
func (m *materialInventoryHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := m.MaterialInventoryService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get material inventory",
		"data":    data,
	})
}

func NewMaterialInventoryHandler(
	materialInventoryService domain.MaterialInventoryService,
	config *config.AppConfig,
) MaterialInventoryHandler {
	return &materialInventoryHandlerImpl{
		MaterialInventoryService: materialInventoryService,
		Config:                   config,
	}
}
