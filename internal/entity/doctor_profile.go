package entity

import "time"

type DoctorProfile struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	UserID          int       `json:"user_id"`
	SpecialtyID     int       `json:"specialty_id"`
	HospitalID      int       `json:"hospital_id"`
	ExperienceYears int       `json:"experience_years"`
	ConsultationFee int       `json:"consultation_fee"`
	Bio             string    `json:"bio"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Specialty *Specialty `json:"specialty,omitempty" gorm:"foreignKey:SpecialtyID"`
	Hospital  *Hospital  `json:"hospital,omitempty" gorm:"foreignKey:HospitalID"`
}
