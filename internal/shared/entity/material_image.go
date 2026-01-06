package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialImage struct {
	ID           uuid.UUID
	MaterialID   *uuid.UUID
	ImageGroupID string
	Type         string
	URL          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Material     *Material `gorm:"foreignKey:MaterialID"`
}

func (e *MaterialImage) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *MaterialImage) TableName() string {
	return "material_images"
}
