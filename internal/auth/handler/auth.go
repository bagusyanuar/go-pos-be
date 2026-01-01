package handler

import (
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/auth/domain"
	"github.com/bagusyanuar/go-pos-be/internal/auth/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandler interface {
		Login(ctx *fiber.Ctx) error
	}

	authHandlerImpl struct {
		AuthService domain.AuthService
		Config      *config.AppConfig
	}
)

// Login implements AuthHandler.
func (a *authHandlerImpl) Login(ctx *fiber.Ctx) error {
	req := new(schema.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": exception.ErrInvalidRequestBody.Error(),
		})
	}

	messages, err := util.Validate(a.Config.Validator, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": exception.ErrValidation.Error(),
			"errors":  messages,
		})
	}

	accessToken, refreshToken, err := a.AuthService.Login(ctx.UserContext(), req)
	if err != nil {
		if errors.Is(err, exception.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": err.Error(),
			})
		}

		if errors.Is(err, exception.ErrPasswordMissmatch) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	res := schema.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully login",
		"data":    res,
	})

}

func NewAuthHandler(
	authService domain.AuthService,
	config *config.AppConfig,
) AuthHandler {
	return &authHandlerImpl{
		AuthService: authService,
		Config:      config,
	}
}
