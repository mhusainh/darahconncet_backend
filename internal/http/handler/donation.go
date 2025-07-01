package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/response"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/midtrans"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
)

type DonationHandler struct {
	midtransService midtrans.MidtransService
	notificationService service.NotificationService
}

func NewDonationHandler(midtransService midtrans.MidtransService, notificationService service.NotificationService) *DonationHandler {
	return &DonationHandler{
		midtransService: midtransService,
		notificationService: notificationService,
	}
}


func (h *DonationHandler) WebHookTransaction(ctx echo.Context) error {
	var req dto.DonationsCreate
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.midtransService.WebHookTransaction(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully update donation status", nil))
}

func (h *DonationHandler) CreateTransaction(ctx echo.Context) error {
	var req dto.PaymentRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}
	// Generate order ID
	req.OrderID = "ORDER-" + strconv.FormatInt(claimsData.Id, 10) + "-" + time.Now().Format("20060102150405")
	req.Fullname = claimsData.Name
	req.Email = claimsData.Email

	redirectURL, err := h.midtransService.CreateTransaction(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	notificationData := dto.NotificationCreateRequest{
		UserId:      claimsData.Id,
		Title:       "Pemberitahuan Donasi",
		Message:     "Terima kasih telah berdonasi. Sialhkan cek cara pembayaran di email anda.",
		NotificationType: "Donation",
	}
	if err := h.notificationService.Create(ctx.Request().Context(), notificationData); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully created transaction", map[string]interface{}{
		"redirect_url": redirectURL,
	}))
}
