package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/response"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"

	"github.com/labstack/echo/v4"
)

type BloodRequestHandler struct {
	bloodRequestService service.BloodRequestService
	notificationService service.NotificationService
}

func NewBloodRequestHandler(
	bloodRequestService service.BloodRequestService,
	notificationService service.NotificationService,
	) BloodRequestHandler {
	return BloodRequestHandler{
		bloodRequestService,
		notificationService,
	}
}

func (h *BloodRequestHandler) CreateBloodRequest(ctx echo.Context) error {
	var req dto.BloodRequestCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if imageFile, err := ctx.FormFile("image");err != nil {
		// Jika error bukan karena file tidak ada, berarti ada masalah lain.
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "error processing image file: "+err.Error()))
		}
		// Jika errornya adalah http.ErrMissingFile, tidak apa-apa, karena gambar bersifat opsional.
		// req.Image akan tetap nil.
	} else {
		// Jika file ada, masukkan ke dalam struct request.
		req.Image = imageFile
	}
	var acceptedImages = map[string]struct{}{
		"image/png":  {},
		"image/jpeg": {},
		"image/jpg": {},
	}
	if req.Image != nil {
		if _, ok := acceptedImages[req.Image.Header.Get("Content-Type")]; !ok {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "unsupported image type"))
		}
	}	

	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	req.UserId = claimsData.Id

	if err := h.bloodRequestService.CreateBloodRequest(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	notificationData := dto.NotificationCreateRequest{
		UserId:      req.UserId,
		Title:       "Pemberitahuan Permintaan Darah",
		Message:	 "Anda telah membuat permintaan darah. Mohon tunggu konfirmasi dari pihak admin.",
		NotificationType: "Request",
	}
	if err := h.notificationService.Create(ctx.Request().Context(),notificationData ); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully create blood request", nil))
}

func (h *BloodRequestHandler) CreateCampaign(ctx echo.Context) error {
	var req dto.CampaignCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	if imageFile, err := ctx.FormFile("image");err != nil {
		// Jika error bukan karena file tidak ada, berarti ada masalah lain.
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "error processing image file: "+err.Error()))
		}
		// Jika errornya adalah http.ErrMissingFile, tidak apa-apa, karena gambar bersifat opsional.
		// req.Image akan tetap nil.
	} else {
		// Jika file ada, masukkan ke dalam struct request.
		req.Image = imageFile
	}
	var acceptedImages = map[string]struct{}{
		"image/png":  {},
		"image/jpeg": {},
		"image/jpg": {},
	}
	if req.Image != nil {
		if _, ok := acceptedImages[req.Image.Header.Get("Content-Type")]; !ok {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "unsupported image type"))
		}
	}
	
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	req.UserId = claimsData.Id

	if err := h.bloodRequestService.CreateCampaign(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully create campaign", nil))
}

