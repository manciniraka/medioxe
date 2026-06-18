package service

import (
	"errors"

	"github.com/manciniraka/go-common/password"
	"github.com/manciniraka/medioxe/internal/entity"
)

type CreateDoctorInput struct {
	FullName        string `json:"full_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	SpecialtyID     int    `json:"specialty_id" validate:"required"`
	HospitalID      int    `json:"hospital_id" validate:"required"`
	ExperienceYears int    `json:"experience_years" validate:"required"`
	ConsultationFee int    `json:"consultation_fee" validate:"required"`
	Bio             string `json:"bio" validate:"required"`
}

type UpdateDoctorInput struct {
	SpecialtyID     int    `json:"specialty_id"`
	HospitalID      int    `json:"hospital_id"`
	ExperienceYears int    `json:"experience_years"`
	ConsultationFee int    `json:"consultation_fee"`
	Bio             string `json:"bio"`
}

type doctorService struct {
	doctorRepo DoctorRepository
	userRepo   UserRepository
}

func NewDoctorService(
	doctorRepo DoctorRepository,
	userRepo UserRepository,
) DoctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
		userRepo:   userRepo,
	}
}

func (s *doctorService) CreateDoctor(input CreateDoctorInput) (*entity.DoctorProfile, error) {
	existingUser, _ := s.userRepo.GetByEmail(input.Email)

	if existingUser != nil {
		return nil,
			errors.New("email already registered")
	}

	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "doctor",
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	doctor := entity.DoctorProfile{
		UserID:          user.ID,
		SpecialtyID:     input.SpecialtyID,
		HospitalID:      input.HospitalID,
		ExperienceYears: input.ExperienceYears,
		ConsultationFee: input.ConsultationFee,
		Bio:             input.Bio,
		IsActive:        true,
	}

	err = s.doctorRepo.CreateDoctor(&doctor)
	if err != nil {
		return nil, err
	}

	return s.doctorRepo.GetByID(
		doctor.ID,
	)
}

func (s *doctorService) UpdateDoctor(id int, input UpdateDoctorInput) (*entity.DoctorProfile, error) {
	doctor, err := s.doctorRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("doctor not found")
	}

	if input.SpecialtyID != 0 {
		doctor.SpecialtyID = input.SpecialtyID
	}

	if input.HospitalID != 0 {
		doctor.HospitalID = input.HospitalID
	}

	if input.ExperienceYears != 0 {
		doctor.ExperienceYears = input.ExperienceYears
	}

	if input.ConsultationFee != 0 {
		doctor.ConsultationFee = input.ConsultationFee
	}

	if input.Bio != "" {
		doctor.Bio = input.Bio
	}

	err = s.doctorRepo.UpdateDoctor(doctor)
	if err != nil {
		return nil, err
	}

	return s.doctorRepo.GetByID(id)
}

func (s *doctorService) ActivateDoctor(id int) error {
	doctor, err := s.doctorRepo.GetByID(id)
	if err != nil {
		return errors.New("doctor not found")
	}

	doctor.IsActive = true

	return s.doctorRepo.UpdateDoctor(doctor)
}

func (s *doctorService) DeactivateDoctor(id int) error {
	doctor, err := s.doctorRepo.GetByID(id)
	if err != nil {
		return errors.New("doctor not found")
	}

	doctor.IsActive = false

	return s.doctorRepo.UpdateDoctor(doctor)
}

func (s *doctorService) GetDoctors(specialtyID int, hospitalID int) ([]entity.DoctorProfile, error) {
	if specialtyID != 0 {
		return s.doctorRepo.GetBySpecialtyID(
			specialtyID,
		)
	}

	if hospitalID != 0 {
		return s.doctorRepo.GetByHospitalID(
			hospitalID,
		)
	}

	return s.doctorRepo.GetAll()
}

func (s *doctorService) GetDoctorByID(id int) (*entity.DoctorProfile, error) {
	return s.doctorRepo.GetByID(id)
}
