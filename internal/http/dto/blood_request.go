package dto

import (
	"mime/multipart"
	"time"
)

type BloodRequestCreateRequest struct {
	UserId       int64                 `json:"user_id" form:"user_id"`
	HospitalId   int64                 `json:"hospital_id" form:"hospital_id" validate:"required"`
	PatientName  string                `json:"patient_name" form:"patient_name" validate:"required"`
	EventName    string                `json:"event_name" form:"event_name"`
	EventDate    time.Time             `json:"event_date" form:"event_date"`
	BloodType    string                `json:"blood_type" form:"blood_type"` // Unique identifier for the health passport
	Quantity     int64                 `json:"quantity" form:"quantity"`
	UrgencyLevel string                `json:"urgency_level" form:"urgency_level"` // Unique identifier for the health passport
	Diagnosis    string                `json:"diagnosis" form:"diagnosis"`         // Unique identifier for the health passport
	Image        *multipart.FileHeader `json:"image" form:"image"`
}

type CampaignCreateRequest struct {
	UserId         int64                 `json:"user_id" form:"user_id"`
	HospitalId     int64                 `json:"hospital_id" form:"hospital_id" validate:"required"`
	EventName      string                `json:"event_name" form:"event_name"`
	EventDate      time.Time             `json:"event_date" form:"event_date"`
	StartTime      time.Time             `json:"start_time" form:"start_time"`
	EndTime        time.Time             `json:"end_time" form:"end_time"`
	SlotsAvailable int64                 `json:"slots_available" form:"slots_available"`
	SlotsBooked    int64                 `json:"slots_booked" form:"slots_booked"`
	Image          *multipart.FileHeader `json:"image" form:"image"`
}

type BloodRequestUpdateRequest struct {
	Id           int64                 `param:"id" validate:"required"`
	PatientName  string                `json:"patient_name" form:"patient_name"`
	EventName    string                `json:"event_name" form:"event_name"`
	EventDate    time.Time             `json:"event_date" form:"event_date"`
	BloodType    string                `json:"blood_type" form:"blood_type"` // Unique identifier for the health passport
	Quantity     int64                 `json:"quantity" form:"quantity"`
	UrgencyLevel string                `json:"urgency_level" form:"urgency_level"` // Unique identifier for the health passport
	Diagnosis    string                `json:"diagnosis" form:"diagnosis"`         // Unique identifier for the health passport
	Status       string                `json:"status" form:"status"`
	Image        *multipart.FileHeader `json:"image" form:"image"`
}

type CampaignUpdateRequest struct {
	Id             int64                 `param:"id" validate:"required"`
	EventName      string                `json:"event_name" form:"event_name"`
	EventDate      time.Time             `json:"event_date" form:"event_date"`
	StartTime      time.Time             `json:"start_time" form:"start_time"`
	EndTime        time.Time             `json:"end_time" form:"end_time"`
	Image          *multipart.FileHeader `json:"image" form:"image"`
	SlotsAvailable int64                 `json:"slots_available" form:"slots_available"`
	SlotsBooked    int64                 `json:"slots_booked" form:"slots_booked"`
}

type BloodRequestByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type BloodRequestByUserIdRequest struct {
	UserId int64 `param:"user_id" validate:"required"`
}

type BloodRequestByHospitalIdRequest struct {
	HospitalId int64 `param:"hospital_id" validate:"required"`
}


type GetAllBloodRequestRequest struct {
	Page         int64  `query:"page"`
	Limit        int64  `query:"limit"`
	Search       string `query:"search"`
	Sort         string `query:"sort"`
	Order        string `query:"order"`
	StartDate    string `query:"start_date"`
	EndDate      string `query:"end_date"`
	MaxQuantity  int64  `query:"max_quantity"`
	MinQuantity  int64  `query:"min_quantity"`
	UrgencyLevel string `query:"urgency_level"`
	BloodType    string `query:"blood_type"`
	EventType    string `query:"event_type"`
}