func (h *BloodRequestHandler) GetBloodRequests(ctx echo.Context) error {
	var req dto.GetAllBloodRequestRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	bloodRequests, total, err := h.bloodRequestService.GetAllBloodRequest(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully showing all blood requests", bloodRequests, req.Page, req.Limit, total))
}

func (h *BloodRequestHandler) GetBloodRequestByUser(ctx echo.Context) error {
	var req dto.GetAllBloodRequestRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	bloodRequests, total, err := h.bloodRequestService.GetAllBloodRequestByUser(ctx.Request().Context(), claimsData.Id, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully showing all blood requests", bloodRequests, req.Page, req.Limit, total))
}

func (h *BloodRequestHandler) GetBloodRequestsByAdmin(ctx echo.Context) error {
	var req dto.GetAllBloodRequestRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	bloodRequests, total, err := h.bloodRequestService.GetAllAdminBloodRequest(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully showing all blood requests", bloodRequests, req.Page, req.Limit, total))
}

func (h *BloodRequestHandler) GetCampaigns(ctx echo.Context) error {
	var req dto.GetAllBloodRequestRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	bloodRequests, total, err := h.bloodRequestService.GetAllCampaign(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully showing all blood requests", bloodRequests, req.Page, req.Limit, total))
}

func (h *BloodRequestHandler) GetById(ctx echo.Context) error {
	var req dto.BloodRequestByIdRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	bloodRequest, err := h.bloodRequestService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing blood request", bloodRequest))
}

func (h *BloodRequestHandler) UpdateBloodRequest(ctx echo.Context) error {
	var req dto.BloodRequestUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if imageFile, err := ctx.FormFile("image");err != nil {
		// Jika error bukan karena file tidak ada, berarti ada masalah lain.
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "error processing image file: "+err.Error()))
		}
		// Jika errornya adalah http.ErrMissingFile, tidak apa-apa, karena gambar bersifat opsional.
		// req.Image akan tetap nil.
	} else {
		// Jika file ada, masukkan ke dalam struct request.
		req.Image = imageFile
	}
	var acceptedImages = map[string]struct{}{
		"image/png":  {},
		"image/jpeg": {},
		"image/jpg": {},
	}
	if req.Image != nil {
		if _, ok := acceptedImages[req.Image.Header.Get("Content-Type")]; !ok {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "unsupported image type"))
		}
	}

	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	bloodRequest, err := h.bloodRequestService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	if claimsData.Role == "User" {
		if claimsData.Id != bloodRequest.UserId {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "You are not authorized to update this request"))
		}
		if bloodRequest.Status == "Completed" || bloodRequest.Status == "Verified" {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Permintaan Sudah tidak bisa diupdate"))
		}
		if req.Status != "Canceled" {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Kamu hanya bisa membatalkan permintaan"))
		}
	}

	if err := h.bloodRequestService.UpdateBloodRequest(ctx.Request().Context(), req, bloodRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully updating blood request", bloodRequest))
}

func (h *BloodRequestHandler) UpdateCampaign(ctx echo.Context) error {
	var req dto.CampaignUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if imageFile, err := ctx.FormFile("image");err != nil {
		// Jika error bukan karena file tidak ada, berarti ada masalah lain.
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "error processing image file: "+err.Error()))
		}
		// Jika errornya adalah http.ErrMissingFile, tidak apa-apa, karena gambar bersifat opsional.
		// req.Image akan tetap nil.
	} else {
		// Jika file ada, masukkan ke dalam struct request.
		req.Image = imageFile
	}
	var acceptedImages = map[string]struct{}{
		"image/png":  {},
		"image/jpeg": {},
		"image/jpg": {},
	}
	if req.Image != nil {
		if _, ok := acceptedImages[req.Image.Header.Get("Content-Type")]; !ok {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "unsupported image type"))
		}
	}

	bloodRequest, err := h.bloodRequestService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.bloodRequestService.UpdateCampaign(ctx.Request().Context(), req, bloodRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully updating campaign", bloodRequest))
}

func (h *BloodRequestHandler) StatusBloodRequest(ctx echo.Context) error {
	var req dto.BloodRequestUpdateRequest
	var notif dto.NotificationCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	bloodRequest, err := h.bloodRequestService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.bloodRequestService.UpdateBloodRequest(ctx.Request().Context(), req, bloodRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	notif.UserId = bloodRequest.UserId
	notif.Title = "Permintaan Darah"
	notif.Message = "Permintaan darah anda telah di" + req.Status
	notif.NotificationType = "info"
	if err := h.notificationService.Create(ctx.Request().Context(), notif); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully creating blood request", nil))
}

func (h *BloodRequestHandler) DeleteBloodRequest(ctx echo.Context) error {
	var req dto.BloodRequestByIdRequest
	// Mengambil ID dari parameter URL, bukan dari body request
	idStr := ctx.Param("id")
	if idStr == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "id parameter is required"))
	}
	
	// Parse ID ke int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid id format"))
	}
	
	req.Id = id

	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	bloodRequest, err := h.bloodRequestService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	if claimsData.Role == "User" {
		if claimsData.Id != bloodRequest.UserId {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Anda tdak mempunyai akses"))
		}

		if bloodRequest.Status != "pending" {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "campaign tidak dapat dihapus karena status " + bloodRequest.Status))
		}
	}
	if err := h.bloodRequestService.Delete(ctx.Request().Context(), req.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully deleting blood request", nil))
}
