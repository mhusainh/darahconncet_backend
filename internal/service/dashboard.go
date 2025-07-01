package service

import (
	"context"
	"errors"
	"time"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
)

type DashboardService interface {
	DashboardUser(ctx context.Context, userId int64) (map[string]interface{}, error)
	DashboardAdmin(ctx context.Context) (map[string]interface{}, error)
	LandingPage(ctx context.Context) (map[string]interface{}, error)
}

type dashboardService struct {
	bloodDonationRepository repository.BloodDonationRepository
	bloodRequestRepository  repository.BloodRequestRepository
	userRepository          repository.UserRepository
}

func NewDashboardService(
	bloodDonationRepository repository.BloodDonationRepository,
	bloodRequestRepository repository.BloodRequestRepository,
	userRepository repository.UserRepository,
) DashboardService {
	return &dashboardService{
		bloodDonationRepository,
		bloodRequestRepository,
		userRepository,
	}
}

func (s *dashboardService) DashboardUser(ctx context.Context, userId int64) (map[string]interface{}, error) {
	bloodDonations, err := s.bloodDonationRepository.GetByUser(ctx, userId)
	if err != nil {
		return nil, errors.New("Gagal mendapatkan data dashboard")
	}
	var lastDonation interface{}
	if len(bloodDonations) > 0 {
		lastDonation = bloodDonations[0].CreatedAt
	}
	data := map[string]interface{}{
		"total_donor":      int(len(bloodDonations)),
		"last_donation":    lastDonation,
		"total_sertifikat": countCompletedDonations(bloodDonations),
	}
	return data, nil
}

func (s *dashboardService) DashboardAdmin(ctx context.Context) (map[string]interface{}, error) {
	totalDonation, err := s.bloodRequestRepository.CountTotal(ctx, "blood_request")
	if err != nil {
		return nil, errors.New("Gagal menghitung total Blood Request")
	}
	totalCampaign, err := s.bloodRequestRepository.CountTotal(ctx, "campaign")
	if err != nil {
		return nil, errors.New("Gagal menghitung total Campaign")
	}
	DonorTerverifikasi, err := s.bloodRequestRepository.CountBloodRequest(ctx, "verified", "blood_request")
	if err != nil {
		return nil, errors.New("Gagal menghitung total Donor Terverifikasi")
	}
	RequestPending, err := s.bloodRequestRepository.CountBloodRequest(ctx, "pending", "blood_request")
	if err != nil {
		return nil, errors.New("Gagal menghitung total Request Pending")
	}
	CampaignActive, err := s.bloodRequestRepository.CountCampaignActive(ctx, "pending", "blood_request")
	if err != nil {
		return nil, errors.New("Gagal menghitung total Request Pending")
	}

	data := map[string]interface{}{
		"total_donor":         totalDonation,
		"total_campaign":      totalCampaign,
		"donor_terverifikasi": DonorTerverifikasi,
		"request_pending":     RequestPending,
		"campaign_active":     CampaignActive,
	}
	return data, nil
}

func (s *dashboardService) LandingPage(ctx context.Context) (map[string]interface{}, error) {
	totalDonation, err := s.bloodRequestRepository.CountAllTotal(ctx)
	if err != nil {
		return nil, errors.New("Gagal menghitung total Blood Request")
	}
	donationSuccess, err := s.bloodDonationRepository.CountSuccessDonation(ctx)
	if err != nil {
		return nil, errors.New("Gagal menghitung total Donasi Sukses")
	}
	totalUser, err := s.userRepository.CountUser(ctx)
	if err != nil {
		return nil, errors.New("Gagal menghitung total User")
	}
	currentMonth := time.Now().Format("01")
	currentYear := time.Now().Format("2006")
	bloodRequest, err := s.bloodRequestRepository.CountByMonth(ctx, currentMonth, currentYear)
	if err != nil {
		return nil, errors.New("Gagal menghitung total Blood Request")
	}

	data := map[string]interface{}{
		"donasi_terdaftar": totalDonation,
		"donasi_sukses":    donationSuccess,
		"total_user":       totalUser,
		"event_bulan_ini":  bloodRequest,
	}
	return data, nil
}

func countCompletedDonations(donations []entity.BloodDonation) int {
	count := 0
	for _, donation := range donations {
		if donation.Status == "completed" {
			count++
		}
	}
	return count
}
