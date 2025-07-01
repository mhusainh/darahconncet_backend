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

type HealthPassportHandler struct {
	healthPassportService service.HealthPassportService
}

func NewHealthPassportHandler(healthPassportService service.HealthPassportService) HealthPassportHandler {
	return HealthPassportHandler{
		healthPassportService,
	}
}

func (h *HealthPassportHandler) GetHealthPassports(ctx echo.Context) error {
	var req dto.GetAllHealthPassportRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	healthPassport, total, err := h.healthPassportService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully get health passport", healthPassport, total, req.Page, req.Limit))
}

func (h *HealthPassportHandler) GetHealthPassport(ctx echo.Context) error {
	var req dto.HealthPassportByIdRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	healthPassport, err := h.healthPassportService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing health passport", healthPassport))
}

func (h *HealthPassportHandler) GetHealthPassportByUser(ctx echo.Context) error {
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

	healthPassport, err := h.healthPassportService.GetByUserId(ctx.Request().Context(), claimsData.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK,
		response.SuccessResponse("successfully showing health passport", healthPassport))
}

func (h *HealthPassportHandler) CreateHealthPassport(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user claims"))
	}
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user information from claims"))
	}

	healthPassport,_ := h.healthPassportService.GetByUserId(ctx.Request().Context(), claimsData.Id)
	
	if healthPassport != nil {
		if err := h.healthPassportService.UpdateByUser(ctx.Request().Context(), healthPassport); err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully update health passport", nil))
	}

	if err := h.healthPassportService.Create(ctx.Request().Context(), claimsData.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusCreated, response.SuccessResponse("successfully create health passport", nil))
}

func (h *HealthPassportHandler) UpdateStatusHealthPassport(ctx echo.Context) error {
	var req dto.HealthPassportUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	healthPassport, err := h.healthPassportService.GetById(ctx.Request().Context(), req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.healthPassportService.Update(ctx.Request().Context(), req, healthPassport); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully update health passport", nil))
}

func (h *HealthPassportHandler) DeleteHealthPassport(ctx echo.Context) error {
	var req dto.HealthPassportByIdRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := h.healthPassportService.Delete(ctx.Request().Context(), req.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully delete health passport", nil))
}
