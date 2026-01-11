package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MaterialInventory struct {
	ID         uuid.UUID
	MaterialID *uuid.UUID
	Quantity   decimal.Decimal `gorm:"type:numeric(15,3);default:0;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Material   *Material `gorm:"foreignKey:MaterialID"`
}

func (e *MaterialInventory) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *MaterialInventory) TableName() string {
	return "material_inventories"
}
