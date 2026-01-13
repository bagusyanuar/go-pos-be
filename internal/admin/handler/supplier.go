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
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusBadRequest,
				Message: exception.ErrInvalidRequestBody.Error(),
			})
	}

	messages, err := util.Validate(s.Config.Validator, request)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusUnprocessableEntity,
				Message: exception.ErrValidation.Error(),
				Errors:  messages,
			})
	}

	supplier, err := s.SupplierService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(util.APIResponse[*schema.SupplierCreateResponse]{
			Code:    fiber.StatusInternalServerError,
			Message: "successfully create new supplier",
			Data:    supplier,
		})
}

// Delete implements SupplierHandler.
func (s *supplierHandlerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := s.SupplierService.Delete(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(util.APIResponse[any]{
				Code:    fiber.StatusNotFound,
				Message: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(util.APIResponse[any]{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(util.APIResponse[any]{
		Code:    fiber.StatusOK,
		Message: "successfully delete supplier",
	})
}

// Find implements SupplierHandler.
func (s *supplierHandlerImpl) Find(ctx *fiber.Ctx) error {
	queryParams := new(schema.SupplierQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.APIResponse[any]{
			Code:    fiber.StatusBadRequest,
			Message: exception.ErrInvalidQueryParameters.Error(),
		})
	}

	data, totalItems, err := s.SupplierService.Find(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(util.APIResponse[any]{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	suppliers := mapper.ToSuppliers(data)
	pagination := util.MakePagination(queryParams.Page, queryParams.PageSize, totalItems)

	return ctx.Status(fiber.StatusOK).JSON(util.APIResponse[*[]schema.SupplierResponse]{
		Code:    fiber.StatusBadRequest,
		Message: "successfully get suppliers",
		Data:    &suppliers,
		Meta:    &pagination,
	})
}

// FindByID implements SupplierHandler.
func (s *supplierHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := s.SupplierService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, exception.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(util.APIResponse[any]{
				Code:    fiber.StatusNotFound,
				Message: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(util.APIResponse[any]{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(util.APIResponse[any]{
		Code:    fiber.StatusOK,
		Message: "successfully get suppliers",
		Data:    data,
	})
}

// Update implements SupplierHandler.
func (s *supplierHandlerImpl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(schema.SupplierRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusBadRequest,
				Message: exception.ErrInvalidRequestBody.Error(),
			})
	}

	messages, err := util.Validate(s.Config.Validator, request)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusUnprocessableEntity,
				Message: exception.ErrValidation.Error(),
				Errors:  messages,
			})
	}

	err = s.SupplierService.Update(ctx.UserContext(), id, request)
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(util.APIResponse[any]{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(util.APIResponse[*schema.SupplierCreateResponse]{
			Code:    fiber.StatusInternalServerError,
			Message: "successfully update supplier",
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
