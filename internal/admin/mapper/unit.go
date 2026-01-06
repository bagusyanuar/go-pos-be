package mapper

import (
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
)

func ToUnit(unit *entity.Unit) *schema.UnitResponse {
	if unit == nil {
		return nil
	}
	return &schema.UnitResponse{
		ID:        unit.ID.String(),
		Name:      unit.Name,
		CreatedAt: unit.CreatedAt.Format(constant.BaseDateTimeLayout),
		UpdatedAt: unit.UpdatedAt.Format(constant.BaseDateTimeLayout),
	}
}

func ToUnits(units []entity.Unit) []schema.UnitResponse {
	responses := make([]schema.UnitResponse, 0, len(units))
	for _, v := range units {
		res := ToUnit(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
