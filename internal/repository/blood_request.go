package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type BloodRequestRepository interface {
	Create(ctx context.Context, bloodRequest *entity.BloodRequest) error
	GetById(ctx context.Context, id int64) (*entity.BloodRequest, error)
	GetByUserId(ctx context.Context, userId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetByHospitalId(ctx context.Context, hospitalId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllAdminBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllCampaign(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	Update(ctx context.Context, bloodRequest *entity.BloodRequest) error
	Delete(ctx context.Context, bloodRequest *entity.BloodRequest) error
	CountBloodRequest(ctx context.Context, status string, eventType string) (int64, error)
	CountCampaignActive(ctx context.Context, status string, eventType string) (int64, error)
	CountTotal(ctx context.Context, eventType string) (int64, error)
	CountAllTotal(ctx context.Context) (int64, error)
	CountByMonth(ctx context.Context, month string, year string) (int64, error)
}

type bloodRequestRepository struct {
	db *gorm.DB
}

func NewBloodRequestRepository(db *gorm.DB) BloodRequestRepository {
	return &bloodRequestRepository{db}
}

func (r *bloodRequestRepository) applyFilters(query *gorm.DB, req dto.GetAllBloodRequestRequest) (*gorm.DB, dto.GetAllBloodRequestRequest) {
	if req.UrgencyLevel != "" {
		query = query.Where("LOWER(urgency_level) = ?", req.UrgencyLevel)
	}

	if req.EventType != "" {
		query = query.Where("LOWER(event_type) = ?", req.EventType)
	}

	if req.BloodType != "" {
		query = query.Where("LOWER(blood_type) = ?", req.BloodType)
	}

	if req.MaxQuantity > 0 && req.MinQuantity > 0 {
		query = query.Where("quantity BETWEEN ? AND ?", req.MinQuantity, req.MaxQuantity)
	}

	// Filter berdasarkan tanggal event
	if req.StartDate != "" && req.EndDate != "" {
		query = query.Where("expiry_date BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	// Filter berdasarkan Search (pada nama user, catatan registrasi, atau nama event)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN users ON users.id = blood_requests.user_id").
			Joins("LEFT JOIN hospitals ON hospitals.id = blood_requests.hospital_id").
			Where("LOWER(blood_requests.patient_name) LIKE ? OR LOWER(users.name) LIKE ? OR LOWER(blood_requests.patient_name) LIKE ? OR LOWER(hospitals.name) LIKE ? OR LOWER(hospitals.city) LIKE ? OR LOWER(hospitals.province) LIKE ?",
				"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
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

func (r *bloodRequestRepository) Create(ctx context.Context, bloodRequest *entity.BloodRequest) error {
	return r.db.WithContext(ctx).Create(bloodRequest).Error
}

func (r *bloodRequestRepository) GetById(ctx context.Context, id int64) (*entity.BloodRequest, error) {
	result := new(entity.BloodRequest)
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("id = ?", id).Preload("User").Preload("Hospital").First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *bloodRequestRepository) GetAllBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	var bloodRequest []entity.BloodRequest
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("event_type = ? AND status = ?", "blood_request", "verified").Preload("User").Preload("Hospital")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodRequest).Error; err != nil {
		return nil, 0, err
	}

	return bloodRequest, total, nil
}

func (r *bloodRequestRepository) GetAllAdminBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	var bloodRequest []entity.BloodRequest
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Preload("User").Preload("Hospital")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodRequest).Error; err != nil {
		return nil, 0, err
	}

	return bloodRequest, total, nil
}

func (r *bloodRequestRepository) GetAllCampaign(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	var bloodRequest []entity.BloodRequest
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("event_type = ?", "campaign").Preload("User").Preload("Hospital")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodRequest).Error; err != nil {
		return nil, 0, err
	}

	return bloodRequest, total, nil
}

func (r *bloodRequestRepository) GetByUserId(ctx context.Context, userId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	var bloodRequest []entity.BloodRequest
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Preload("User").Preload("Hospital").Where("user_id = ?", userId)
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodRequest).Error; err != nil {
		return nil, 0, err
	}

	return bloodRequest, total, nil
}

func (r *bloodRequestRepository) GetByHospitalId(ctx context.Context, hospitalId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	var bloodRequest []entity.BloodRequest
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Preload("User").Preload("Hospital").Where("hospital_id = ?", hospitalId)
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodRequest).Error; err != nil {
		return nil, 0, err
	}

	return bloodRequest, total, nil
}


func (r *bloodRequestRepository) Update(ctx context.Context, bloodRequest *entity.BloodRequest) error {
	return r.db.WithContext(ctx).Model(bloodRequest).Updates(bloodRequest).Error
}

func (r *bloodRequestRepository) Delete(ctx context.Context, bloodRequest *entity.BloodRequest) error {
	return r.db.WithContext(ctx).Delete(bloodRequest).Error
}

func (r *bloodRequestRepository) CountBloodRequest(ctx context.Context, status string, eventType string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("status = ? AND event_type = ?", status, eventType).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *bloodRequestRepository) CountTotal(ctx context.Context, eventType string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("event_type = ?", eventType).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *bloodRequestRepository) CountCampaignActive(ctx context.Context, status string, eventType string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("status != ? AND event_type = ?", status, eventType).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *bloodRequestRepository) CountAllTotal(ctx context.Context) (int64, error){
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *bloodRequestRepository) CountByMonth(ctx context.Context, month string, year string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodRequest{}).Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}