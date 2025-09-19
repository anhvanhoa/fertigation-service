package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/repository"
)

type DeleteFertilizerScheduleUsecaseI interface {
	Execute(ctx context.Context, id string) error
}

// DeleteFertilizerScheduleUsecase handles deleting fertilizer schedules
type DeleteFertilizerScheduleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
}

// NewDeleteFertilizerScheduleUsecase creates a new instance of DeleteFertilizerScheduleUsecase
func NewDeleteFertilizerScheduleUsecase(fertilizerScheduleRepo repository.FertilizerScheduleRepository) DeleteFertilizerScheduleUsecaseI {
	return &DeleteFertilizerScheduleUsecase{
		fertilizerScheduleRepo: fertilizerScheduleRepo,
	}
}

// Execute deletes a fertilizer schedule by ID
func (u *DeleteFertilizerScheduleUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	// Check if fertilizer schedule exists
	existingSchedule, err := u.fertilizerScheduleRepo.GetByID(ctx, id)
	if existingSchedule == nil {
		return ErrFertilizerScheduleNotFound
	}
	if err != nil {
		return err
	}

	// Delete the schedule
	err = u.fertilizerScheduleRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
