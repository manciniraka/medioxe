package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *doctorRepository {
	return &doctorRepository{
		db: db,
	}
}

func (r *doctorRepository) CreateDoctor(profile *entity.DoctorProfile) error {
	return r.db.Create(profile).Error
}

func (r *doctorRepository) GetAll() ([]entity.DoctorProfile, error) {
	var doctors []entity.DoctorProfile

	err := r.db.
		Preload("User").
		Preload("Specialty").
		Preload("Hospital").
		Where("is_active = ?", true).
		Find(&doctors).
		Error

	return doctors, err
}

func (r *doctorRepository) GetByID(id int) (*entity.DoctorProfile, error) {
	var doctor entity.DoctorProfile

	err := r.db.
		Preload("User").
		Preload("Specialty").
		Preload("Hospital").
		Find(&doctor, id).
		Error

	return &doctor, err
}

func (r *doctorRepository) GetByUserID(userID int) (*entity.DoctorProfile, error) {
	var doctor entity.DoctorProfile

	err := r.db.
		Where("user_id = ?", userID).
		First(&doctor).
		Error

	return &doctor, err
}

func (r *doctorRepository) GetBySpecialtyID(specialtyID int) ([]entity.DoctorProfile, error) {
	var doctors []entity.DoctorProfile

	err := r.db.
		Where("specialty_id = ?", specialtyID).
		Where("is_active = ?", true).
		Find(&doctors).
		Error

	return doctors, err
}

func (r *doctorRepository) GetByHospitalID(hospitalID int) ([]entity.DoctorProfile, error) {
	var doctors []entity.DoctorProfile

	err := r.db.
		Preload("User").
		Preload("Hospital").
		Preload("Hospital").
		Where("hospital_id = ?", hospitalID).
		Where("is_active = ?", true).
		Find(&doctors).
		Error

	return doctors, err
}

func (r *doctorRepository) UpdateDoctor(profile *entity.DoctorProfile) error {
	return r.db.Save(profile).Error
}
