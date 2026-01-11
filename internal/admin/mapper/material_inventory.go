package mapper

import (
	"sort"

	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/shopspring/decimal"
)

func ToMaterialInventory(material *entity.Material) *schema.MaterialInventoryResponse {
	if material == nil {
		return nil
	}

	category := new(schema.MaterialInventoryMaterialCategory)
	if material.MaterialCategory != nil {
		category.ID = material.MaterialCategory.ID.String()
		category.Name = material.MaterialCategory.Name
	}

	var globalStok decimal.Decimal
	if material.Inventory != nil {
		globalStok = material.Inventory.Quantity
	} else {
		globalStok = decimal.Zero
	}

	var quantity float64
	if material.Inventory != nil {
		quantity, _ = material.Inventory.Quantity.Float64()
	}

	units := make([]schema.MaterialInventoryUnit, 0, len(material.Units))

	sort.Slice(material.Units, func(i, j int) bool {
		return material.Units[i].ConversionRate > material.Units[j].ConversionRate
	})

	remaining := globalStok
	for i, v := range material.Units {

		if v.ConversionRate <= 0 {
			continue
		}

		rate := decimal.NewFromFloat(v.ConversionRate)

		// do flooring if not last conversion rate
		var qty decimal.Decimal
		if i == len(material.Units)-1 {
			qty = remaining.Div(rate)
		} else {
			qty = remaining.Div(rate).Floor()
		}

		qtyFloat, _ := qty.Float64()
		unit := schema.MaterialInventoryUnit{
			UnitID:         v.UnitID.String(),
			Name:           v.Unit.Name,
			ConversionRate: v.ConversionRate,
			IsDefault:      v.IsDefault,
			Quantity:       qtyFloat,
		}
		units = append(units, unit)

		remaining = remaining.Sub(qty.Mul(rate))
	}

	return &schema.MaterialInventoryResponse{
		ID:          material.ID.String(),
		Name:        material.Name,
		Description: material.Description,
		Quantity:    quantity,
		CreatedAt:   material.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt:   material.UpdatedAt.Format(constant.BaseDateTimeLayout),
		Category:    category,
		Units:       units,
	}
}

func ToMaterialInventories(materials []entity.Material) []schema.MaterialInventoryResponse {
	responses := make([]schema.MaterialInventoryResponse, 0, len(materials))
	for _, v := range materials {
		res := ToMaterialInventory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
