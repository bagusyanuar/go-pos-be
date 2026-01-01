package service

import (
	"context"
	"errors"
	"time"

	"github.com/bagusyanuar/go-pos-be/internal/auth/domain"
	"github.com/bagusyanuar/go-pos-be/internal/auth/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	UserRepository domain.UserRepository
	Config         *config.AppConfig
}

// Login implements domain.AuthService.
func (a *authServiceImpl) Login(
	ctx context.Context,
	schema *schema.LoginRequest,
) (
	accessToken string,
	refreshToken string,
	err error,
) {
	email := schema.Email
	password := schema.Password
	user, err := a.UserRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", exception.ErrPasswordMissmatch
		}
		return "", "", err
	}

	accessToken, err = a.generateToken(
		user,
		time.Minute*time.Duration(a.Config.JWT.Expiration),
		a.Config.JWT.Secret,
		false,
	)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = a.generateToken(
		user,
		time.Minute*time.Duration(a.Config.JWT.ExpirationRefresh),
		a.Config.JWT.SecretRefresh,
		true,
	)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authServiceImpl) generateToken(
	user *entity.User,
	duration time.Duration,
	secret string,
	isRefresh bool,
) (string, error) {
	exp := time.Now().Add(duration)

	var claims jwt.Claims
	if isRefresh {
		claims = jwt.RegisteredClaims{
			Issuer:    a.Config.JWT.Issuer,
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   user.ID.String(),
		}
	} else {
		claims = util.JWTAppClaims{
			Email: user.Email,
			Roles: []string{"admin"},
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    a.Config.JWT.Issuer,
				ExpiresAt: jwt.NewNumericDate(exp),
				Subject:   user.ID.String(),
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func NewAuthService(
	userRepository domain.UserRepository,
	config *config.AppConfig,
) domain.AuthService {
	return &authServiceImpl{
		UserRepository: userRepository,
		Config:         config,
	}
}
