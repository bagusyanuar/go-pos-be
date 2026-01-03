package schema

import (
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

/* === Request & Query Unit Category === */

type UnitRequest struct {
	Name string `json:"name" validate:"required"`
}

type UnitQuery struct {
	Param string `json:"param" query:"param"`
	util.QueryPagination
	util.QuerySort
}

/* === Response Unit Category === */

type UnitResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToUnit(unit *entity.Unit) *UnitResponse {
	if unit == nil {
		return nil
	}
	return &UnitResponse{
		ID:   unit.ID.String(),
		Name: unit.Name,
	}
}

func ToUnits(units []entity.Unit) []UnitResponse {
	responses := make([]UnitResponse, 0, len(units))
	for _, v := range units {
		res := ToUnit(&v)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
