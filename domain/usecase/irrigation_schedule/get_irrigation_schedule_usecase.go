package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetIrrigationScheduleUsecase handles retrieving irrigation schedules
type GetIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewGetIrrigationScheduleUsecase creates a new instance of GetIrrigationScheduleUsecase
func NewGetIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *GetIrrigationScheduleUsecase {
	return &GetIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute retrieves an irrigation schedule by ID
func (u *GetIrrigationScheduleUsecase) Execute(ctx context.Context, id string) (*entity.IrrigationSchedule, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	irrigationSchedule, err := u.irrigationScheduleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if irrigationSchedule == nil {
		return nil, ErrIrrigationScheduleNotFound
	}

	return irrigationSchedule, nil
}
