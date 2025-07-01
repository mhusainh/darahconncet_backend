package dto

type DonorScheduleCreateRequest struct {
	UserId         int64     `json:"user_id" form:"user_id" validate:"required"`
	RequestId      int64     `json:"request_id" form:"request_id"`
	HospitalId     int64     `json:"hospital_id" form:"hospital_id" validate:"required"`
	Description    string    `json:"description" form:"description" validate:"required"`
}

type DonorScheduleUpdateRequest struct {
	Id             int64     `param:"id" validate:"required"`
	Description    string    `json:"description" form:"description"`
	Status         string    `json:"status" form:"status"`
}

type DonorScheduleByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type GetAllDonorScheduleRequest struct {
	Page           int64  `query:"page"`
	Limit          int64  `query:"limit"`
	Search         string `query:"search"`
	Sort           string `query:"sort"`
	Order          string `query:"order"`
	StartDate      string `query:"start_date"`
	EndDate        string `query:"end_date"`
	Status         string `query:"status"`
}