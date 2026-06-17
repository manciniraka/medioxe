package entity

import "time"

type AppointmentHistory struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	AppointmentID int       `json:"appointment_id"`
	Status        string    `json:"status"`
	Remarks       string    `json:"remarks"`
	CreatedAt     time.Time `json:"created_at"`

	Appointment Appointment `json:"appointment" gorm:"foreignKey:AppointmentID"`
}
