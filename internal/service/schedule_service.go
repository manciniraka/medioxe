package service

import (
	"errors"
	"time"

	"github.com/manciniraka/medioxe/internal/entity"
)

type CreateScheduleInput struct {
	Date      time.Time `json:"date" validate:"required"`
	StartTime string    `json:"start_time" validate:"required"`
	EndTime   string    `json:"end_time" validate:"required"`
}

type UpdateScheduleInput struct {
	Date      time.Time `json:"date" validate:"required"`
	StartTime string    `json:"start_time" validate:"required"`
	EndTime   string    `json:"end_time" validate:"required"`
}

type scheduleService struct {
	scheduleRepo ScheduleRepository
	doctorRepo   DoctorRepository
}

func NewScheduleService(
	scheduleRepo ScheduleRepository,
	doctorRepo DoctorRepository,
) ScheduleService {
	return &scheduleService{
		scheduleRepo: scheduleRepo,
		doctorRepo:   doctorRepo,
	}
}

func (s *scheduleService) CreateSchedule(userID int, input CreateScheduleInput) (*entity.Schedule, error) {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	schedule := entity.Schedule{
		DoctorID:  doctor.ID,
		Date:      input.Date,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		IsBooked:  false,
	}

	err = s.scheduleRepo.CreateSchedule(&schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (s *scheduleService) GetMySchedules(userID int) ([]entity.Schedule, error) {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.scheduleRepo.GetByDoctorID(doctor.ID)
}

func (s *scheduleService) GetDoctorSchedules(doctorID int) ([]entity.Schedule, error) {
	return s.scheduleRepo.GetByDoctorID(doctorID)
}

func (s *scheduleService) UpdateSchedule(
	userID int,
	scheduleID int,
	input UpdateScheduleInput,
) (*entity.Schedule, error) {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	schedule, err := s.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return nil, err
	}

	if schedule.DoctorID != doctor.ID {
		return nil, errors.New(
			"you are not allowed to access this schedule",
		)
	}

	if !input.Date.IsZero() {
		schedule.Date = input.Date
	}

	if input.StartTime != "" {
		schedule.StartTime = input.StartTime
	}

	if input.EndTime != "" {
		schedule.EndTime = input.EndTime
	}

	err = s.scheduleRepo.UpdateSchedule(schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (s *scheduleService) DeleteSchedule(userID int, scheduleID int) error {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	schedule, err := s.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return err
	}

	if schedule.DoctorID != doctor.ID {
		return errors.New(
			"you are not allowed to access this schedule",
		)
	}

	return s.scheduleRepo.DeleteSchedule(scheduleID)
}
