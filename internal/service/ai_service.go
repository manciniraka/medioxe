package service

import (
	"strings"

	geminiai "github.com/manciniraka/medioxe/external/geminiAI"
	"github.com/manciniraka/medioxe/internal/entity"
)

type SymptomAnalysisInput struct {
	Symptoms string `json:"symptoms" validate:"required"`
}

type SymptomAnalysisResponse struct {
	AnalysisID           int                    `json:"analysis_id"`
	RecommendedSpecialty entity.Specialty       `json:"recommended_specialty"`
	RecommendedDoctor    []entity.DoctorProfile `json:"recommended_doctors"`
	AISummary            string                 `json:"ai_summary"`
}

type aiService struct {
	geminiClient        *geminiai.GeminiClient
	specialtyRepo       SpecialtyRepository
	doctorRepo          DoctorRepository
	symptomAnalysisRepo SymptomAnalysisRepository
	scheduleRepo        ScheduleRepository
}

func NewAIService(
	geminiClient *geminiai.GeminiClient,
	specialtyRepo SpecialtyRepository,
	doctorRepo DoctorRepository,
	symptomAnalysisRepo SymptomAnalysisRepository,
	scheduleRepo ScheduleRepository,
) AIService {
	return &aiService{
		geminiClient:        geminiClient,
		specialtyRepo:       specialtyRepo,
		doctorRepo:          doctorRepo,
		symptomAnalysisRepo: symptomAnalysisRepo,
		scheduleRepo:        scheduleRepo,
	}
}

func (s *aiService) AnalyzeSymptoms(
	patientID int,
	input SymptomAnalysisInput,
) (*SymptomAnalysisResponse, error) {
	geminiResult, err := s.geminiClient.AnalyzeSymptoms(input.Symptoms)
	if err != nil {
		return nil, err
	}

	specialty, err := s.specialtyRepo.GetByName(geminiResult.RecommendedSpecialty)
	if err != nil {
		return nil, err
	}

	var doctors []entity.DoctorProfile

	hasTimePreference := geminiResult.PreferredTimeStart != nil &&
		geminiResult.PreferredTimeEnd != nil &&
		strings.TrimSpace(*geminiResult.PreferredTimeStart) != "" &&
		strings.TrimSpace(*geminiResult.PreferredTimeEnd) != ""

	if hasTimePreference {
		doctors, err = s.scheduleRepo.GetDoctorsBySpecialtyAndTime(
			specialty.ID,
			*geminiResult.PreferredTimeStart,
			*geminiResult.PreferredTimeEnd,
		)

		if err != nil {
			return nil, err
		}
	} else {
		doctors, err = s.doctorRepo.GetBySpecialtyID(specialty.ID)
		if err != nil {
			return nil, err
		}
	}

	analysis := entity.SymptomAnalysis{
		PatientID:              patientID,
		Symptoms:               input.Symptoms,
		RecommendedSpecialtyID: specialty.ID,
		AISummary:              geminiResult.Summary,
	}

	err = s.symptomAnalysisRepo.CreateSymptomAnalysis(&analysis)

	if err != nil {
		return nil, err
	}

	return &SymptomAnalysisResponse{
		AnalysisID:           analysis.ID,
		RecommendedSpecialty: *specialty,
		RecommendedDoctor:    doctors,
		AISummary:            geminiResult.Summary,
	}, nil
}
