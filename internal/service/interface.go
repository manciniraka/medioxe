package service

import "github.com/manciniraka/medioxe/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetByID(id int) (*entity.User, error)
}

type AuthService interface {
	Register(input RegisterInput) (*entity.User, error)
	Login(input LoginInput) (string, error)
}

type UserService interface {
	GetProfile(userID int) (*entity.User, error)
}

type SpecialtyRepository interface {
	GetAll() ([]entity.Specialty, error)
	GetByID(id int) (*entity.Specialty, error)
	GetByName(name string) (*entity.Specialty, error)
}

type HospitalRepository interface {
	GetAll() ([]entity.Hospital, error)
	GetByID(id int) (*entity.Hospital, error)
}

type DoctorRepository interface {
	CreateDoctor(profile *entity.DoctorProfile) error
	GetAll() ([]entity.DoctorProfile, error)
	GetByID(id int) (*entity.DoctorProfile, error)
	GetByUserID(userID int) (*entity.DoctorProfile, error)
	GetBySpecialtyID(specialtyID int) ([]entity.DoctorProfile, error)
	GetByHospitalID(hospitalID int) ([]entity.DoctorProfile, error)
	UpdateDoctor(profile *entity.DoctorProfile) error
}

type SymptomAnalysisRepository interface {
	CreateSymptomAnalysis(analysis *entity.SymptomAnalysis) error
	GetByID(id int) (*entity.SymptomAnalysis, error)
}

type DoctorService interface {
	CreateDoctor(input CreateDoctorInput) (*entity.DoctorProfile, error)
	UpdateDoctor(id int, input UpdateDoctorInput) (*entity.DoctorProfile, error)
	ActivateDoctor(id int) error
	DeactivateDoctor(id int) error
	GetDoctors(specialtyID int, hospitalID int) ([]entity.DoctorProfile, error)
	GetDoctorByID(id int) (*entity.DoctorProfile, error)
	GetMyProfile(userID int) (*entity.DoctorProfile, error)
	UpdateMyProfile(userID int, input UpdateMyProfileInput) (*entity.DoctorProfile, error)
}

type ScheduleRepository interface {
	CreateSchedule(schedule *entity.Schedule) error
	GetByID(id int) (*entity.Schedule, error)
	UpdateSchedule(schedule *entity.Schedule) error
	DeleteSchedule(id int) error
	GetByDoctorID(doctorID int) ([]entity.Schedule, error)
	GetDoctorsBySpecialtyAndTime(
		specialtyID int,
		startTime string,
		endTime string,
	) ([]entity.DoctorProfile, error)
}

type ScheduleService interface {
	CreateSchedule(userID int, input CreateScheduleInput) (*entity.Schedule, error)
	UpdateSchedule(
		userID int,
		scheduleID int,
		input UpdateScheduleInput,
	) (*entity.Schedule, error)
	DeleteSchedule(userID int, scheduleID int) error
	GetMySchedules(userID int) ([]entity.Schedule, error)
	GetDoctorSchedules(doctorID int) ([]entity.Schedule, error)
}

type AppointmentRepository interface {
	CreateAppointment(appointment *entity.Appointment) error
	GetByID(id int) (*entity.Appointment, error)
	GetByPatientID(patientID int) ([]entity.Appointment, error)
	GetByDoctorID(doctorID int) ([]entity.Appointment, error)
	UpdateAppointment(appointment *entity.Appointment) error
}

type AppointmentHistoryRepository interface {
	CreateAppointmentHistory(history *entity.AppointmentHistory) error
	GetByAppointmentID(appointmentID int) ([]entity.AppointmentHistory, error)
	GetAll() ([]entity.AppointmentHistory, error)
}

type AppointmentService interface {
	CreateAppointment(patientID int, input CreateAppointmentInput) (*entity.Appointment, error)
	GetMyAppointments(patientID int) ([]entity.Appointment, error)
	GetDoctorAppointments(userID int) ([]entity.Appointment, error)
	ConfirmAppointment(userID int, appointmentID int) error
	CompleteAppointment(userID int, appointmentID int) error
	CancelAppointment(userID int, appointmentID int) error
	GetAppointmentHistory(appointmentID int) ([]entity.AppointmentHistory, error)
	GetAll() ([]entity.AppointmentHistory, error)
}

type NotificationRepository interface {
	SendEmail(
		toName string,
		toEmail string,
		subject string,
		message string,
	) error
}

type AIService interface {
	AnalyzeSymptoms(patientID int, input SymptomAnalysisInput) (*SymptomAnalysisResponse, error)
}
