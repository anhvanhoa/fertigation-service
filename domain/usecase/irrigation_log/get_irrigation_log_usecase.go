package irrigation_log

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type GetIrrigationLogUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.IrrigationLog, error)
}

type GetIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

func NewGetIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) GetIrrigationLogUsecaseI {
	return &GetIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

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
