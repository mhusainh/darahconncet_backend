package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type HealthPassportRepository interface {
	Create(ctx context.Context, healthPassport *entity.HealthPassport) error
	GetById(ctx context.Context, id int64) (*entity.HealthPassport, error)
	GetAll(ctx context.Context, req dto.GetAllHealthPassportRequest) ([]entity.HealthPassport, int64, error)
	GetByUserId(ctx context.Context, userId int64) (*entity.HealthPassport, error)
	Update(ctx context.Context, healthPassport *entity.HealthPassport) error
	Delete(ctx context.Context, healthPassport *entity.HealthPassport) error
}

type healthPassportRepository struct {
	db *gorm.DB
}

func NewHealthPassportRepository(db *gorm.DB) HealthPassportRepository {
	return &healthPassportRepository{db}
}

func (r *healthPassportRepository) applyFilters(query *gorm.DB, req dto.GetAllHealthPassportRequest) (*gorm.DB, dto.GetAllHealthPassportRequest) {
	// Filter berdasarkan IsRead
	if req.Status != "" {
		query = query.Where("LOWER(status) = ?", req.Status)
	}

	// Filter berdasarkan Search (pada judul atau pesan)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN users ON users.id = health_passports.user_id").
			Where("LOWER(health_passports.passport_number) LIKE ? OR LOWER(users.name) LIKE ?",
				"%"+search+"%", "%"+search+"%")
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

func (r *healthPassportRepository) Create(ctx context.Context, healthPassport *entity.HealthPassport) error {
	return r.db.WithContext(ctx).Create(healthPassport).Error
}

func (r *healthPassportRepository) GetById(ctx context.Context, id int64) (*entity.HealthPassport, error) {
	result := new(entity.HealthPassport)
	if err := r.db.WithContext(ctx).Where("id = ?", id).Preload("User").First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *healthPassportRepository) GetAll(ctx context.Context, req dto.GetAllHealthPassportRequest) ([]entity.HealthPassport, int64, error) {
	var healthPassports []entity.HealthPassport
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.HealthPassport{}).Preload("User")
	query, req = r.applyFilters(query, req)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (req.Page - 1) * req.Limit
	query = query.Offset(int(offset)).Limit(int(req.Limit))

	// Execute query
	if err := query.Find(&healthPassports).Error; err != nil {
		return nil, 0, err
	}

	return healthPassports, total, nil
}

func (r *healthPassportRepository) GetByUserId(ctx context.Context, userId int64) (*entity.HealthPassport, error) {
	result := new(entity.HealthPassport)
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Preload("User").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *healthPassportRepository) Update(ctx context.Context, healthPassport *entity.HealthPassport) error {
	return r.db.WithContext(ctx).Model(healthPassport).Updates(healthPassport).Error
}

func (r *healthPassportRepository) Delete(ctx context.Context, healthPassport *entity.HealthPassport) error {
	return r.db.WithContext(ctx).Delete(healthPassport).Error
}
