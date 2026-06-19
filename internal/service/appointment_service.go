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
	userRepo        UserRepository

	notificationRepo NotificationRepository
}

func NewAppointmentService(
	appointmentRepo AppointmentRepository,
	scheduleRepo ScheduleRepository,
	doctorRepo DoctorRepository,
	userRepo UserRepository,
	notificationRepo NotificationRepository,
) AppointmentService {
	return &appointmentService{
		appointmentRepo:  appointmentRepo,
		scheduleRepo:     scheduleRepo,
		doctorRepo:       doctorRepo,
		userRepo:         userRepo,
		notificationRepo: notificationRepo,
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

	patient, userErr := s.userRepo.GetByID(
		patientID,
	)

	if userErr == nil {
		_ = s.notificationRepo.SendEmail(
			patient.FullName,
			patient.Email,
			"Appointment Created",
			"Your appointment has been successfully created.",
		)
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

	if appointment.Status != AppointmentPending {
		return errors.New(
			"appointment already processed",
		)
	}

	appointment.Status = AppointmentConfirmed

	err = s.appointmentRepo.UpdateAppointment(appointment)
	if err != nil {
		return err
	}

	patient, userErr := s.userRepo.GetByID(
		appointment.PatientID,
	)

	if userErr == nil {
		_ = s.notificationRepo.SendEmail(
			patient.FullName,
			patient.Email,
			"Appointment Confirmed",
			"Your appointment has been confirmed.",
		)
	}

	return nil
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

	if appointment.Status != AppointmentConfirmed {
		return errors.New(
			"appointment must be confirmed first",
		)
	}

	appointment.Status = AppointmentCompleted

	err = s.appointmentRepo.UpdateAppointment(appointment)
	if err != nil {
		return err
	}

	patient, userErr := s.userRepo.GetByID(
		appointment.PatientID,
	)

	if userErr == nil {
		_ = s.notificationRepo.SendEmail(
			patient.FullName,
			patient.Email,
			"Appointment Completed",
			"Your appointment has been completed.",
		)
	}

	return nil
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

	if appointment.Status == AppointmentCompleted {
		return errors.New(
			"completed appointment cannot be cancelled",
		)
	}

	appointment.Status = AppointmentCancelled

	err = s.appointmentRepo.UpdateAppointment(appointment)
	if err != nil {
		return err
	}

	patient, userErr := s.userRepo.GetByID(
		appointment.PatientID,
	)

	if userErr == nil {
		_ = s.notificationRepo.SendEmail(
			patient.FullName,
			patient.Email,
			"Appointment Cancelled",
			"Your appointment has been cancelled.",
		)
	}

	return nil
}
