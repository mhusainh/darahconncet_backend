package router

import (
	"net/http"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/handler"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/route"
)

var (
	adminOnly = []string{"Administrator"}
	userOnly  = []string{"User"}
	allRoles  = []string{"Administrator", "User"}
)

func PublicRoutes(
	userHandler handler.UserHandler,
	bloodRequestHandler handler.BloodRequestHandler,
	bloodDonationHandler handler.BloodDonationHandler,
	certificateHandler handler.CertificateHandler,
	donationHandler *handler.DonationHandler,
	dashboardHandler handler.Dashboard,
) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodGet,
			Path:    "landing",
			Handler: dashboardHandler.GetLandingPage,
		},
		// User Handler
		{
			Method:  http.MethodPost,
			Path:    "login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodGet,
			Path:    "login/:provider",
			Handler: userHandler.LoginGoogleAuth,
		},
		{
			Method:  http.MethodGet,
			Path:    "login/:provider/callback",
			Handler: userHandler.CallbackGoogleAuth,
		},
		{
			Method:  http.MethodPost,
			Path:    "register",
			Handler: userHandler.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "request-reset-password",
			Handler: userHandler.ResetPasswordRequest,
		},
		{
			Method:  http.MethodPost,
			Path:    "reset-password",
			Handler: userHandler.ResetPassword,
		},
		{
			Method:  http.MethodGet,
			Path:    "verify-email",
			Handler: userHandler.VerifyEmail,
		},
		{
			Method:  http.MethodPost,
			Path:    "/resend-token-verify-email",
			Handler: userHandler.ResendTokenVerifyEmail,
		},
		// Blood Request/Campaign Handler
		{
			Method:  http.MethodGet,
			Path:    "blood-request",
			Handler: bloodRequestHandler.GetBloodRequests,
		},
		{
			Method:  http.MethodGet,
			Path:    "campaign",
			Handler: bloodRequestHandler.GetCampaigns,
		},
		{
			Method:  http.MethodGet,
			Path:    "campaign-bloodRequest/:id",
			Handler: bloodRequestHandler.GetById,
		},
		// Donation Handler
		{
			Method:  http.MethodPost,
			Path:    "donation/webhook",
			Handler: donationHandler.WebHookTransaction,
		},
		// Certificate Handler
	}
}

