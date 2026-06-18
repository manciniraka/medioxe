package entity

import "time"

type SymptomAnalysis struct {
	ID                     int       `json:"id" gorm:"primaryKey"`
	PatientID              int       `json:"patient_id"`
	Symptoms               string    `json:"symptoms"`
	RecommendedSpecialtyID int       `json:"recommended_specialty_id"`
	RecommendedDoctorID    *int      `json:"recommended_doctor_id"`
	AISummary              string    `json:"ai_summary"`
	CreatedAt              time.Time `json:"created_at"`

	Patient              User           `json:"patient" gorm:"foreignKey:PatientID"`
	RecommendedSpecialty Specialty      `json:"recommended_specialty" gorm:"foreignKey:RecommendedSpecialtyID"`
	RecommendedDoctor    *DoctorProfile `json:"recommended_doctor" gorm:"foreignKey:RecommendedDoctorID"`
}
