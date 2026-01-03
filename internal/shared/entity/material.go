package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Material struct {
	ID                 uuid.UUID
	MaterialCategoryID *uuid.UUID
	Name               string
	Description        *string
	Image              *string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
	MaterialCategory   *MaterialCategory `gorm:"foreignKey:MaterialCategoryID"`
}

func (e *Material) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *Material) TableName() string {
	return "materials"
}
