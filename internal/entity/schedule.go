package entity

import "time"

type Schedule struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	DoctorID  int       `json:"doctor_id"`
	Date      time.Time `json:"date"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	IsBooked  bool      `json:"is_booked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Doctor *DoctorProfile `json:"doctor,omitempty" gorm:"foreignKey:DoctorID"`
}
