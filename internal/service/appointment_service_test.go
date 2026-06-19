package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/manciniraka/medioxe/internal/entity"
	service "github.com/manciniraka/medioxe/internal/service"
	mocks "github.com/manciniraka/medioxe/mocks"
)

func TestCreateAppointmentSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppointmentRepo := mocks.NewMockAppointmentRepository(ctrl)
	mockAppointmentHistoryRepo := mocks.NewMockAppointmentHistoryRepository(ctrl)
	mockScheduleRepo := mocks.NewMockScheduleRepository(ctrl)
	mockDoctorRepo := mocks.NewMockDoctorRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockNotificationRepo := mocks.NewMockNotificationRepository(ctrl)

	svc := service.NewAppointmentService(
		mockAppointmentRepo,
		mockAppointmentHistoryRepo,
		mockScheduleRepo,
		mockDoctorRepo,
		mockUserRepo,
		mockNotificationRepo,
	)

	schedule := &entity.Schedule{
		ID:       1,
		DoctorID: 1,
		IsBooked: false,
	}

	patient := &entity.User{
		ID:       1,
		FullName: "Raka",
		Email:    "raka@test.com",
	}

	input := service.CreateAppointmentInput{
		ScheduleID: 1,
		Notes:      "Sakit kepala",
	}

	mockScheduleRepo.
		EXPECT().
		GetByID(1).
		Return(schedule, nil)

	mockAppointmentRepo.
		EXPECT().
		CreateAppointment(gomock.Any()).
		Return(nil)

	mockAppointmentHistoryRepo.
		EXPECT().
		CreateAppointmentHistory(gomock.Any()).
		Return(nil)

	mockScheduleRepo.
		EXPECT().
		UpdateSchedule(gomock.Any()).
		Return(nil)

	mockUserRepo.
		EXPECT().
		GetByID(1).
		Return(patient, nil)

	mockNotificationRepo.
		EXPECT().
		SendEmail(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).
		Return(nil)

	result, err := svc.CreateAppointment(
		1,
		input,
	)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(
		t,
		service.AppointmentPending,
		result.Status,
	)
}
