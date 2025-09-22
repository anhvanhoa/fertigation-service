package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type GetIrrigationScheduleUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.IrrigationSchedule, error)
}

type GetIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

func NewGetIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) GetIrrigationScheduleUsecaseI {
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
