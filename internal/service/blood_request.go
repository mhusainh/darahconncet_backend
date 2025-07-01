package service

import (
	"context"
	"errors"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
)

type BloodRequestService interface {
	RegistrationDonate(ctx context.Context, registrationStatus string, bloodRequest *entity.BloodRequest) error
	CreateBloodRequest(ctx context.Context, req dto.BloodRequestCreateRequest) error
	CreateCampaign(ctx context.Context, req dto.CampaignCreateRequest) error
	GetAllBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllBloodRequestByUser(ctx context.Context, userId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllAdminBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetAllCampaign(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error)
	GetById(ctx context.Context, id int64) (*entity.BloodRequest, error)
	UpdateCampaign(ctx context.Context, req dto.CampaignUpdateRequest, bloodRequest *entity.BloodRequest) error
	UpdateBloodRequest(ctx context.Context, req dto.BloodRequestUpdateRequest, bloodRequest *entity.BloodRequest) error
	Delete(ctx context.Context, id int64) error
}

type bloodRequestService struct {
	bloodRequestRepository repository.BloodRequestRepository
	cloudinaryService     cloudinary.Service
}

func NewBloodRequestService(bloodRequestRepository repository.BloodRequestRepository, cloudinaryService cloudinary.Service) BloodRequestService {
	return &bloodRequestService{
		bloodRequestRepository,
		cloudinaryService,
	}
}

func (s *bloodRequestService) CreateBloodRequest(ctx context.Context, req dto.BloodRequestCreateRequest) error {
	bloodRequest := new(entity.BloodRequest)
	bloodRequest.UserId = req.UserId
	bloodRequest.PatientName = req.PatientName
	bloodRequest.HospitalId = req.HospitalId
	bloodRequest.EventName = req.EventName
	bloodRequest.BloodType = req.BloodType
	bloodRequest.Quantity = req.Quantity
	bloodRequest.UrgencyLevel = req.UrgencyLevel
	bloodRequest.Diagnosis = req.Diagnosis
	bloodRequest.Status = "pending"
	bloodRequest.EventDate = req.EventDate
	bloodRequest.EventType = "blood_request"

	// Upload gambar jika ada
	if req.Image != nil {
		UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodRequests")
		if err != nil {
			return errors.New("Gagal mengupload gambar")
		}
		bloodRequest.UrlFile = UrlFile
		bloodRequest.PublicId = publicId
	}

	if err := s.bloodRequestRepository.Create(ctx, bloodRequest); err != nil {
		// Jika gagal menyimpan ke database, hapus gambar yang sudah diupload
		if bloodRequest.PublicId != "" {
			_ = s.cloudinaryService.DeleteFile(bloodRequest.PublicId)
		}
		return errors.New("Gagal membuat permintaan darah")
	}

	return nil
}

func (s *bloodRequestService) CreateCampaign(ctx context.Context, req dto.CampaignCreateRequest) error {
	bloodRequest := new(entity.BloodRequest)
	bloodRequest.UserId = req.UserId
	bloodRequest.HospitalId = req.HospitalId
	bloodRequest.EventName = req.EventName
	bloodRequest.StartTime = req.StartTime
	bloodRequest.EndTime = req.EndTime
	bloodRequest.SlotsAvailable = req.SlotsAvailable
	bloodRequest.SlotsBooked = req.SlotsBooked
	bloodRequest.EventDate = req.EventDate
	bloodRequest.EventType = "campaign"

	if req.Image != nil {
		UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodRequests")
		if err != nil {
			return errors.New("Gagal mengupload gambar")
		}
		bloodRequest.UrlFile = UrlFile
		bloodRequest.PublicId = publicId
	}

	if err := s.bloodRequestRepository.Create(ctx, bloodRequest); err != nil {
		// Jika gagal menyimpan ke database, hapus gambar yang sudah diupload
		if bloodRequest.PublicId != "" {
			_ = s.cloudinaryService.DeleteFile(bloodRequest.PublicId)
		}
		return errors.New("Gagal membuat permintaan darah")
	}

	return nil
}

func (s *bloodRequestService) GetAllBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	bloodRequests, total, err := s.bloodRequestRepository.GetAllBloodRequest(ctx, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan permintaan darah")
	}

	return bloodRequests, total, nil
}

func (s *bloodRequestService) GetAllBloodRequestByUser(ctx context.Context, userId int64, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	bloodRequests, total, err := s.bloodRequestRepository.GetByUserId(ctx, userId, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan permintaan darah")
	}

	return bloodRequests, total, nil
}

func (s *bloodRequestService) GetAllAdminBloodRequest(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	bloodRequests, total, err := s.bloodRequestRepository.GetAllAdminBloodRequest(ctx, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan permintaan darah")
	}

	return bloodRequests, total, nil
}

func (s *bloodRequestService) GetAllCampaign(ctx context.Context, req dto.GetAllBloodRequestRequest) ([]entity.BloodRequest, int64, error) {
	bloodRequests, total, err := s.bloodRequestRepository.GetAllCampaign(ctx, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan permintaan darah")
	}

	return bloodRequests, total, nil
}

