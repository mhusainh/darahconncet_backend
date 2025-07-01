package dto

import (
	"mime/multipart"
	"time"
)

type BloodDonationCreateRequest struct {
	UserId         int64                 `json:"user_id" form:"user_id"`
	HospitalId     int64                 `json:"hospital_id" form:"hospital_id" validate:"required"`
	RegistrationId int64                 `json:"registration_id" form:"registration_id" validate:"required"`
	DonationDate   time.Time             `json:"donation_date" form:"donation_date" validate:"required"`
	BloodType      string                `json:"blood_type" form:"blood_type" validate:"required"` // e.g., "A+", "O-", etc.
	Status         string                `json:"status" form:"status" validate:"required"`
	Image          *multipart.FileHeader `json:"image" form:"image" validate:"required"`
}

type BloodDonationUpdateRequest struct {
	Id           int64                 `param:"id" validate:"required"`
	DonationDate time.Time             `json:"donation_date" form:"donation_date"`
	BloodType    string                `json:"blood_type" form:"blood_type"` // e.g., "A+", "O-", etc.
	Status       string                `json:"status" form:"status"`         //'Completed', 'Rejected', 'Deferred'
	Image        *multipart.FileHeader `form:"image" validate:"omitempty"`
}

type BloodDonationByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type GetAllBloodDonationRequest struct {
	Page      int64  `query:"page"`
	Limit     int64  `query:"limit"`
	Search    string `query:"search"`
	Sort      string `query:"sort"`
	Order     string `query:"order"`
	Status    string `query:"status"`
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
	BloodType string `query:"blood_type"`
}
