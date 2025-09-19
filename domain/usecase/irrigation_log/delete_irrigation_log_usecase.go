package irrigation_log

import (
	"context"
	"fertigation-Service/domain/repository"
)

// DeleteIrrigationLogUsecase handles deleting irrigation logs
type DeleteIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewDeleteIrrigationLogUsecase creates a new instance of DeleteIrrigationLogUsecase
func NewDeleteIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) *DeleteIrrigationLogUsecase {
	return &DeleteIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

// Execute deletes an irrigation log by ID
func (u *DeleteIrrigationLogUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	// Check if irrigation log exists
	existingLog, err := u.irrigationLogRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingLog == nil {
		return ErrIrrigationLogNotFound
	}

	// Delete the log
	err = u.irrigationLogRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
