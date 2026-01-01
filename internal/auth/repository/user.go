package repository

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/auth/domain"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

// FindByEmail implements domain.UserRepository.
func (u *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	tx := u.DB.WithContext(ctx)

	user := new(entity.User)

	if err := tx.Where("email = ?", email).
		First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}
