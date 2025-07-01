package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type DonorScheduleRepository interface {
	Create(ctx context.Context, donorSchedule *entity.DonorSchedule) error
	GetById(ctx context.Context, id int64) (*entity.DonorSchedule, error)
	GetAll(ctx context.Context, UserId int64, req dto.GetAllDonorScheduleRequest) ([]entity.DonorSchedule, int64, error)
	Update(ctx context.Context, donorSchedule *entity.DonorSchedule) error
	Delete(ctx context.Context, donorSchedule *entity.DonorSchedule) error
	GetByRequestId(ctx context.Context, requestId int64) (*entity.DonorSchedule, error)
}

type donorScheduleRepository struct {
	db *gorm.DB
}

func NewDonorScheduleRepository(db *gorm.DB) DonorScheduleRepository {
	return &donorScheduleRepository{db}
}

func (r *donorScheduleRepository) GetByRequestId(ctx context.Context, requestId int64) (*entity.DonorSchedule, error) {
	result := new(entity.DonorSchedule)
	if err := r.db.WithContext(ctx).Where("request_id = ?", requestId).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *donorScheduleRepository) applyFilters(query *gorm.DB, req dto.GetAllDonorScheduleRequest) (*gorm.DB, dto.GetAllDonorScheduleRequest) {
	// Filter berdasarkan EventDate
	if req.StartDate != "" && req.EndDate != "" {
		query = query.Where("event_date BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}


	if req.Status != "" {
		query = query.Where("LOWER(status) = ?", req.Status)
	}

	// Filter berdasarkan Search (pada judul atau pesan)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN hospitals ON hospitals.id = donor_schedules.hospital_id").
		Where("LOWER(donor_schedules.event_name) LIKE ? OR LOWER(hospitals.name) LIKE ? OR LOWER(donor_schedules.description) LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
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

func (r *donorScheduleRepository) Create(ctx context.Context, donorSchedule *entity.DonorSchedule) error {
	return r.db.WithContext(ctx).Create(donorSchedule).Error
}

func (r *donorScheduleRepository) GetById(ctx context.Context, id int64) (*entity.DonorSchedule, error) {
	result := new(entity.DonorSchedule)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *donorScheduleRepository) GetAll(ctx context.Context, UserId int64, req dto.GetAllDonorScheduleRequest) ([]entity.DonorSchedule, int64, error) {
	var donorSchedule []entity.DonorSchedule
	var total int64

	dataQuery := r.db.WithContext(ctx).Model(&entity.DonorSchedule{}).Where("user_id = ?", UserId).Preload("Hospital").Preload("BloodRequest.Hospital")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&donorSchedule).Error; err != nil {
		return nil, 0, err
	}

	return donorSchedule, total, nil
}

func (r *donorScheduleRepository) Update(ctx context.Context, donorSchedule *entity.DonorSchedule) error {
	return r.db.WithContext(ctx).Model(donorSchedule).Updates(donorSchedule).Error
}

func (r *donorScheduleRepository) Delete(ctx context.Context, donorSchedule *entity.DonorSchedule) error {
	return r.db.WithContext(ctx).Delete(donorSchedule).Error
}
