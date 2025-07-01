package entity

import "time"

type User struct {
	Id                 int64     `json:"id"`
	Name               string    `json:"name"`
	Gender             string    `json:"gender"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Phone              string    `json:"phone"`
	BloodType          string    `json:"blood_type"`
	BirthDate          time.Time `json:"birth_date"`
	Address            string    `json:"address"`
	Role               string    `json:"role"`
	ResetPasswordToken string    `json:"reset_password_token"`
	VerifyEmailToken   string    `json:"verify_email_token"`
	IsVerified         bool       `json:"is_verified"`
	PublicId           string    `json:"public_id"`
	UrlFile            string    `json:"url_file"`
	WalletAddress      string    `json:"wallet_address"`
	TokenExpiresAt     time.Time `json:"token_expires_at"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "public.users"
}
