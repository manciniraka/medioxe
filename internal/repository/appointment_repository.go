package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(
	db *gorm.DB,
) *appointmentRepository {
	return &appointmentRepository{
		db: db,
	}
}

func (r *appointmentRepository) CreateAppointment(appointment *entity.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetByID(id int) (*entity.Appointment, error) {
	var appointment entity.Appointment

	err := r.db.
		First(&appointment, id).
		Error

	return &appointment, err
}

func (r *appointmentRepository) GetByPatientID(patientID int) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := r.db.
		Where("patient_id = ?", patientID).
		Find(&appointments).
		Error

	return appointments, err
}

func (r *appointmentRepository) GetByDoctorID(doctorID int) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	err := r.db.
		Where("doctor_id = ?", doctorID).
		Find(&appointments).
		Error

	return appointments, err
}

func (r *appointmentRepository) UpdateAppointment(appointment *entity.Appointment) error {
	return r.db.Save(appointment).Error
}
