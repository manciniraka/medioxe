package repository

import (
	"github.com/manciniraka/medioxe/internal/entity"
	"gorm.io/gorm"
)

type symptomAnalysisRepository struct {
	db *gorm.DB
}

func NewSymptomAnalysisRepository(db *gorm.DB) *symptomAnalysisRepository {
	return &symptomAnalysisRepository{
		db: db,
	}
}

func (r *symptomAnalysisRepository) CreateSymptomAnalysis(analysis *entity.SymptomAnalysis) error {
	return r.db.Create(analysis).Error
}

func (r *symptomAnalysisRepository) GetByID(id int) (*entity.SymptomAnalysis, error) {
	var analysis entity.SymptomAnalysis

	err := r.db.
		Preload("Patient").
		Preload("RecommendedSpecialty").
		First(&analysis, id).
		Error

	if err != nil {
		return nil, err
	}

	return &analysis, nil
}
