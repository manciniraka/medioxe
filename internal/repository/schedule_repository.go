package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *scheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r *scheduleRepository) CreateSchedule(schedule *entity.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) GetByID(id int) (*entity.Schedule, error) {
	var schedule entity.Schedule

	err := r.db.
		Preload("Doctor").
		First(&schedule, id).
		Error

	return &schedule, err
}

func (r *scheduleRepository) UpdateSchedule(schedule *entity.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) DeleteSchedule(id int) error {
	return r.db.Delete(&entity.Schedule{}, id).Error
}
