package service

import (
	"context"
	"errors"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
)

type NotificationService interface {
	Create(ctx context.Context, req dto.NotificationCreateRequest) error
	GetById(ctx context.Context, id int64) (*entity.Notification, error)
	GetAll(ctx context.Context, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error)
	Update(ctx context.Context, req dto.NotificationUpdateRequest, notification *entity.Notification) error
	Delete(ctx context.Context, id int64) error
	GetByUserId(ctx context.Context, userId int64, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error)
	GetUnreadCountByUserId(ctx context.Context, userId int64) (int64, error)
}

type notificationService struct {
	notificationRepository repository.NotificationRepository
	userRepository repository.UserRepository
}

func NewNotificationService(notificationRepository repository.NotificationRepository, userRepository repository.UserRepository,
) NotificationService {
	return &notificationService{notificationRepository, userRepository}
}

func (s *notificationService) GetAll(ctx context.Context, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error) {
	notifications, total, err := s.notificationRepository.GetAll(ctx, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan daftar notifikasi")
	}
	return notifications, total, nil
}

func (s *notificationService) GetById(ctx context.Context, id int64) (*entity.Notification, error) {
	notification, err := s.notificationRepository.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Notifikasi tidak ditemukan")
	}
	return notification, nil
}

func (s *notificationService) GetByUserId(ctx context.Context, userId int64, req dto.GetAllNotificationRequest) ([]entity.Notification, int64, error) {
	notifications, total, err := s.notificationRepository.GetByUserId(ctx, userId, req)
	if err != nil {
		return nil, 0, errors.New("Notifikasi tidak ditemukan untuk pengguna ini")
	}
	return notifications, total, nil
}

func (s *notificationService) GetUnreadCountByUserId(ctx context.Context, userId int64) (int64, error) {
	count, err := s.notificationRepository.GetUnreadCountByUserId(ctx, userId)
	if err != nil {
		return 0, errors.New("Gagal mendapatkan jumlah notifikasi belum dibaca")
	}
	return count, nil
}

func (s *notificationService) Create(ctx context.Context, req dto.NotificationCreateRequest) error {
	notification := new(entity.Notification)
	notification.UserId = req.UserId
	notification.Title = req.Title
	notification.Message = req.Message
	notification.NotificationType = req.NotificationType
	notification.IsRead = false // Default to unread
	
	if _,err := s.userRepository.GetById(ctx,notification.UserId); err != nil{
		return errors.New("user tidak ditemukan")
	}
	if err := s.notificationRepository.Create(ctx, notification); err != nil {
		return errors.New("Notifikasi gagal dibuat")
	}
	return nil
}

func (s *notificationService) Update(ctx context.Context, req dto.NotificationUpdateRequest, notification *entity.Notification) error {
	if req.Title != "" {
		notification.Title = req.Title
	}
	if req.Message != "" {
		notification.Message = req.Message
	}
	if req.NotificationType != "" {
		notification.NotificationType = req.NotificationType
	}
	if req.IsRead != false {
		notification.IsRead = req.IsRead
	}

	if err := s.notificationRepository.Update(ctx, notification); err != nil {
		return errors.New("Notifikasi gagal diperbarui")
	}
	return nil
}

func (s *notificationService) Delete(ctx context.Context, id int64) error {
	notification, err := s.notificationRepository.GetById(ctx, id)
	if err != nil {
		return errors.New("Notifikasi tidak ditemukan")
	}

	if err := s.notificationRepository.Delete(ctx, notification); err != nil {
		return errors.New("Notifikasi gagal dihapus")
	}
	return nil
}
