package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/repository"
)

type DeleteIrrigationScheduleUsecaseI interface {
	Execute(ctx context.Context, id string) error
}

type DeleteIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

func NewDeleteIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) DeleteIrrigationScheduleUsecaseI {
	return &DeleteIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

func (u *DeleteIrrigationScheduleUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	// Check if irrigation schedule exists
	existingSchedule, err := u.irrigationScheduleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingSchedule == nil {
		return ErrIrrigationScheduleNotFound
	}

	// Delete the schedule
	err = u.irrigationScheduleRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
