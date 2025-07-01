package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims jwt.Claims) (string, error)
}

type tokenUseCase struct {
	secretKey string
}

func NewTokenUseCase(secretKey string) TokenUseCase {
	return &tokenUseCase{secretKey}
}

type JwtCustomClaims struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Metamask bool   `json:"metamask"`
	jwt.RegisteredClaims
}

type ResetPasswordClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type GoogleOAuthClaims struct {
	Id         int64 `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	PictureURL string `json:"picture_url,omitempty"`
	Provider   string `json:"provider"` // "google"
	Metamask   bool   `json:"metamask"`
	IsNew      bool   `json:"is_new"`
	jwt.RegisteredClaims
}

func (t *tokenUseCase) GenerateAccessToken(claims jwt.Claims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
