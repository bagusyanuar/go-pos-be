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

type ProductCategoryHandler struct {
	ProductCategoryService domain.ProductCategoryService
	Config                 *config.AppConfig
}

func NewProductCategoryHandler(
	productCategoryService domain.ProductCategoryService,
	config *config.AppConfig,
) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		ProductCategoryService: productCategoryService,
		Config:                 config,
	}
}

func (h *ProductCategoryHandler) FindAll(ctx *fiber.Ctx) error {

	queryParams := new(schema.ProductCategoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidQueryParameters.Error(),
			"error":   err.Error(),
		})
	}

	data, pagination, err := h.ProductCategoryService.FindAll(ctx.UserContext(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully get product categories",
		"data":    data,
		"meta":    pagination,
	})
}

func (h *ProductCategoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := h.ProductCategoryService.FindByID(ctx.UserContext(), id)
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
		"message": "successfully get product category",
		"data":    data,
	})
}

func (h *ProductCategoryHandler) Create(ctx *fiber.Ctx) error {
	req := new(schema.ProductCategoryRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(h.Config.Validator, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	err = h.ProductCategoryService.Create(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "successfully create new product category",
	})
}
