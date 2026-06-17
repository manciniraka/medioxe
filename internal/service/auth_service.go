package service

import (
	"errors"
	"os"

	jwtgo "github.com/golang-jwt/jwt/v5"
	commonJWT "github.com/manciniraka/go-common/jwt"
	commonPassword "github.com/manciniraka/go-common/password"
	"github.com/manciniraka/medioxe/internal/entity"
)

type authService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

type RegisterInput struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Creates a new user
func (s *authService) Register(input RegisterInput) (*entity.User, error) {
	// Check existing Email
	existingUser, err := s.userRepo.GetByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New(
			"email already registered",
		)
	}

	// Hashing password using my personal module packages for repeatable usage
	hashedPassword, err := commonPassword.Hash(
		input.Password,
	)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "patient",
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

// Validates user credentials to login
func (s *authService) Login(input LoginInput) (string, error) {
	user, err := s.userRepo.GetByEmail(input.Email)
	if err != nil {
		return "", errors.New(
			"invalid email or password",
		)
	}

	// Comparing hashed password using my personal module packages
	err = commonPassword.Compare(
		user.Password,
		input.Password,
	)
	if err != nil {
		return "", errors.New(
			"invalid email or password",
		)
	}

	claims := jwtgo.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
	}

	// Generating JWT Token using my personal module packages
	token, err := commonJWT.GenerateToken(
		claims,
		os.Getenv("JWT_SECRET"),
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
