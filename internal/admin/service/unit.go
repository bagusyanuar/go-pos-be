package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
)

type unitServiceImpl struct {
	UnitRepository domain.UnitRepository
	Config         *config.AppConfig
}

// Create implements domain.UnitService.
func (u *unitServiceImpl) Create(ctx context.Context, schema *schema.UnitRequest) error {
	e := entity.Unit{
		Name: schema.Name,
	}

	_, err := u.UnitRepository.Create(ctx, &e)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.UnitService.
func (u *unitServiceImpl) Delete(ctx context.Context, id string) error {
	_, err := u.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = u.UnitRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Find implements domain.UnitService.
func (u *unitServiceImpl) Find(ctx context.Context, queryParams *schema.UnitQuery) ([]schema.UnitResponse, *util.PaginationMeta, error) {
	data, pagination, err := u.UnitRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.UnitResponse{}, nil, err
	}

	res := schema.ToUnits(data)
	return res, pagination, nil
}

// FindByID implements domain.UnitService.
func (u *unitServiceImpl) FindByID(ctx context.Context, id string) (*schema.UnitResponse, error) {
	data, err := u.UnitRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := schema.ToUnit(data)
	return res, nil
}

// Update implements domain.UnitService.
func (u *unitServiceImpl) Update(ctx context.Context, id string, schema *schema.UnitRequest) error {
	unit, err := u.UnitRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	unit.Name = schema.Name

	_, err = u.UnitRepository.Update(ctx, unit)
	if err != nil {
		return err
	}
	return nil
}

func NewUnitService(
	unitRepository domain.UnitRepository,
	config *config.AppConfig,
) domain.UnitService {
	return &unitServiceImpl{
		UnitRepository: unitRepository,
		Config:         config,
	}
}
