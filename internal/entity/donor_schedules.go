package entity

import "time"

type DonorSchedule struct {
	Id             int64             `json:"id"`
	UserId         int64             `json:"user_id"`
	RequestId      int64             `json:"request_id"`
	BloodRequest   BloodRequest      `gorm:"foreignKey:RequestId;references:Id" json:"BloodRequest"`
	HospitalId     int64             `json:"hospital_id"`
	Hospital       Hospital          `gorm:"foreignKey:HospitalId;references:Id" json:"Hospital"` // Add this line to embed the Hospital entity
	Description    string            `json:"description"`
	Status         string            `json:"status"` //'Upcoming', 'Ongoing', 'Completed', 'Cancelled'
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

func (DonorSchedule) TableName() string {
	return "public.donor_schedules"
}
