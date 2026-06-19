package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type appointmentHistoryRepository struct {
	db *gorm.DB
}

func NewAppointmentHistoryRepository(
	db *gorm.DB,
) *appointmentHistoryRepository {
	return &appointmentHistoryRepository{
		db: db,
	}
}

func (r *appointmentHistoryRepository) CreateAppointmentHistory(history *entity.AppointmentHistory) error {
	return r.db.Create(history).Error
}

func (r *appointmentHistoryRepository) GetByAppointmentID(appointmentID int) ([]entity.AppointmentHistory, error) {
	var histories []entity.AppointmentHistory

	err := r.db.
		Where("appointment_id = ?", appointmentID).
		Order("created_at ASC").
		Find(&histories).
		Error

	return histories, err
}

func (r *appointmentHistoryRepository) GetAll() ([]entity.AppointmentHistory, error) {
	var histories []entity.AppointmentHistory

	err := r.db.
		Preload("Appointment").
		Order("created_at DESC").
		Find(&histories).
		Error

	return histories, err
}
