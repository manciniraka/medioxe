package entity

import "time"

type Appointment struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	PatientID         int       `json:"patient_id"`
	DoctorID          int       `json:"doctor_id"`
	ScheduleID        int       `json:"schedule_id"`
	SymptomAnalysisID *int      `json:"symptom_analysis_id"`
	Notes             string    `json:"notes"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	Patient         *User            `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Doctor          *DoctorProfile   `json:"doctor,omitempty" gorm:"foreignKey:DoctorID"`
	Schedule        *Schedule        `json:"schedule,omitempty" gorm:"foreignKey:ScheduleID"`
	SymptomAnalysis *SymptomAnalysis `json:"symptom_analysis" gorm:"foreignKey:SymptomAnalysisID"`
}
