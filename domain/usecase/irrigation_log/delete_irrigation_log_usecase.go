package irrigation_log

import (
	"context"
	"fertigation-Service/domain/repository"
)

type DeleteIrrigationLogUsecaseI interface {
	Execute(ctx context.Context, id string) error
}

type DeleteIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewDeleteIrrigationLogUsecase creates a new instance of DeleteIrrigationLogUsecase
func NewDeleteIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) DeleteIrrigationLogUsecaseI {
	return &DeleteIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

func (u *DeleteIrrigationLogUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

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
