package handler

import (
	"net/http"
	"strconv"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
	googleoauth "github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/googleOauth"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/response"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService        service.UserService
	cloudinaryService  *cloudinary.Service
	GoogleOauthService *googleoauth.Service
}

func NewUserHandler(userService service.UserService, cloudinaryService *cloudinary.Service, GoogleOauthService *googleoauth.Service) UserHandler {
	return UserHandler{userService, cloudinaryService, GoogleOauthService}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	// Untuk endpoint GET, gunakan query parameters
	var req dto.GetAllUserRequest
	
	// Ambil query parameters secara manual
	if page := ctx.QueryParam("page"); page != "" {
		pageInt, err := strconv.ParseInt(page, 10, 64)
		if err == nil {
			req.Page = pageInt
		}
	}
	
	if limit := ctx.QueryParam("limit"); limit != "" {
		limitInt, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			req.Limit = limitInt
		}
	}
	
	// Ambil query parameters string langsung
	req.Email = ctx.QueryParam("email")
	req.Search = ctx.QueryParam("search")
	req.Sort = ctx.QueryParam("sort")
	req.Order = ctx.QueryParam("order")
	req.BloodType = ctx.QueryParam("blood_type")

	// Gunakan req yang sudah diisi dengan query parameters
	users, total, err := h.userService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponseWithPagi("successfully showing all users", users, req.Page, req.Limit, total))
}

func (h *UserHandler) GetUser(ctx echo.Context) error {
	// Ambil ID dari parameter path
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID format"))
	}

	user, err := h.userService.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing user", user))
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var loginRequest dto.UserLoginRequest

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	token, isVerified, err := h.userService.Login(ctx.Request().Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
	}
	if !isVerified {
		return ctx.JSON(http.StatusOK, response.SuccessResponse("Email belum diverifikasi, silahkan verifikasi email anda", map[string]interface{}{
			"verify_expired_at": token,
		}))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully login", map[string]interface{}{
		"token": token,
	}))
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var req dto.UserRegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := h.userService.Register(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse("successfully registered", nil))
}

func (h *UserHandler) GetProfile(ctx echo.Context) error {

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

	user, err := h.userService.GetById(ctx.Request().Context(), claimsData.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing a user", user))
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	// Ambil ID dari parameter path
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID format"))
	}

	user, err := h.userService.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.userService.Delete(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully delete user", nil))
}

func (h *UserHandler) ResetPassword(ctx echo.Context) error {
	var req dto.ResetPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	req.Token = ctx.QueryParam("token")
	if req.Token == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "token is required"))
	}

	err := h.userService.ResetPassword(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully reset a password", nil))
}

func (h *UserHandler) ResetPasswordRequest(ctx echo.Context) error {
	var req dto.RequestResetPassword

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.userService.RequestResetPassword(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully request reset password", nil))
}

func (h *UserHandler) ResendTokenVerifyEmail(ctx echo.Context) error {
	var req dto.ResendTokenVerifyEmailRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	user, err := h.userService.ResendTokenVerifyEmail(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Resend token verify email", user))
}

func (h *UserHandler) VerifyEmail(ctx echo.Context) error {
	// Get token directly from query parameter instead of using Bind
	token := ctx.QueryParam("token")
	if token == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "token is required"))
	}

	// Create the request with the token
	req := dto.VerifyEmailRequest{
		Token: token,
	}

	err := h.userService.VerifyEmail(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully verify email", nil))
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var req dto.UpdateUserRequest

	// --- PERUBAHAN LOGIKA DIMULAI DI SINI ---

	// Langkah 1: Bind data form (non-file) terlebih dahulu.
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Langkah 2: Tangani file upload secara manual dan terpisah.
	// Ini membuat penanganan file opsional menjadi lebih eksplisit dan aman.
	if imageFile, err := ctx.FormFile("image"); err != nil {
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
		"image/jpg":  {},
	}
	if req.Image != nil {
		if _, ok := acceptedImages[req.Image.Header.Get("Content-Type")]; !ok {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "unsupported image type"))
		}
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

	req.Id = claimsData.Id

	// Update user data
	err := h.userService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse("successfully update user", nil))
}

func (h *UserHandler) LoginGoogleAuth(ctx echo.Context) error {
	// This will redirect to Google OAuth, so no JSON response needed
	_, err := h.GoogleOauthService.Login(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// This will not be reached due to redirect
	return nil
}

func (h *UserHandler) CallbackGoogleAuth(ctx echo.Context) error {
	data, err := h.GoogleOauthService.Callback(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Redirect(http.StatusTemporaryRedirect, data)
}

func (h *UserHandler) WalletAddress(ctx echo.Context) error {
	var req dto.WalletAddressRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user claims"))
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*token.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "unable to get user information from claims"))
	}

	user, err := h.userService.GetById(ctx.Request().Context(), claimsData.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.userService.WalletAddress(ctx.Request().Context(), user, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully update wallet address", nil))
}
