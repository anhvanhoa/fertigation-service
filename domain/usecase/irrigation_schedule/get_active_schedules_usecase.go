package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetActiveSchedulesUsecase handles retrieving active irrigation schedules
type GetActiveSchedulesUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewGetActiveSchedulesUsecase creates a new instance of GetActiveSchedulesUsecase
func NewGetActiveSchedulesUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *GetActiveSchedulesUsecase {
	return &GetActiveSchedulesUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute retrieves all active irrigation schedules
func (u *GetActiveSchedulesUsecase) Execute(ctx context.Context) ([]*entity.IrrigationSchedule, error) {
	schedules, err := u.irrigationScheduleRepo.GetActiveSchedules(ctx)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
