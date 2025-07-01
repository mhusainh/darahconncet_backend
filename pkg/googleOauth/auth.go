package googleoauth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"
)

type GoogleAuthService interface {
	Login(ctx echo.Context) error
	Callback(ctx echo.Context) error
}
type Service struct {
	tokenService token.TokenUseCase
	userService  service.UserService
	cfg          *configs.GoogleOauth
}

func NewGoogleOAuthService(tokenService token.TokenUseCase, userService service.UserService, cfg *configs.GoogleOauth) *Service {
	return &Service{
		tokenService,
		userService,
		cfg,
	}
}

func InitGoogle(cfg *configs.GoogleOauth) error {
	if cfg.ClientId == "" || cfg.ClientSecret == "" || cfg.CallbackURL == "" {
		return errors.New("Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
	}

	log.Printf("Initializing Google OAuth with callback URL: %s", cfg.CallbackURL)

	goth.UseProviders(
		google.New(cfg.ClientId, cfg.ClientSecret, cfg.CallbackURL),
	)
	return nil
}

func (s *Service) Login(ctx echo.Context) (goth.User, error) {
	provider := ctx.Param("provider")
	if provider == "" {
		return goth.User{}, errors.New("Provider not specified")
	}

	// Clear any existing session first
	gothic.Logout(ctx.Response().Writer, ctx.Request())

	q := ctx.Request().URL.Query()
	q.Add("provider", provider)
	ctx.Request().URL.RawQuery = q.Encode()

	req := ctx.Request()
	res := ctx.Response().Writer

	log.Printf("Starting OAuth login for provider: %s", provider)
	log.Printf("Request URL: %s", req.URL.String())

	// Begin authentication process - this will redirect to Google
	gothic.BeginAuthHandler(res, req)

	// This won't be reached due to redirect, but we need to return something
	return goth.User{}, nil
}

func (s *Service) Callback(ctx echo.Context) (string, error) {
	// Add provider parameter to query for goth
	provider := ctx.Param("provider")
	if provider == "" {
		return "", errors.New("Provider not specified")
	}

	q := ctx.Request().URL.Query()
	q.Add("provider", provider)
	ctx.Request().URL.RawQuery = q.Encode()

	req := ctx.Request()
	res := ctx.Response().Writer

	log.Printf("Processing OAuth callback for provider: %s", provider)
	log.Printf("Callback URL: %s", req.URL.String())
	log.Printf("Query parameters: %s", req.URL.RawQuery)

	// Add debugging for the authorization code
	authCode := req.URL.Query().Get("code")
	if authCode == "" {
		log.Printf("No authorization code found in callback")
		return "", errors.New("authorization code not found")
	}
	log.Printf("Authorization code received (length: %d)", len(authCode))

	// Check for error parameter
	if errParam := req.URL.Query().Get("error"); errParam != "" {
		log.Printf("OAuth error parameter: %s", errParam)
		return "", fmt.Errorf("OAuth error: %s", errParam)
	}

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Printf("Error completing user auth: %v", err)
		return "", fmt.Errorf("authentication failed: %v", err)
	}

	log.Printf("Successfully authenticated user: %s", user.Email)

	// Check if user already exists in the database
	metamask := false
	userEntity, IsNew, err := s.userService.CheckGoogleOAuth(ctx.Request().Context(), user.Email, &user)
	if err != nil {
		log.Printf("Error checking Google OAuth user: %v", err)
		return "", errors.New("ada kesalahan saat check google oauth")
	}
	if userEntity.WalletAddress != "" {
		metamask = true
	}

	// Buat JWT claims dari data Google OAuth
	claims := &token.GoogleOAuthClaims{
		Id:         userEntity.Id,
		Email:      user.Email,
		Name:       userEntity.Name,
		Role:       userEntity.Role,
		PictureURL: user.AvatarURL,
		Provider:   "google",
		Metamask:   metamask,
		IsNew:      IsNew,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate token
	accessToken, err := s.tokenService.GenerateAccessToken(claims)
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		return "", errors.New("ada kesalahan saat generate token")
	}
	// Return token dan data user
	return s.cfg.RedirectURL + accessToken, nil
}