func PrivateRoutes(
	userHandler handler.UserHandler,
	notificationHandler handler.NotificationHandler,
	healthPassportHandler handler.HealthPassportHandler,
	bloodRequestHandler handler.BloodRequestHandler,
	donorRegistrationHandler handler.DonorRegistrationHandler,
	donorScheduleHandler handler.DonorScheduleHandler,
	hospitalHandler handler.HospitalHandler,
	bloodDonationHandler handler.BloodDonationHandler,
	certificateHandler handler.CertificateHandler,
	donationHandler *handler.DonationHandler,
	dashboardHandler handler.Dashboard,
) []route.Route {
	return []route.Route{
		// =============================================
		// USER ONLY ROUTES
		// =============================================
		// Dashboard - User Only
		{
			Method:  http.MethodGet,
			Path:    "user/dashboard",
			Handler: dashboardHandler.DashboardUser,
			Roles:   userOnly,
		},
		// Health Passport - User Only
		{
			Method:  http.MethodGet,
			Path:    "user/health-passport",
			Handler: healthPassportHandler.GetHealthPassportByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/health-passport",
			Handler: healthPassportHandler.CreateHealthPassport,
			Roles:   userOnly,
		},
		// Notification - User Only
		{
			Method:  http.MethodGet,
			Path:    "user/notifications/",
			Handler: notificationHandler.GetNotificationsByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/notifications/:id",
			Handler: notificationHandler.GetNotificationByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/notifications/count",
			Handler: notificationHandler.GetUnreadNotificationCount,
			Roles:   userOnly,
		},
		// Donor Registration - User Only
		{
			Method:  http.MethodPost,
			Path:    "user/donor-registration",
			Handler: donorRegistrationHandler.CreateDonorRegistration,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/schedules",
			Handler: donorScheduleHandler.GetDonorSchedules,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/schedule/:id",
			Handler: donorScheduleHandler.GetDonorSchedule,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/schedule/",
			Handler: donorScheduleHandler.CreateDonorSchedule,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "user/schedule/:id",
			Handler: donorScheduleHandler.UpdateDonorSchedule,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodDelete,
			Path:    "user/schedule/:id",
			Handler: donorScheduleHandler.DeleteDonorSchedule,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/donation/transaction",
			Handler: donationHandler.CreateTransaction,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/certificates",
			Handler: certificateHandler.GetByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/certificate/:id",
			Handler: certificateHandler.GetById,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/blood-request",
			Handler: bloodRequestHandler.GetBloodRequestByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/create-blood-request",
			Handler: bloodRequestHandler.CreateBloodRequest,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "user/update-blood-request/:id",
			Handler: bloodRequestHandler.UpdateBloodRequest,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "user/blood-donations",
			Handler: bloodDonationHandler.GetByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/wallet-address",
			Handler: userHandler.WalletAddress,
			Roles:   userOnly,
		},
		// Certificate - User Only
		{
			Method:  http.MethodGet,
			Path:    "user/certificates",
			Handler: certificateHandler.GetByUser,
			Roles:   userOnly,
		},
		// =============================================
		// ADMIN ONLY ROUTES
		// =============================================
		// Dasboard - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/dashboard",
			Handler: dashboardHandler.DashboardAdmin,
			Roles:   adminOnly,
		},
		// Health Passport - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/health-passports",
			Handler: healthPassportHandler.GetHealthPassports,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "admin/health-passport/:id",
			Handler: healthPassportHandler.GetHealthPassport,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "admin/health-passport/:id",
			Handler: healthPassportHandler.UpdateStatusHealthPassport,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodDelete,
			Path:    "admin/health-passport/:id",
			Handler: healthPassportHandler.DeleteHealthPassport,
			Roles:   adminOnly,
		},
		// Blood Request - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/blood-requests",
			Handler: bloodRequestHandler.GetBloodRequestsByAdmin,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "admin/blood-request/:id",
			Handler: bloodRequestHandler.StatusBloodRequest,
			Roles:   adminOnly,
		},
		// Blood Request/Campaign - Admin Only
		{
			Method:  http.MethodPost,
			Path:    "admin/campaign",
			Handler: bloodRequestHandler.CreateCampaign,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "admin/campaign/:id",
			Handler: bloodRequestHandler.UpdateCampaign,
			Roles:   adminOnly,
		},
		// Notification - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/notifications",
			Handler: notificationHandler.GetNotifications,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "admin/notification/:id",
			Handler: notificationHandler.GetNotification,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "admin/notifications/user/:user_id",
			Handler: notificationHandler.GetNotificationByUserId,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "admin/notification",
			Handler: notificationHandler.CreateNotification,
			Roles:   adminOnly,
		},
		// User Management - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/users",
			Handler: userHandler.GetUsers,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "admin/users/:id",
			Handler: userHandler.GetUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodDelete,
			Path:    "admin/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "blood-donations",
			Handler: bloodDonationHandler.GetAll,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "blood-donation/:id/status",
			Handler: bloodDonationHandler.StatusBloodDonation,
			Roles:   adminOnly,
		},
		// Hospital - Admin Only
		{
			Method:  http.MethodGet,
			Path:    "admin/hospital/:id",
			Handler: hospitalHandler.Update,
			Roles:   adminOnly,
		},
		// =============================================
		// ALL ROLES ROUTES (Admin & User)
		// =============================================
		// User Profile - All Roles
		{
			Method:  http.MethodGet,
			Path:    "user/profile",
			Handler: userHandler.GetProfile,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPut,
			Path:    "user/profile",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPost,
			Path:    "user/profile/picture",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},
		// Donor Registration - All Roles
		{
			Method:  http.MethodGet,
			Path:    "donor-registrations",
			Handler: donorRegistrationHandler.GetDonorRegistrations,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodGet,
			Path:    "donor-registration/:id",
			Handler: donorRegistrationHandler.GetDonorRegistration,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPut,
			Path:    "donor-registration/",
			Handler: donorRegistrationHandler.UpdateDonorRegistration,
			Roles:   allRoles,
		},
		// Hospital - All Roles
		{
			Method:  http.MethodGet,
			Path:    "hospital",
			Handler: hospitalHandler.GetAll,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodGet,
			Path:    "hospital/:id",
			Handler: hospitalHandler.GetById,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPost,
			Path:    "hospital",
			Handler: hospitalHandler.Create,
			Roles:   allRoles,
		},
		// Blood Donation - All Roles
		{
			Method:  http.MethodGet,
			Path:    "blood-donation/:id",
			Handler: bloodDonationHandler.GetById,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPost,
			Path:    "blood-donation",
			Handler: bloodDonationHandler.Create,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodDelete,
			Path:    "blood-donation/:id",
			Handler: bloodDonationHandler.Delete,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodDelete,
			Path:    "campaign/:id",
			Handler: bloodRequestHandler.DeleteBloodRequest,
			Roles:   allRoles,
		},
	}
}
