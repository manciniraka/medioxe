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

type DoctorService interface{}

type ScheduleService interface{}

type AppointmentService interface{}

type PaymentService interface{}

type AIService interface{}
