package repository

import (
	"context"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"

	"gorm.io/gorm"
)

type DonationsRepository interface {
	Create(ctx context.Context, donation *entity.Donation) error
	GetById(ctx context.Context, id int64) (*entity.Donation, error)
	GetAll(ctx context.Context) ([]entity.Donation, error)
}

type donationsRepository struct {
	db *gorm.DB
}

func NewDonationsRepository(db *gorm.DB) DonationsRepository {
	return &donationsRepository{db}
}

func (r *donationsRepository) Create(ctx context.Context, donation *entity.Donation) error {
	return r.db.WithContext(ctx).Create(donation).Error
}

func (r *donationsRepository) GetById(ctx context.Context, id int64) (*entity.Donation, error) {
	result := new(entity.Donation)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *donationsRepository) GetAll(ctx context.Context) ([]entity.Donation, error) {
	result := make([]entity.Donation, 0)
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
