package domain

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
