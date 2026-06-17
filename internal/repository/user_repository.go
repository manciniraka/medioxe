package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *userRepository) GetByID(id int) (*entity.User, error) {
	var user entity.User

	err := r.db.
		First(&user, id).
		Error

	return &user, err
}
