package repository

import (
	"context"
	"time"
	"strings"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *entity.Notification) error
	GetById(ctx context.Context, id int64) (*entity.Notification, error)
	GetAll(ctx context.Context, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error)
	Update(ctx context.Context, notification *entity.Notification) error
	Delete(ctx context.Context, notification *entity.Notification) error
	GetByUserId(ctx context.Context, userId int64, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error)
	GetUnreadCountByUserId(ctx context.Context, userId int64) (int64, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) Create(ctx context.Context, notification *entity.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

// applyFilters menerapkan filter, sorting, dan pagination ke query GORM
func (r *notificationRepository) applyFilters(query *gorm.DB, req dto.GetAllNotificationRequest) (*gorm.DB, dto.GetAllNotificationRequest) {
	// Filter berdasarkan IsRead
	if req.IsRead != nil {
		query = query.Where("LOWER(is_read) = ?", *req.IsRead)
	}

	// Filter berdasarkan Search (pada judul atau pesan)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Joins("LEFT JOIN users ON users.id = notifications.user_id").
			Where("LOWER(notifications.title) LIKE ? OR LOWER(notifications.notification_type) LIKE ? OR LOWER(users.name) LIKE ?",
				"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Filter berdasarkan tanggal
	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			query = query.Where("created_at >= ?", startDate)
		}
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			// Tambahkan 1 hari ke endDate untuk mencakup seluruh hari
			endDate = endDate.Add(24 * time.Hour)
			query = query.Where("created_at < ?", endDate)
		}
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

// GetAll mengambil semua notifikasi dengan filter dan pagination
func (r *notificationRepository) GetAll(ctx context.Context, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error) {
	var notifications []entity.Notification
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Notification{}).Preload("User")
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&notifications).Error; err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

func (r *notificationRepository) GetById(ctx context.Context, id int64) (*entity.Notification, error) {
	result := new(entity.Notification)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *notificationRepository) Update(ctx context.Context, notification *entity.Notification) error {
	return r.db.WithContext(ctx).Model(notification).Updates(notification).Error
}

func (r *notificationRepository) Delete(ctx context.Context, notification *entity.Notification) error {
	return r.db.WithContext(ctx).Delete(notification).Error
}

// GetByUserId mengambil notifikasi berdasarkan ID pengguna dengan filter dan pagination
func (r *notificationRepository) GetByUserId(ctx context.Context, userId int64, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error) {
	var notifications []entity.Notification
	var total int64

	// Hitung total item sebelum pagination
	dataQuery := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("user_id = ?", userId)
	dataQuery, req = r.applyFilters(dataQuery, req)
	if err := dataQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	dataQuery = dataQuery.Limit(int(req.Limit)).Offset(int(offset))

	if err := dataQuery.Find(&notifications).Error; err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

func (r *notificationRepository) GetUnreadCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("user_id = ? AND is_read = false", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
