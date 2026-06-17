package entity

import "time"

type Payment struct {
	ID            int        `json:"id" gorm:"primaryKey"`
	AppointmentID int        `json:"appointment_id"`
	Amount        int        `json:"amount"`
	Status        string     `json:"status"`
	TransactionID int        `json:"transaction_id"`
	PaidAt        *time.Time `json:"paid_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`

	Appointment Appointment `json:"appointment" gorm:"foreignKey:AppointmentID"`
}
