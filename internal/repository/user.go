package repository

import (
	"context"
	"strings"


	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]entity.User, int64, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
	GetByResetPasswordToken(ctx context.Context, token string) (*entity.User, error)
	GetByVerifyEmailToken(ctx context.Context, token string) (*entity.User, error)
	CountUser(ctx context.Context) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// applyFilters menerapkan filter, sorting, dan pagination ke query GORM
func (r *userRepository) applyFilters(query *gorm.DB, req dto.GetAllUserRequest) (*gorm.DB, dto.GetAllUserRequest) {
	// Filter berdasarkan BloodType
	if req.BloodType != "" {
		query = query.Where("LOWER(blood_type) = ?", req.BloodType)
	}
	if req.Email != "" {
		query = query.Where("LOWER(email) = ?", req.Email)
	}
	// Filter berdasarkan Search (pada judul atau pesan)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Where("LOWER(name) LIKE ?", "%"+search+"%").
			Or("LOWER(gender) LIKE ?", "%"+search+"%").
			Or("LOWER(email) LIKE ?", "%"+search+"%").
			Or("LOWER(phone) LIKE ?", "%"+search+"%").
			Or("LOWER(birth_date) LIKE ?", "%"+search+"%").
			Or("LOWER(address) LIKE ?", "%"+search+"%")
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

func (r *userRepository) GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.User{})
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) GetById(ctx context.Context, id int64) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Model(&user).Updates(&user).Error
}

func (r *userRepository) Delete(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Delete(&user).Error
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByResetPasswordToken(ctx context.Context, token string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("reset_password_token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByVerifyEmailToken(ctx context.Context, token string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("verify_email_token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) CountUser(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("role = ?", "User").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