func (s *bloodRequestService) GetById(ctx context.Context, id int64) (*entity.BloodRequest, error) {
	bloodRequest, err := s.bloodRequestRepository.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Permintaan darah tidak ditemukan")
	}

	return bloodRequest, nil
}

func (s *bloodRequestService) UpdateBloodRequest(ctx context.Context, req dto.BloodRequestUpdateRequest, bloodRequest *entity.BloodRequest) error {
	if req.EventName != "" {
		bloodRequest.EventName = req.EventName
	}
	if req.BloodType != "" {
		bloodRequest.BloodType = req.BloodType
	}
	if req.PatientName != "" {
		bloodRequest.PatientName = req.PatientName
	}
	if req.Quantity != 0 {
		bloodRequest.Quantity = req.Quantity
	}
	if req.UrgencyLevel != "" {
		bloodRequest.UrgencyLevel = req.UrgencyLevel
	}
	if req.Diagnosis != "" {
		bloodRequest.Diagnosis = req.Diagnosis
	}
	if req.Status != "" {
		bloodRequest.Status = req.Status
	}
	if !req.EventDate.IsZero() {
		bloodRequest.EventDate = req.EventDate
	}

	// Simpan publicId lama sebelum mengubahnya
	oldPublicId := bloodRequest.PublicId
	var newPublicId string

	// Upload gambar baru jika ada
	if req.Image != nil {
		UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodRequests")
		if err != nil {
			return errors.New("Gagal mengupload gambar")
		}
		newPublicId = publicId
		bloodRequest.UrlFile = UrlFile
		bloodRequest.PublicId = publicId
	}

	if err := s.bloodRequestRepository.Update(ctx, bloodRequest); err != nil {
		// Jika gagal update database dan ada gambar baru yang diupload, hapus gambar baru
		if newPublicId != "" {
			_ = s.cloudinaryService.DeleteFile(newPublicId)
		}
		return errors.New("Gagal mengupdate permintaan darah")
	}

	// Jika berhasil update database dan ada gambar baru, hapus gambar lama
	if oldPublicId != "" && newPublicId != "" {
		_ = s.cloudinaryService.DeleteFile(oldPublicId)
	}

	return nil
}

func (s *bloodRequestService) UpdateCampaign(ctx context.Context, req dto.CampaignUpdateRequest, bloodRequest *entity.BloodRequest) error {
	if req.EventName != "" {
		bloodRequest.EventName = req.EventName
	}
	if !req.StartTime.IsZero() {
		bloodRequest.StartTime = req.StartTime
	}
	if !req.EndTime.IsZero() {
		bloodRequest.EndTime = req.EndTime
	}
	if !req.EventDate.IsZero() {
		bloodRequest.EventDate = req.EventDate
	}
	bloodRequest.SlotsAvailable = req.SlotsAvailable
	bloodRequest.SlotsBooked = req.SlotsBooked

	// Simpan publicId lama sebelum mengubahnya
	oldPublicId := bloodRequest.PublicId
	var newPublicId string

	// Upload gambar baru jika ada
	if req.Image != nil {
		UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodRequests")
		if err != nil {
			return errors.New("Gagal mengupload gambar")
		}
		newPublicId = publicId
		bloodRequest.UrlFile = UrlFile
		bloodRequest.PublicId = publicId
	}

	if err := s.bloodRequestRepository.Update(ctx, bloodRequest); err != nil {
		// Jika gagal update database dan ada gambar baru yang diupload, hapus gambar baru
		if newPublicId != "" {
			_ = s.cloudinaryService.DeleteFile(newPublicId)
		}
		return errors.New("Gagal mengupdate permintaan darah")
	}

	// Jika berhasil update database dan ada gambar baru, hapus gambar lama
	if oldPublicId != "" && newPublicId != "" {
		_ = s.cloudinaryService.DeleteFile(oldPublicId)
	}

	return nil
}

func (s *bloodRequestService) Delete(ctx context.Context, id int64) error {
	bloodRequest, err := s.bloodRequestRepository.GetById(ctx, id)
	if err != nil {
		return errors.New("Permintaan darah tidak ditemukan")
	}

	// Simpan publicId untuk dihapus setelah data dihapus dari database
	publicId := bloodRequest.PublicId

	if err := s.bloodRequestRepository.Delete(ctx, bloodRequest); err != nil {
		return errors.New("Gagal menghapus permintaan darah")
	}

	// Hapus gambar dari cloudinary jika ada
	if publicId != "" {
		_ = s.cloudinaryService.DeleteFile(publicId)
	}

	return nil
}

func (s *bloodRequestService) RegistrationDonate(ctx context.Context, registrationStatus string, bloodRequest *entity.BloodRequest) error {
	switch registrationStatus {
	case "registered":
		bloodRequest.SlotsAvailable = bloodRequest.SlotsAvailable - 1
		bloodRequest.SlotsBooked = bloodRequest.SlotsBooked + 1

	case "cancelled":
		bloodRequest.SlotsAvailable = bloodRequest.SlotsAvailable + 1
		bloodRequest.SlotsBooked = bloodRequest.SlotsBooked - 1

	default:
		return errors.New("Status tidak valid")

	}

	if err := s.bloodRequestRepository.Update(ctx, bloodRequest); err != nil {
		return errors.New("Gagal mengupdate permintaan darah")
	}
	return nil
}
