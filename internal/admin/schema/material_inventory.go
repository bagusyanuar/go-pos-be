package schema

import (
	"sort"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/shopspring/decimal"
)

type MaterialInventoryQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

type MaterialInventoryResponse struct {
	MaterialResponse
	Quantity float64 `json:"quantity"`
}

func ToMaterialInventory(material *entity.Material) *MaterialInventoryResponse {
	if material == nil {
		return nil
	}

	category := new(MaterialMaterialCategory)
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

	units := make([]MaterialUnit, 0, len(material.Units))

	sort.Slice(material.Units, func(i, j int) bool {
		return material.Units[i].ConversionRate > material.Units[j].ConversionRate
	})

	remaining := globalStok
	for _, v := range material.Units {

		if v.ConversionRate <= 0 {
			continue
		}

		rate := decimal.NewFromFloat(v.ConversionRate)
		qty := remaining.Div(rate).Floor()

		qtyFloat, _ := qty.Float64()
		unit := MaterialUnit{
			UnitID:         v.UnitID.String(),
			Name:           v.Unit.Name,
			ConversionRate: v.ConversionRate,
			IsDefault:      v.IsDefault,
			Quantity:       qtyFloat,
		}
		units = append(units, unit)

		remaining = remaining.Sub(qty.Mul(rate))
	}

	return &MaterialInventoryResponse{
		MaterialResponse: MaterialResponse{
			ID:          material.ID.String(),
			Name:        material.Name,
			Description: material.Description,
			Category:    category,
			Units:       units,
		},
		Quantity: quantity,
	}
}

func ToMaterialInventories(materials []entity.Material) []MaterialInventoryResponse {
	responses := make([]MaterialInventoryResponse, 0, len(materials))
	for _, v := range materials {
		res := ToMaterialInventory(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
