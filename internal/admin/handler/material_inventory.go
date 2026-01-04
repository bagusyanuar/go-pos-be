package handler

import (
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
		Create(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	materialInventoryHandlerImpl struct {
		MaterialInventoryService domain.MaterialInventoryService
		Config                   *config.AppConfig
	}
)

// Create implements MaterialInventoryHandler.
func (m *materialInventoryHandlerImpl) Create(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements MaterialInventoryHandler.
func (m *materialInventoryHandlerImpl) Delete(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

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
	panic("unimplemented")
}

// Update implements MaterialInventoryHandler.
func (m *materialInventoryHandlerImpl) Update(ctx *fiber.Ctx) error {
	panic("unimplemented")
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
