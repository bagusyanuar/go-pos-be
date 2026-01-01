package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/auth/schema"
)

type AuthService interface {
	Login(ctx context.Context, schema *schema.LoginRequest) (accessToken, refreshToken string, err error)
}
