package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/response"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
)

type NotificationHandler struct {
	notificationService service.NotificationService
}

func NewNotificationHandler(
	notificationService service.NotificationService,
	) NotificationHandler {
	return NotificationHandler{
		notificationService,
	}
}

func (h *NotificationHandler) GetNotifications(ctx echo.Context) error {
	// Bind query parameters ke struct request
	var req dto.GetAllNotificationRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	notifications, total, err := h.notificationService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, 
		response.SuccessResponseWithPagi("successfully showing all notifications", notifications, req.Page, req.Limit, total))
}

func (h *NotificationHandler) GetNotification(ctx echo.Context) error {
	var req dto.NotificationByIdRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	notification, err := h.notificationService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing notification", notification))
}

func (h *NotificationHandler) GetNotificationByUserId(ctx echo.Context) error{
	// Bind query parameters ke struct request
	var userReq dto.NotificationByUserIdRequest
	if err := ctx.Bind(&userReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Bind query parameters untuk filter dan pagination
	var req dto.GetAllNotificationRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	notifications, total, err := h.notificationService.GetByUserId(ctx.Request().Context(), userReq.UserId, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, 
		response.SuccessResponseWithPagi("successfully showing all notifications", notifications, req.Page, req.Limit, total,))
}

func (h *NotificationHandler) GetNotificationsByUser(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user claims"))
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user information from claims"))
	}

	// Bind query parameters untuk filter dan pagination
	var req dto.GetAllNotificationRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	notifications, total, err := h.notificationService.GetByUserId(ctx.Request().Context(), claimsData.Id, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, 
		response.SuccessResponseWithPagi("successfully showing all notifications by user", notifications, req.Page, req.Limit, total))
}

func (h *NotificationHandler) GetNotificationByUser(ctx echo.Context) error {
	var req dto.NotificationByIdRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user claims"))
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user information from claims"))
	}

	notification, err := h.notificationService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// Check if the notification belongs to the user
	if notification.UserId != claimsData.Id {
		return ctx.JSON(http.StatusForbidden, response.ErrorResponse(http.StatusForbidden, "unauthorized access"))
	}
	
	var update dto.NotificationUpdateRequest
	update.IsRead = true
	if err := h.notificationService.Update(ctx.Request().Context(), update, notification); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing notification by user", notification))
}

func (h *NotificationHandler) GetUnreadNotificationCount(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user claims"))
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user information from claims"))
	}

	count, err := h.notificationService.GetUnreadCountByUserId(ctx.Request().Context(), claimsData.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing unread notification count", count))
}

func (h *NotificationHandler) CreateNotification(ctx echo.Context) error {
	var req dto.NotificationCreateRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := ctx.Validate(&req); err != nil{
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if err := h.notificationService.Create(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusCreated, response.SuccessResponse("successfully created notification", nil))
}

func (h *NotificationHandler) UpdateNotification(ctx echo.Context) error {
	var req dto.NotificationUpdateRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	notification, err := h.notificationService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.notificationService.Update(ctx.Request().Context(), req, notification); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusCreated, response.SuccessResponse("successfully updated notification", nil))
}

func (h *NotificationHandler) DeleteNotification(ctx echo.Context) error {
	var req dto.NotificationByIdRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := h.notificationService.Delete(ctx.Request().Context(), req.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully deleted notification", nil))
}
