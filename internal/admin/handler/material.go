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
	MaterialHandler interface {
		Find(ctx *fiber.Ctx) error
		FindByID(ctx *fiber.Ctx) error
		Create(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
		UploadImage(ctx *fiber.Ctx) error
	}

	materialHandlerImpl struct {
		MaterialService domain.MaterialService
		Config          *config.AppConfig
	}
)

// Create implements MaterialHandler.
func (m *materialHandlerImpl) Create(ctx *fiber.Ctx) error {
	req := new(schema.MaterialRequest)
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

	err = m.MaterialService.Create(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new material",
	})
}

// Delete implements MaterialHandler.
func (m *materialHandlerImpl) Delete(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// Find implements MaterialHandler.
func (m *materialHandlerImpl) Find(ctx *fiber.Ctx) error {
	queryParams := new(schema.MaterialQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidQueryParameters.Error(),
			"error":   err.Error(),
		})
	}

	data, pagination, err := m.MaterialService.Find(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get materials",
		"data":    data,
		"meta":    pagination,
	})
}

// FindByID implements MaterialHandler.
func (m *materialHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := m.MaterialService.FindByID(ctx.UserContext(), id)
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
		"message": "successfully get material",
		"data":    data,
	})
}

// Update implements MaterialHandler.
func (m *materialHandlerImpl) Update(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UploadImage implements MaterialHandler.
func (m *materialHandlerImpl) UploadImage(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := new(schema.MaterialImageRequest)
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

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "file not found",
		})
	}

	req.Image = fileHeader

	schema := req

	err = m.MaterialService.UploadImage(ctx.UserContext(), id, schema)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new material images",
	})
}

func NewMaterialHandler(
	materialService domain.MaterialService,
	config *config.AppConfig,
) MaterialHandler {
	return &materialHandlerImpl{
		MaterialService: materialService,
		Config:          config,
	}
}
