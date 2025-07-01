package service

import (
	"context"
	"errors"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
)

type HospitalService interface {
	Create(ctx context.Context, req dto.HospitalCreateRequest) (*entity.Hospital, error)
	GetById(ctx context.Context, id int64) (*entity.Hospital, error)
	GetAll(ctx context.Context, req dto.GetAllHospitalRequest) ([]entity.Hospital, int64, error)
	Update(ctx context.Context, req dto.HospitalUpdateRequest, hospital *entity.Hospital) error
	Delete(ctx context.Context, id int64) error
}

type hospitalService struct {
	hospitalRepository repository.HospitalRepository
}

func NewHospitalService(hospitalRepository repository.HospitalRepository) HospitalService {
	return &hospitalService{hospitalRepository}
}

func (s *hospitalService) GetAll(ctx context.Context, req dto.GetAllHospitalRequest) ([]entity.Hospital, int64, error) {
	hospitals, total, err := s.hospitalRepository.GetAll(ctx, dto.GetAllHospitalRequest{})
	if err != nil {
		return nil, 0, errors.New("Gagal mendapatkan daftar rumah sakit")
	}
	return hospitals, total, nil
}

func (s *hospitalService) GetById(ctx context.Context, id int64) (*entity.Hospital, error) {
	hospital, err := s.hospitalRepository.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Rumah sakit tidak ditemukan")
	}
	return hospital, nil
}

func (s *hospitalService) Create(ctx context.Context, req dto.HospitalCreateRequest) (*entity.Hospital, error) {
	hospital := new(entity.Hospital)
	hospital.Name = req.Name
	hospital.Address = req.Address
	hospital.City = req.City
	hospital.Province = req.Province
	hospital.Latitude = req.Latitude
	hospital.Longitude = req.Longitude

	if err := s.hospitalRepository.Create(ctx, hospital); err != nil {
		return nil, errors.New("gagal membuat rumah sakit" + err.Error())
	}
	return hospital, nil
}

func (s *hospitalService) Update(ctx context.Context, req dto.HospitalUpdateRequest, hospital *entity.Hospital) error {
	if req.Name != "" {
		hospital.Name = req.Name
	}
	if req.Address != "" {
		hospital.Address = req.Address
	}
	if req.City != "" {
		hospital.City = req.City
	}
	if req.Province != "" {
		hospital.Province = req.Province
	}
	if req.Latitude != 0 {
		hospital.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		hospital.Longitude = req.Longitude
	}

	if err := s.hospitalRepository.Update(ctx, hospital); err != nil {
		return errors.New("Rumah sakit gagal diperbarui")
	}
	return nil
}

func (s *hospitalService) Delete(ctx context.Context, id int64) error {
	hospital, err := s.hospitalRepository.GetById(ctx, id)
	if err != nil {
		return errors.New("Rumah sakit tidak ditemukan")
	}

	if err := s.hospitalRepository.Delete(ctx, hospital); err != nil {
		return errors.New("Rumah sakit gagal dihapus")
	}
	return nil
}
