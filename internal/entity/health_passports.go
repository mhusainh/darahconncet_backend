package entity

import "time"

type HealthPassport struct {
	Id             int64     `json:"id"`
	UserId         int64     `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId;references:Id" json:"user"` // Add this line to embed the User entity
	PassportNumber string    `json:"passport_number"` // Unique identifier for the health passport
	ExpiryDate     time.Time `json:"expiry_date"`     // Expiry date of the health passport
	Status         string    `json:"status"`          // e.g., "active", "expired", "revoked"
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (HealthPassport) TableName() string {
	return "public.health_passports"
}
