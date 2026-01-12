package repositories

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	userRepositoryImpl struct {
		DB *gorm.DB
	}
)

// FindByEmail implements UserRepository.
func (u *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}
