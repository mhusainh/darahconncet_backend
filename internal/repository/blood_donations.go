package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type BloodDonationRepository interface {
	Create(ctx context.Context, bloodDonation *entity.BloodDonation) error
	GetById(ctx context.Context, id int64) (*entity.BloodDonation, error)
	GetAll(ctx context.Context, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error)
	GetByUserId(ctx context.Context, userId int64, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error)
	Update(ctx context.Context, bloodDonation *entity.BloodDonation) error
	Delete(ctx context.Context, bloodDonation *entity.BloodDonation) error
	GetByUser(ctx context.Context, userId int64) ([]entity.BloodDonation, error)
	CountSuccessDonation(ctx context.Context) (int64, error)
}

type bloodDonationRepository struct {
	db *gorm.DB
}

func NewBloodDonationRepository(db *gorm.DB) BloodDonationRepository {
	return &bloodDonationRepository{db}
}

func (r *bloodDonationRepository) applyFilters(query *gorm.DB, req dto.GetAllBloodDonationRequest) (*gorm.DB, dto.GetAllBloodDonationRequest) {
	if req.Status != "" {
		query = query.Where("LOWER(status) = ?", req.Status)
	}

	if req.BloodType != "" {
		query = query.Where("LOWER(blood_type) = ?", req.BloodType)
	}

	// Filter berdasarkan tanggal event
	if req.StartDate != "" && req.EndDate != "" {
		query = query.Where("donation_date BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	// Filter berdasarkan Search (pada nama user, catatan registrasi, atau nama event)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN users ON users.id = blood_donations.user_id").
			Joins("LEFT JOIN hospitals ON hospitals.id = blood_donations.hospital_id").
			Joins("LEFT JOIN registrations ON registrations.id = blood_donations.registration_id").
			Where("LOWER(users.name) LIKE ? OR LOWER(registrations.note) LIKE ? OR LOWER(hospitals.name) LIKE ? OR LOWER(hospitals.city) LIKE ? OR LOWER(hospitals.province) LIKE ?",
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

func (r *bloodDonationRepository) Create(ctx context.Context, bloodDonation *entity.BloodDonation) error {
	return r.db.WithContext(ctx).Create(bloodDonation).Error
}

func (r *bloodDonationRepository) GetById(ctx context.Context, id int64) (*entity.BloodDonation, error) {
	result := new(entity.BloodDonation)
	if err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Hospital").Preload("Registration").First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *bloodDonationRepository) GetAll(ctx context.Context, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error) {
	var bloodDonation []entity.BloodDonation
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodDonation{}).Preload("Hospital").Preload("Registration")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodDonation).Error; err != nil {
		return nil, 0, err
	}

	return bloodDonation, total, nil
}

func (r *bloodDonationRepository) GetByUserId(ctx context.Context, UserId int64, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error) {
	var bloodDonation []entity.BloodDonation
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.BloodDonation{}).Where("user_id = ?", UserId).Preload("Hospital").Preload("Registration")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&bloodDonation).Error; err != nil {
		return nil, 0, err
	}

	return bloodDonation, total, nil
}

func (r *bloodDonationRepository) Update(ctx context.Context, bloodDonation *entity.BloodDonation) error {
	return r.db.WithContext(ctx).Model(bloodDonation).Updates(bloodDonation).Error
}

func (r *bloodDonationRepository) Delete(ctx context.Context, bloodDonation *entity.BloodDonation) error {
	return r.db.WithContext(ctx).Model(&entity.BloodDonation{}).Delete(bloodDonation).Error
}

func (r *bloodDonationRepository) GetByUser(ctx context.Context, userId int64) ([]entity.BloodDonation, error) {
	result := make([]entity.BloodDonation, 0)
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *bloodDonationRepository) CountSuccessDonation(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.BloodDonation{}).Where("status = ?", "completed").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}