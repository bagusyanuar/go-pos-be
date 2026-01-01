package provider

import (
	"github.com/bagusyanuar/go-pos-be/internal/auth/domain"
	"github.com/bagusyanuar/go-pos-be/internal/auth/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User domain.UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: repository.NewUserRepository(db),
	}
}
