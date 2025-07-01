package entity

import "time"

type Notification struct {
	Id               int64     `json:"id"`
	UserId           int64     `json:"user_id"`
	User             User      `gorm:"foreignKey:UserId;references:Id" json:"user"` // Add this line to embed the User entity
	Title            string    `json:"title"`
	Message          string    `json:"message"`
	NotificationType string    `json:"notification_type"` // e.g., "alert", "reminder", "info"
	IsRead           bool      `json:"is_read"`           // Indicates if the notification has been read
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (Notification) TableName() string {
	return "public.notifications"
}
