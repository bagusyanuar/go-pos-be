package entity

import "github.com/google/uuid"

type MaterialUnit struct {
	MaterialID     uuid.UUID `gorm:"primaryKey"`
	UnitID         uuid.UUID `gorm:"primaryKey"`
	ConversionRate float64
	IsDefault      bool
	Material       Material `gorm:"foreignKey:MaterialID"`
	Unit           Unit     `gorm:"foreignKey:UnitID"`
}

func (e *MaterialUnit) TableName() string {
	return "material_units"
}
