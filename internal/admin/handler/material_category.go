package handler

import (
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type (
	MaterialCategoryHandler interface {
		Find(ctx *fiber.Ctx) error
		FindByID(ctx *fiber.Ctx) error
		Create(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	materialCategoryHandlerImpl struct {
		MaterialCategoryService domain.MaterialCategoryService
		Config                  *config.AppConfig
	}
)

// Create implements MaterialCategoryHandler.
func (m *materialCategoryHandlerImpl) Create(ctx *fiber.Ctx) error {
	req := new(schema.MaterialCategoryRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(m.Config.Validator, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	err = m.MaterialCategoryService.Create(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new material category",
	})
}

// Delete implements MaterialCategoryHandler.
func (m *materialCategoryHandlerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := m.MaterialCategoryService.Delete(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully delete material category",
	})
}

// Find implements MaterialCategoryHandler.
func (m *materialCategoryHandlerImpl) Find(ctx *fiber.Ctx) error {
	queryParams := new(schema.MaterialCategoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidQueryParameters.Error(),
			"error":   err.Error(),
		})
	}

	data, pagination, err := m.MaterialCategoryService.Find(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get material categories",
		"data":    data,
		"meta":    pagination,
	})
}

// FindByID implements MaterialCategoryHandler.
func (m *materialCategoryHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := m.MaterialCategoryService.FindByID(ctx.UserContext(), id)
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
		"message": "successfully get material category",
		"data":    data,
	})
}

// Update implements MaterialCategoryHandler.
func (m *materialCategoryHandlerImpl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := new(schema.MaterialCategoryRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(m.Config.Validator, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	err = m.MaterialCategoryService.Update(ctx.UserContext(), id, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully update material category",
	})
}

func NewMaterialCategoryHandler(
	materialCategoryService domain.MaterialCategoryService,
	config *config.AppConfig,
) MaterialCategoryHandler {
	return &materialCategoryHandlerImpl{
		MaterialCategoryService: materialCategoryService,
		Config:                  config,
	}
}
