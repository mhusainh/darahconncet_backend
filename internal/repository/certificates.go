package repository

import (
	"context"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type CertificateRepository interface {
	Create(ctx context.Context, certificate *entity.Certificate) error
	GetById(ctx context.Context, id int64) (*entity.Certificate, error)
	GetAll(ctx context.Context, req dto.GetAllCertificateRequest) ([]entity.Certificate, int64, error)
	GetByUser(ctx context.Context, userId int64, req dto.GetAllCertificateRequest) ([]entity.Certificate, int64, error)
	Update(ctx context.Context, certificate *entity.Certificate) error
	Delete(ctx context.Context, certificate *entity.Certificate) error
}

type certificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &certificateRepository{db}
}

func (r *certificateRepository) applyFilters(query *gorm.DB, req dto.GetAllCertificateRequest) (*gorm.DB, dto.GetAllCertificateRequest) {
	// Filter berdasarkan Search (pada nama user, catatan registrasi, atau nama event)
	if req.DonationId != "" {
		query = query.Where("donation_id = ?", req.DonationId)
	}

	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN users ON users.id = certificates.user_id").
			Where("LOWER(users.name) LIKE ? OR LOWER(certificates.certificate_number) LIKE ? OR LOWER(certificates.digital_signature) LIKE ?",
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

func (r *certificateRepository) Create(ctx context.Context, certificate *entity.Certificate) error {
	return r.db.WithContext(ctx).Create(certificate).Error
}

func (r *certificateRepository) GetById(ctx context.Context, id int64) (*entity.Certificate, error) {
	result := new(entity.Certificate)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *certificateRepository) GetAll(ctx context.Context, req dto.GetAllCertificateRequest) ([]entity.Certificate, int64, error) {
	var certificates []entity.Certificate
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Certificate{}).Preload("User")
	if req.UserId != "" {
		dataQuery = dataQuery.Where("user_id = ?", req.UserId)
	}
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&certificates).Error; err != nil {
		return nil, 0, err
	}

	return certificates, total, nil
}

func (r *certificateRepository) GetByUser(ctx context.Context, userId int64, req dto.GetAllCertificateRequest) ([]entity.Certificate, int64, error) {

	var certificates []entity.Certificate
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Certificate{}).Preload("User").Where("user_id = ?", userId)
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&certificates).Error; err != nil {
		return nil, 0, err
	}

	return certificates, total, nil
}

func (r *certificateRepository) GetByDonationId(ctx context.Context, donationId int64, req dto.GetAllCertificateRequest) ([]entity.Certificate, int64, error) {
	var certificates []entity.Certificate
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Certificate{}).Where("donation_id = ?", donationId)
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&certificates).Error; err != nil {
		return nil, 0, err
	}

	return certificates, total, nil
}

func (r *certificateRepository) Update(ctx context.Context, certificate *entity.Certificate) error {
	return r.db.WithContext(ctx).Model(certificate).Updates(certificate).Error
}

func (r *certificateRepository) Delete(ctx context.Context, certificate *entity.Certificate) error {
	return r.db.WithContext(ctx).Delete(certificate).Error
}
