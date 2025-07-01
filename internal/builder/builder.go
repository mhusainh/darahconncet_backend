package builder

import (
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/handler"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/router"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cache"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
	googleoauth "github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/googleOauth"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/mailer"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/midtrans"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/route"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *configs.Config, db *gorm.DB, rdb *redis.Client, cloudinaryService *cloudinary.Service, mailer *mailer.Mailer, blockchain service.BlockchainService) []route.Route {
	cacheable := cache.NewCacheable(rdb)
	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)

	//repository
	userRepository := repository.NewUserRepository(db)
	bloodRequestRepository := repository.NewBloodRequestRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	bloodDonationRepository := repository.NewBloodDonationRepository(db)
	certificateRepository := repository.NewCertificateRepository(db)
	donorRegistrationRepository := repository.NewDonorRegistrationRepository(db)
	donationsRepository := repository.NewDonationsRepository(db)
	//end

	//service
	userService := service.NewUserService(userRepository, tokenUseCase, cacheable, cfg, mailer, cloudinaryService)
	bloodRequestService := service.NewBloodRequestService(bloodRequestRepository, *cloudinaryService)
	notificationService := service.NewNotificationService(notificationRepository,userRepository)
	bloodDonationService := service.NewBloodDonationService(bloodDonationRepository, *cloudinaryService)
	certificateService := service.NewCertificateService(certificateRepository)
	donorRegistrationService := service.NewDonorRegistrationService(donorRegistrationRepository)
	// Buat instance midtransService
	midtransService := midtrans.NewMidtransService(&cfg.MidtransConfig)
	// Set donationsRepository
	midtransService.DonationsRepository = donationsRepository
	dashboardService := service.NewDashboardService(bloodDonationRepository, bloodRequestRepository, userRepository)
	googleAuthService := googleoauth.NewGoogleOAuthService(tokenUseCase, userService, &cfg.GoogleOauth)
	//end

	//handler
	userHandler := handler.NewUserHandler(userService, cloudinaryService, googleAuthService)
	bloodRequestHandler := handler.NewBloodRequestHandler(bloodRequestService, notificationService)
	bloodDonationHandler := handler.NewBloodDonationHandler(bloodDonationService, notificationService, certificateService, donorRegistrationService, userService, blockchain)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	certificateHandler := handler.NewCertificateHandler(certificateService)
	donationHandler := handler.NewDonationHandler(midtransService, notificationService)
	//end

	return router.PublicRoutes(userHandler, bloodRequestHandler, bloodDonationHandler, certificateHandler, donationHandler, dashboardHandler)
}

func BuildPrivateRoutes(cfg *configs.Config, db *gorm.DB, rdb *redis.Client, cloudinaryService *cloudinary.Service, mailer *mailer.Mailer, blockchain service.BlockchainService) []route.Route {
	cacheable := cache.NewCacheable(rdb)
	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)

	//repository
	userRepository := repository.NewUserRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	healthPassportRepository := repository.NewHealthPassportRepository(db)
	bloodRequestRepository := repository.NewBloodRequestRepository(db)
	donorRegistrationRepository := repository.NewDonorRegistrationRepository(db)
	donorScheduleRepository := repository.NewDonorScheduleRepository(db)
	hospitalRepository := repository.NewHospitalRepository(db)
	bloodDonationRepository := repository.NewBloodDonationRepository(db)
	certificateRepository := repository.NewCertificateRepository(db)
	donationsRepository := repository.NewDonationsRepository(db)
	//end

	//service
	userService := service.NewUserService(userRepository, tokenUseCase, cacheable, cfg, mailer,cloudinaryService)
	notificationService := service.NewNotificationService(notificationRepository,userRepository)
	healthPassportService := service.NewHealthPassportService(healthPassportRepository)
	bloodRequestService := service.NewBloodRequestService(bloodRequestRepository, *cloudinaryService)
	donorRegistrationService := service.NewDonorRegistrationService(donorRegistrationRepository)
	donorScheduleService := service.NewDonorScheduleService(donorScheduleRepository)
	hospitalService := service.NewHospitalService(hospitalRepository)
	bloodDonationService := service.NewBloodDonationService(bloodDonationRepository, *cloudinaryService)
	certificateService := service.NewCertificateService(certificateRepository)

	// Buat instance midtransService
	midtransService := midtrans.NewMidtransService(&cfg.MidtransConfig)
	// Set donationsRepository
	midtransService.DonationsRepository = donationsRepository
	googleAuthService := googleoauth.NewGoogleOAuthService(tokenUseCase, userService, &cfg.GoogleOauth)

	dashboardService := service.NewDashboardService(bloodDonationRepository, bloodRequestRepository, userRepository)

	//end

	//handler
	userHandler := handler.NewUserHandler(userService, cloudinaryService, googleAuthService)
	notificationHandler := handler.NewNotificationHandler(notificationService)
	healthPassportHandler := handler.NewHealthPassportHandler(healthPassportService)
	bloodRequestHandler := handler.NewBloodRequestHandler(bloodRequestService, notificationService)
	donorRegistrationHandler := handler.NewDonorRegistrationHandler(donorRegistrationService, healthPassportService, notificationService, bloodRequestService)

	donorScheduleHandler := handler.NewDonorScheduleHandler(donorScheduleService)
	hospitalHandler := handler.NewHospitalHandler(hospitalService)
	bloodDonationHandler := handler.NewBloodDonationHandler(bloodDonationService, notificationService, certificateService, donorRegistrationService, userService, blockchain)

	certificateHandler := handler.NewCertificateHandler(certificateService)
	donationHandler := handler.NewDonationHandler(midtransService, notificationService)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	//end

	return router.PrivateRoutes(userHandler, notificationHandler, healthPassportHandler, bloodRequestHandler, donorRegistrationHandler, donorScheduleHandler, hospitalHandler, bloodDonationHandler, certificateHandler, donationHandler, dashboardHandler)
}
