package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplierContact struct {
	ID         uuid.UUID
	SupplierID *uuid.UUID
	Type       string
	Value      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Supplier   *Material `gorm:"foreignKey:SupplierID"`
}

func (e *SupplierContact) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *SupplierContact) TableName() string {
	return "supplier_contacts"
}
