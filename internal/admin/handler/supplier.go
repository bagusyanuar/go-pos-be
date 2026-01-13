package handler

import (
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type (
	SupplierHandler interface {
		Find(ctx *fiber.Ctx) error
		FindByID(ctx *fiber.Ctx) error
		Create(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	supplierHandlerImpl struct {
		SupplierService domain.SupplierService
		Config          *config.AppConfig
	}
)

// Create implements SupplierHandler.
func (s *supplierHandlerImpl) Create(ctx *fiber.Ctx) error {
	request := new(schema.SupplierRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(s.Config.Validator, request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	supplier, err := s.SupplierService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new supplier",
		"data":    supplier,
	})
}

// Delete implements SupplierHandler.
func (s *supplierHandlerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := s.SupplierService.Delete(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully delete supplier",
	})
}

// Find implements SupplierHandler.
func (s *supplierHandlerImpl) Find(ctx *fiber.Ctx) error {
	queryParams := new(schema.SupplierQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidQueryParameters.Error(),
			"error":   err.Error(),
		})
	}

	data, totalItems, err := s.SupplierService.Find(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	suppliers := mapper.ToSuppliers(data)
	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get suppliers",
		"data":    suppliers,
		"meta":    pagination,
	})
}

// FindByID implements SupplierHandler.
func (s *supplierHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := s.SupplierService.FindByID(ctx.UserContext(), id)
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
		"message": "successfully get supplier",
		"data":    data,
	})
}

// Update implements SupplierHandler.
func (s *supplierHandlerImpl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(schema.SupplierRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(s.Config.Validator, request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	err = s.SupplierService.Update(ctx.UserContext(), id, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully update supplier",
	})
}

func NewSupplierHandler(
	supplierService domain.SupplierService,
	config *config.AppConfig,
) SupplierHandler {
	return &supplierHandlerImpl{
		SupplierService: supplierService,
		Config:          config,
	}
}
