package dto

type NotificationCreateRequest struct {
	UserId           int64  `json:"user_id" form:"user_id" validate:"required"`
	Title            string `json:"title" form:"title" validate:"required"`
	Message          string `json:"message" form:"message" validate:"required"`
	NotificationType string `json:"notification_type" form:"notification_type" validate:"required"` // 'Request', 'Donation', 'Certificate', 'Reminder', 'System'
}

type NotificationUpdateRequest struct {
	Id               int64  `param:"id" validate:"required"`
	Title            string `json:"title" form:"title"`
	Message          string `json:"message" form:"message"`
	NotificationType string `json:"notification_type" form:"notification_type"` // 'Request', 'Donation', 'Certificate', 'Reminder', 'System'
	IsRead           bool   `json:"is_read" form:"is_read"`           // Indicates if the notification has been read
}

type NotificationByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type NotificationByUserIdRequest struct {
	UserId int64 `param:"user_id" validate:"required"`
}

type GetAllNotificationRequest struct {
	Page      int64  `query:"page" `
	Limit     int64  `query:"limit" `
	Search    string `query:"search"`
	Sort      string `query:"sort"`
	Order     string `query:"order"`
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
	IsRead    *bool  `query:"is_read"`
}
