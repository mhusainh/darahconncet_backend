package entity

import "time"

type BloodRequest struct {
	Id             int64     `json:"id"`
	UserId         int64     `json:"user_id"`
	User           User      `json:"user" gorm:"foreignKey:UserId;references:Id"`
	HospitalId     int64     `json:"hospital_id"`
	Hospital       Hospital  `json:"hospital" gorm:"foreignKey:HospitalId;references:Id"`
	PatientName    string    `json:"patient_name"`
	BloodType      string    `json:"blood_type"`
	Quantity       int64     `json:"quantity"`
	UrgencyLevel   string    `json:"urgency_level"`
	Diagnosis      string    `json:"diagnosis"`
	EventName      string    `json:"event_name"`
	EventDate      time.Time `json:"event_date"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	SlotsAvailable int64     `json:"slots_available"`
	SlotsBooked    int64     `json:"slots_booked"`
	Status         string    `json:"status"` //'Pending', 'Verified', 'Fulfilled', 'Cancelled', 'Expired'
	EventType      string    `json:"event_type"`
	UrlFile        string    `json:"url_file"`
	PublicId       string    `json:"public_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (BloodRequest) TableName() string {
	return "public.blood_requests"
}
