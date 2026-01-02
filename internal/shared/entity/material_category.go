package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialCategory struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (e *MaterialCategory) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *MaterialCategory) TableName() string {
	return "material_categories"
}
