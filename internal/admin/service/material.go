package service

import (
	"context"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/minio/minio-go/v7"
)

type materialServiceImpl struct {
	MaterialRepository domain.MaterialRepository
	Config             *config.AppConfig
}

// Find implements domain.MaterialService.
func (m *materialServiceImpl) Find(ctx context.Context, queryParams *schema.MaterialQuery) ([]schema.MaterialResponse, *util.PaginationMeta, error) {
	data, pagination, err := m.MaterialRepository.Find(ctx, queryParams)
	if err != nil {
		return []schema.MaterialResponse{}, nil, err
	}

	res := mapper.ToMaterials(data)
	return res, pagination, nil
}

// FindByID implements domain.MaterialService.
func (m *materialServiceImpl) FindByID(ctx context.Context, id string) (*schema.MaterialResponse, error) {
	data, err := m.MaterialRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := mapper.ToMaterial(data)
	return res, nil
}

// Create implements domain.MaterialService.
func (m *materialServiceImpl) Create(ctx context.Context, schema *schema.MaterialRequest) error {

	unitDefaultCount := 0

	for _, u := range schema.Units {
		if u.IsDefault {
			unitDefaultCount++
			if u.ConversionRate != 1 {
				return exception.ErrUnitConversionRate
			}
		}
	}

	if unitDefaultCount != 1 {
		return exception.ErrUnitDefault
	}

	units := make([]entity.MaterialUnit, 0, len(schema.Units))
	for _, v := range schema.Units {
		unit := entity.MaterialUnit{
			UnitID:         v.UnitID,
			ConversionRate: v.ConversionRate,
			IsDefault:      v.IsDefault,
		}
		units = append(units, unit)
	}

	e := entity.Material{
		MaterialCategoryID: schema.CategoryID,
		Name:               schema.Name,
		Description:        schema.Description,
		Units:              units,
	}

	_, err := m.MaterialRepository.Create(ctx, &e)
	if err != nil {
		return err
	}
	return nil
}

// Update implements domain.MaterialService.
func (m *materialServiceImpl) Update(ctx context.Context, id string, schema *schema.MaterialUpdateRequest) error {
	material, err := m.MaterialRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	material.MaterialCategoryID = schema.CategoryID
	material.Name = schema.Name
	material.Description = schema.Description

	_, err = m.MaterialRepository.Update(ctx, material)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.MaterialService.
func (m *materialServiceImpl) Delete(ctx context.Context, id string) error {
	_, err := m.MaterialRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = m.MaterialRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// UploadImage implements domain.MaterialService.
func (m *materialServiceImpl) UploadImage(ctx context.Context, id string, schema *schema.MaterialImageRequest) error {
	if schema.Image == nil {
		return exception.ErrNoFileAttched
	}

	material, err := m.MaterialRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	minioObj := util.MinioObject{
		Context:    ctx,
		Client:     m.Config.Minio.MinioClient,
		Bucket:     m.Config.Minio.Bucket,
		Path:       "material_categories",
		FileHeader: schema.Image,
	}

	// Upload & Resize secara Concurrent
	groupID, uploadResults, err := minioObj.UploadImageWithThumbnail()
	if err != nil {
		return err
	}

	// Preparing batch data
	images := make([]entity.MaterialImage, 0, len(uploadResults))
	for _, result := range uploadResults {
		img := entity.MaterialImage{
			MaterialID:   &material.ID,
			ImageGroupID: groupID,
			Type:         result.Type,
			URL:          result.ObjectName,
		}
		images = append(images, img)
	}

	// Simpan ke database
	err = m.MaterialRepository.UploadImage(ctx, images)
	if err != nil {

		// CleanUp: Hapus semua file yang baru diupload jika DB gagal
		for _, res := range uploadResults {
			_ = m.Config.Minio.MinioClient.RemoveObject(
				ctx,
				m.Config.Minio.Bucket,
				res.ObjectName,
				minio.RemoveObjectOptions{},
			)
		}
		return err
	}

	return nil
}

func NewMaterialService(
	materialRepository domain.MaterialRepository,
	config *config.AppConfig,
) domain.MaterialService {
	return &materialServiceImpl{
		MaterialRepository: materialRepository,
		Config:             config,
	}
}
