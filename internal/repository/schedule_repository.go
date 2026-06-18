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

func (r *scheduleRepository) GetByDoctorID(doctorID int) ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	err := r.db.
		Where("doctor_id = ?", doctorID).
		Order("date ASC").
		Order("start_time ASC").
		Find(&schedules).
		Error

	return schedules, err
}

func (r *scheduleRepository) GetDoctorsBySpecialtyAndTime(
	specialtyID int,
	startTime string,
	endTime string,
) ([]entity.DoctorProfile, error) {
	var doctors []entity.DoctorProfile
	err := r.db.
		Model(&entity.DoctorProfile{}).
		Distinct().
		Joins(
			"JOIN schedules ON schedules.doctor_id = doctor_profiles.id",
		).
		Where(
			"doctor_profiles.specialty_id = ?",
			specialtyID,
		).
		Where(
			"doctor_profiles.is_active = ?",
			true,
		).
		Where("schedules.is_booked = ?",
			false,
		).
		Where(
			"schedules.start_time >= ?",
			startTime,
		).
		Where(
			"schedules.end_time <= ?",
			endTime,
		).
		Find(&doctors).
		Error

	if err != nil {
		return nil, err
	}

	return doctors, nil
}
