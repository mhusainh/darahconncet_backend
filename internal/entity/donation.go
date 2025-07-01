package entity

import (
	"time"
)

type Donation struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	UserId    int64     `gorm:"not null" json:"user_id"`
	Amount    int64     `gorm:"not null" json:"amount"`
	Status    string    `gorm:"not null" json:"status"`
	TransactionTime time.Time `gorm:"not null" json:"transaction_time"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Donation) TableName() string {
	return "public.donations"
}
