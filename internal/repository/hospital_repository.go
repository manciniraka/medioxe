package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) *hospitalRepository {
	return &hospitalRepository{
		db: db,
	}
}

func (r *hospitalRepository) GetAll() ([]entity.Hospital, error) {
	var hospitals []entity.Hospital

	err := r.db.
		Find(&hospitals).
		Error

	return hospitals, err
}

func (r *hospitalRepository) GetByID(id int) (*entity.Hospital, error) {
	var hospital entity.Hospital

	err := r.db.
		First(&hospital, id).
		Error

	return &hospital, err
}
