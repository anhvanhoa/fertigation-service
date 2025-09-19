package irrigation_log

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetIrrigationLogUsecase handles retrieving irrigation logs
type GetIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewGetIrrigationLogUsecase creates a new instance of GetIrrigationLogUsecase
func NewGetIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) *GetIrrigationLogUsecase {
	return &GetIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

// Execute retrieves an irrigation log by ID
func (u *GetIrrigationLogUsecase) Execute(ctx context.Context, id string) (*entity.IrrigationLog, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	irrigationLog, err := u.irrigationLogRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if irrigationLog == nil {
		return nil, ErrIrrigationLogNotFound
	}

	return irrigationLog, nil
}
