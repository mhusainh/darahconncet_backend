package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type HospitalRepository interface {
	Create(ctx context.Context, hospital *entity.Hospital) error
	GetById(ctx context.Context, id int64) (*entity.Hospital, error)
	GetAll(ctx context.Context, req dto.GetAllHospitalRequest) ([]entity.Hospital, int64, error)
	Update(ctx context.Context, hospital *entity.Hospital) error
	Delete(ctx context.Context, hospital *entity.Hospital) error
}

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) HospitalRepository {
	return &hospitalRepository{db}
}

// applyFilters menerapkan filter, sorting, dan pagination ke query GORM
func (r *hospitalRepository) applyFilters(query *gorm.DB, req dto.GetAllHospitalRequest) (*gorm.DB, dto.GetAllHospitalRequest) {
	// Filter berdasarkan Search (pada judul atau pesan)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Where("LOWER(name) LIKE ?", "%"+search+"%").
			Or("LOWER(address) LIKE ?", "%"+search+"%").
			Or("LOWER(city) LIKE ?", "%"+search+"%").
			Or("LOWER(province) LIKE ?", "%"+search+"%").
			Or("LOWER(latitude) LIKE ?", "%"+search+"%").
			Or("LOWER(longitude) LIKE ?", "%"+search+"%")
	}

	// Set default values jika tidak ada
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	// Sorting
	sortBy := "created_at"
	if req.Sort != "" {
		sortBy = req.Sort
	}

	orderBy := "desc"
	if req.Order != "" {
		orderBy = req.Order
	}

	query = query.Order(sortBy + " " + orderBy)

	return query, req
}

func (r *hospitalRepository) Create(ctx context.Context, hospital *entity.Hospital) error {
	return r.db.WithContext(ctx).Create(&hospital).Error
}

func (r *hospitalRepository) GetById(ctx context.Context, id int64) (*entity.Hospital, error) {
	result := new(entity.Hospital)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *hospitalRepository) GetAll(ctx context.Context, req dto.GetAllHospitalRequest) ([]entity.Hospital, int64, error) {
	var hospital []entity.Hospital
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Hospital{})
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&hospital).Error; err != nil {
		return nil, 0, err
	}

	return hospital, total, nil
}

func (r *hospitalRepository) Update(ctx context.Context, hospital *entity.Hospital) error {
	return r.db.WithContext(ctx).Model(hospital).Updates(hospital).Error
}

func (r *hospitalRepository) Delete(ctx context.Context, hospital *entity.Hospital) error {
	return r.db.WithContext(ctx).Delete(hospital).Error
}
