package service

import (
	"errors"

	"github.com/manciniraka/medioxe/internal/entity"
)

const (
	AppointmentPending   = "pending"
	AppointmentConfirmed = "confirmed"
	AppointmentCompleted = "completed"
	AppointmentCancelled = "cancelled"
)

type CreateAppointmentInput struct {
	ScheduleID        int    `json:"schedule_id" validate:"required"`
	SymptomAnalysisID *int   `json:"symptom_analysis_id"`
	Notes             string `json:"notes"`
}

type appointmentService struct {
	appointmentRepo AppointmentRepository
	scheduleRepo    ScheduleRepository
	doctorRepo      DoctorRepository
}

func NewAppointmentService(
	appointmentRepo AppointmentRepository,
	scheduleRepo ScheduleRepository,
	doctorRepo DoctorRepository,
) AppointmentService {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
		scheduleRepo:    scheduleRepo,
		doctorRepo:      doctorRepo,
	}
}

func (s *appointmentService) CreateAppointment(patientID int, input CreateAppointmentInput) (*entity.Appointment, error) {
	schedule, err := s.scheduleRepo.GetByID(input.ScheduleID)
	if err != nil {
		return nil, err
	}

	if schedule.IsBooked {
		return nil,
			errors.New(
				"schedule already booked",
			)
	}

	appointment := entity.Appointment{
		PatientID:         patientID,
		DoctorID:          schedule.DoctorID,
		ScheduleID:        schedule.ID,
		SymptomAnalysisID: input.SymptomAnalysisID,
		Notes:             input.Notes,
		Status:            AppointmentPending,
	}

	err = s.appointmentRepo.CreateAppointment(&appointment)
	if err != nil {
		return nil, err
	}

	schedule.IsBooked = true

	err = s.scheduleRepo.UpdateSchedule(schedule)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (s *appointmentService) GetMyAppointments(patientID int) ([]entity.Appointment, error) {
	return s.appointmentRepo.GetByPatientID(patientID)
}

func (s *appointmentService) GetDoctorAppointments(userID int) ([]entity.Appointment, error) {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.appointmentRepo.GetByDoctorID(doctor.ID)
}

func (s *appointmentService) ConfirmAppointment(userID int, appointmentID int) error {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	appointment, err := s.appointmentRepo.GetByID(appointmentID)
	if err != nil {
		return err
	}

	if appointment.DoctorID != doctor.ID {
		return errors.New(
			"you are not allowed to access this appointment",
		)
	}

	appointment.Status = AppointmentConfirmed

	return s.appointmentRepo.UpdateAppointment(appointment)
}

func (s *appointmentService) CompleteAppointment(userID int, appointmentID int) error {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	appointment, err := s.appointmentRepo.GetByID(appointmentID)
	if err != nil {
		return err
	}

	if appointment.DoctorID != doctor.ID {
		return errors.New(
			"you are not allowed to access this appointment",
		)
	}

	appointment.Status = AppointmentCompleted

	return s.appointmentRepo.UpdateAppointment(appointment)
}

func (s *appointmentService) CancelAppointment(userID int, appointmentID int) error {
	doctor, err := s.doctorRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	appointment, err := s.appointmentRepo.GetByID(appointmentID)
	if err != nil {
		return err
	}

	if appointment.DoctorID != doctor.ID {
		return errors.New(
			"you are not allowed to access this appointment",
		)
	}

	appointment.Status = AppointmentCancelled

	return s.appointmentRepo.UpdateAppointment(appointment)
}
