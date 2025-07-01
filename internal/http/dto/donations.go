package dto

type DonationsCreate struct {
	UserID int64 `json:"user_id" form:"user_id" validate:"required"`
	Amount int64 `json:"amount" form:"amount" validate:"required"`
	Transaction_time string `json:"transaction_time" form:"transaction_time" validate:"required"`
	Transaction_status string `json:"transaction_status" form:"transaction_status" validate:"required"`
}

type PaymentRequest struct {
	OrderID  string `json:"order_id" form:"order_id"`
	Amount   int64 `json:"amount" form:"amount" validate:"required"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}
