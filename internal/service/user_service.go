package service

import "github.com/manciniraka/medioxe/internal/entity"

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetProfile(userID int) (*entity.User, error) {
	return s.userRepo.GetByID(userID)
}
