package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type specialtyRepository struct {
	db *gorm.DB
}

func NewSpecialtyRepository(db *gorm.DB) *specialtyRepository {
	return &specialtyRepository{
		db: db,
	}
}

func (r *specialtyRepository) GetAll() ([]entity.Specialty, error) {
	var specialties []entity.Specialty

	err := r.db.
		Find(&specialties).
		Error

	return specialties, err
}

func (r *specialtyRepository) GetByID(id int) (*entity.Specialty, error) {
	var specialty entity.Specialty

	err := r.db.
		First(&specialty, id).
		Error

	return &specialty, err
}
