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

type DoctorRepository interface {
	CreateDoctor(profile *entity.DoctorProfile) error
	GetAll() ([]entity.DoctorProfile, error)
	GetByID(id int) (*entity.DoctorProfile, error)
	GetBySpecialtyID(specialtyID int) ([]entity.DoctorProfile, error)
	GetByHospitalID(hospitalID int) ([]entity.DoctorProfile, error)
	UpdateDoctor(profile *entity.DoctorProfile) error
}

type SymptomAnalysisRepository interface {
	Create(analysis *entity.SymptomAnalysis) error
	GetByID(id int) (*entity.SymptomAnalysis, error)
}

type DoctorService interface {
	CreateDoctor(input CreateDoctorInput) (*entity.DoctorProfile, error)
	UpdateDoctor(id int, input UpdateDoctorInput) (*entity.DoctorProfile, error)
	ActivateDoctor(id int) error
	DeactivateDoctor(id int) error
	GetDoctors(specialtyID int, hospitalID int) ([]entity.DoctorProfile, error)
	GetDoctorByID(id int) (*entity.DoctorProfile, error)
}

type ScheduleRepository interface {
	GetDoctorsBySpecialtyAndTime(
		specialtyID int,
		startTime string,
		endTime string,
	) ([]entity.DoctorProfile, error)
}

type ScheduleService interface{}

type AppointmentService interface{}

type PaymentService interface{}

type AIService interface {
	AnalyzeSymptoms(patientID int, input SymptomAnalysisInput) (*SymptomAnalysisResponse, error)
}
