package midtrans

import (
	"context"
	"errors"
	"time"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	CreateTransaction(ctx context.Context, req dto.PaymentRequest) (string, error)
	WebHookTransaction(ctx context.Context, input *dto.DonationsCreate) error
}


type midtransService struct {
	cfg *configs.MidtransConfig
	snapClient snap.Client
	DonationsRepository repository.DonationsRepository
	
}


func InitMidtrans(cfg *configs.MidtransConfig) (MidtransService, error) {
	snapClient := snap.Client{}

	snapClient.New(cfg.ServerKey, midtrans.Sandbox)
	

	return &midtransService{
		cfg:        cfg,
		snapClient: snapClient,
	}, nil
}

func NewMidtransService(cfg *configs.MidtransConfig) *midtransService {
    snapClient := snap.Client{}
    snapClient.New(cfg.ServerKey, midtrans.Sandbox)
    
    return &midtransService{
        cfg: cfg,
        snapClient: snapClient,
    }
}

func (s *midtransService) CreateTransaction(ctx context.Context, req dto.PaymentRequest) (string, error) {
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Fullname,
			Email: req.Email,
		},
	}
	resp, err := s.snapClient.CreateTransaction(request)
	if err != nil {
		return "", err
	}
	return resp.RedirectURL, nil
}

func (s *midtransService) WebHookTransaction(ctx context.Context, input *dto.DonationsCreate) error {
	donation := new(entity.Donation)

	donation.UserId = input.UserID
	donation.Amount = input.Amount
	donation.Status = input.Transaction_status
	donation.CreatedAt = time.Now()
	donation.UpdatedAt = time.Now()
	
	// Parse transaction time
	transactionTime, err := time.Parse("2006-01-02 15:04:05", input.Transaction_time)
	if err != nil {
		return errors.New("Invalid transaction time format")
	}
	donation.TransactionTime = transactionTime

	// Validate transaction status
	if input.Transaction_status != "settlement" && input.Transaction_status != "capture" {
		return errors.New("Invalid transaction status")
	} 
	donation.Status = "success"
	
	// Save donation to database
	if err := s.DonationsRepository.Create(ctx, donation); err != nil {
		return errors.New("Failed to process donation")
	}

	return nil
}