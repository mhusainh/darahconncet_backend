package service

import (
	"context"
	"errors"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
)

type BloodDonationService interface {
	Create(ctx context.Context, req dto.BloodDonationCreateRequest) error
	GetAll(ctx context.Context, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error)
	GetByUserId(ctx context.Context, userId int64, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error)
	GetById(ctx context.Context, id int64) (*entity.BloodDonation, error)
	Update(ctx context.Context, req dto.BloodDonationUpdateRequest, bloodDonation *entity.BloodDonation) (*entity.BloodDonation,error)
	Delete(ctx context.Context, id int64) error
}

type bloodDonationService struct {
	bloodDonationRepository repository.BloodDonationRepository
	cloudinaryService       cloudinary.Service
}

func NewBloodDonationService(
	bloodDonationRepository repository.BloodDonationRepository,
	cloudinaryService cloudinary.Service,
) BloodDonationService {
	return &bloodDonationService{
		bloodDonationRepository,
		cloudinaryService,
	}
}

func (s *bloodDonationService) Create(ctx context.Context, req dto.BloodDonationCreateRequest) error {
	bloodDonation := new(entity.BloodDonation)
	bloodDonation.UserId = req.UserId
	bloodDonation.HospitalId = req.HospitalId
	bloodDonation.RegistrationId = req.RegistrationId
	bloodDonation.DonationDate = req.DonationDate
	bloodDonation.BloodType = req.BloodType
	bloodDonation.Status = req.Status

	UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodDonations")
	if err != nil {
		return errors.New("Gagal mengupload gambar")
	}

	// Simpan URL dan Public ID ke request
	bloodDonation.UrlFile = UrlFile
	bloodDonation.PublicId = publicId

	if err := s.bloodDonationRepository.Create(ctx, bloodDonation); err != nil {
		if bloodDonation.PublicId != "" {
			if err := s.cloudinaryService.DeleteFile(bloodDonation.PublicId); err != nil {
				return errors.New("Gagal menghapus gambar")
			}
		}
		return errors.New("Gagal membuat donasi darah")
	}

	return nil
}

func (s *bloodDonationService) GetAll(ctx context.Context, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error) {
	bloodDonations, total, err := s.bloodDonationRepository.GetAll(ctx, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mengambil donasi darah")
	}

	return bloodDonations, total, nil
}

func (s *bloodDonationService) GetByUserId(ctx context.Context, userId int64, req dto.GetAllBloodDonationRequest) ([]entity.BloodDonation, int64, error) {
	bloodDonations, total, err := s.bloodDonationRepository.GetByUserId(ctx, userId, req)
	if err != nil {
		return nil, 0, errors.New("Gagal mengambil donasi darah")
	}

	return bloodDonations, total, nil
}

func (s *bloodDonationService) GetById(ctx context.Context, id int64) (*entity.BloodDonation, error) {
	bloodDonation, err := s.bloodDonationRepository.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Gagal mengambil donasi darah")
	}

	return bloodDonation, nil
}

func (s *bloodDonationService) Update(ctx context.Context, req dto.BloodDonationUpdateRequest, bloodDonation *entity.BloodDonation) (*entity.BloodDonation, error) {
    if !req.DonationDate.IsZero() {
        bloodDonation.DonationDate = req.DonationDate
    }
    if req.BloodType != "" {
        bloodDonation.BloodType = req.BloodType
    }
    if req.Status != "" {
        bloodDonation.Status = req.Status
    }
    
    var oldPublicId string
    var newPublicId string
    
    if req.Image != nil {
        // Simpan publicId lama sebelum mengubahnya
        oldPublicId = bloodDonation.PublicId
        
        UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "BloodDonations")
        if err != nil {
            return nil, errors.New("Gagal mengupload gambar")

        }
        newPublicId = publicId
        bloodDonation.UrlFile = UrlFile
        bloodDonation.PublicId = publicId
    }

    if err := s.bloodDonationRepository.Update(ctx, bloodDonation); err != nil {
        // Jika database update gagal dan ada gambar baru yang diunggah, hapus gambar baru
        if req.Image != nil {
            if err := s.cloudinaryService.DeleteFile(newPublicId); err != nil {
                return nil, errors.New("Gagal menghapus gambar baru")
            }
        }
        return nil, errors.New("Gagal mengupdate donasi darah")
    }
    
    // Jika berhasil dan ada gambar lama, hapus gambar lama
    if req.Image != nil && oldPublicId != "" {
        if err := s.cloudinaryService.DeleteFile(oldPublicId); err != nil {
            return nil, errors.New("Gagal menghapus gambar lama")
        }
    }
    
    return bloodDonation, nil
}

func (s *bloodDonationService) Delete(ctx context.Context, id int64) error {
	bloodDonation, err := s.bloodDonationRepository.GetById(ctx, id)
	if err != nil {
		return errors.New("Gagal mengambil donasi darah")
	}

	if err := s.bloodDonationRepository.Delete(ctx, bloodDonation); err != nil {
		return errors.New("Gagal menghapus donasi darah")
	}

	return nil
}
