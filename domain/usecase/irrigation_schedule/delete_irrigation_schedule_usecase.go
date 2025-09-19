package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/repository"
)

// DeleteIrrigationScheduleUsecase handles deleting irrigation schedules
type DeleteIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewDeleteIrrigationScheduleUsecase creates a new instance of DeleteIrrigationScheduleUsecase
func NewDeleteIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *DeleteIrrigationScheduleUsecase {
	return &DeleteIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute deletes an irrigation schedule by ID
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
