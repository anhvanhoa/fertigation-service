package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"
)

// GetSchedulesForExecutionUsecase handles retrieving schedules that need to be executed
type GetSchedulesForExecutionUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewGetSchedulesForExecutionUsecase creates a new instance of GetSchedulesForExecutionUsecase
func NewGetSchedulesForExecutionUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *GetSchedulesForExecutionUsecase {
	return &GetSchedulesForExecutionUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute retrieves schedules that need to be executed within a time range
func (u *GetSchedulesForExecutionUsecase) Execute(ctx context.Context, from, to string) ([]*entity.IrrigationSchedule, error) {
	if from == "" || to == "" {
		return nil, ErrInvalidTimeRange
	}

	// Validate time format
	_, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, ErrInvalidTimeFormat
	}

	_, err = time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, ErrInvalidTimeFormat
	}

	schedules, err := u.irrigationScheduleRepo.GetSchedulesForExecution(ctx, from, to)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
