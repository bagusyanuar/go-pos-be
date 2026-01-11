package service

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-pos-be/internal/admin/domain"
	"github.com/bagusyanuar/go-pos-be/internal/admin/mapper"
	"github.com/bagusyanuar/go-pos-be/internal/admin/schema"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
	"github.com/bagusyanuar/go-pos-be/internal/shared/constant"
	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"github.com/bagusyanuar/go-pos-be/pkg/exception"
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/shopspring/decimal"
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
func (m *materialServiceImpl) Create(ctx context.Context, schema *schema.MaterialRequest) (*schema.MaterialCreateResponse, error) {

	e := entity.Material{
		MaterialCategoryID: schema.CategoryID,
		Name:               schema.Name,
		Description:        schema.Description,
	}

	material, err := m.MaterialRepository.Create(ctx, &e)
	if err != nil {
		return nil, err
	}
	res := mapper.ToMaterialCreate(material)
	return res, nil
}

// Update implements domain.MaterialService.
func (m *materialServiceImpl) Update(ctx context.Context, id string, schema *schema.MaterialRequest) error {
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

// ManageUnit implements domain.MaterialService.
func (m *materialServiceImpl) ManageUnit(ctx context.Context, id string, schema *schema.MaterialUnitRequest) error {
	// 1. Ambil data bahan baku & validasi berdasarkan id
	material, err := m.MaterialRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// --- 2. Validasi Umum dan Pemetaan ---

	newUnits := make([]entity.MaterialUnit, 0, len(schema.Units))
	unitDefaultCountInRequest := 0
	var unitIDOldDefault uuid.UUID // Untuk menyimpan ID unit default lama

	// Cari ID unit default lama (jika ada)
	for _, eU := range material.Units {
		if eU.IsDefault {
			unitIDOldDefault = eU.UnitID
			break
		}
	}

	// Validasi data di Request dan hitung default baru
	for _, u := range schema.Units {
		if u.IsDefault == nil {
			return exception.ErrUnitDefaultValue
		}

		if *u.IsDefault {
			unitDefaultCountInRequest++
			// Di mode Calibrate, Rate boleh berapapun. Di mode Create/Append, wajib 1.
			// Namun, karena ini adalah data BARU yang akan jadi default, ConversionRate-nya harus 1
			if u.ConversionRate != 1 && schema.Type != constant.TypeCalibrate {
				return exception.ErrUnitConversionRate
			}
		}

		// Mapping ke Entity (diperlukan untuk semua tipe)
		newUnits = append(newUnits, entity.MaterialUnit{
			MaterialID:     material.ID,
			UnitID:         u.UnitID,
			ConversionRate: u.ConversionRate,
			IsDefault:      *u.IsDefault,
		})
	}

	// --- 3. Logika Flow Control berdasarkan ActionType ---

	switch schema.Type {
	case constant.TypeCreate:
		// Cek: Create awal WAJIB ada tepat 1 default
		if unitDefaultCountInRequest != 1 {
			return exception.ErrUnitDefault
		}
		// Eksekusi Repository: Hapus semua unit lama lalu ganti dengan unit baru
		return m.MaterialRepository.AppendUnit(ctx, material, newUnits)

	case constant.TypeAppend:
		// Cek: Append DILARANG mengirim unit default baru (karena ada endpoint Calibrate)
		if unitDefaultCountInRequest > 0 {
			return errors.New("penggantian unit default harus melalui aksi 'calibrate'")
		}
		// Eksekusi Repository: Tambahkan unit baru (Upsert)
		return m.MaterialRepository.AppendUnit(ctx, material, newUnits)

	case constant.TypeCalibrate:
		// Cek: Calibrate WAJIB mengirim semua unit (existing + new) dan harus ada tepat 1 default
		if unitDefaultCountInRequest != 1 {
			return exception.ErrUnitDefault // Harus ada 1 unit default
		}

		// Logika Bisnis: Hitung Konversi Stok

		// Cari rate unit default LAMA di dalam data input BARU
		var rateOfOldUnitInNewBasis decimal.Decimal
		foundOldUnit := false
		for _, u := range newUnits {
			if u.UnitID == unitIDOldDefault {
				rateOfOldUnitInNewBasis = decimal.NewFromFloat(u.ConversionRate)
				foundOldUnit = true
				break
			}
		}

		if !foundOldUnit {
			return errors.New("unit default lama harus disertakan dalam daftar kalibrasi")
		}

		// Hitung Stok Baru
		var inventoryUpdate map[string]any
		if material.Inventory != nil {
			// StokBaru = StokLama * RateBaru Unit Lama terhadap Basis Baru
			newQuantity := material.Inventory.Quantity.Mul(rateOfOldUnitInNewBasis).Round(3)
			inventoryUpdate = map[string]any{
				"quantity": newQuantity,
				// Tambahkan field update_at agar ada timestamp perubahan stok
			}
		} else {
			// Jika material tidak punya inventory, anggap quantity 0
			inventoryUpdate = map[string]any{"quantity": decimal.Zero}
		}

		// Eksekusi Repository: Update Stok dan Unit Kalibrasi
		// Kita gunakan fungsi repository yang sudah kita buat sebelumnya
		return m.MaterialRepository.CalibrateUnit(ctx, material.ID.String(), inventoryUpdate, newUnits)

	default:
		return errors.New("tipe aksi tidak valid")
	}
}

// DeleteUnit implements domain.MaterialService.
func (m *materialServiceImpl) DeleteUnit(ctx context.Context, materialID string, unitID string) error {
	err := m.MaterialRepository.DeleteUnit(ctx, materialID, unitID)
	if err != nil {
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
