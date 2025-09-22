package irrigation_log

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// ListIrrigationLogUsecase handles listing irrigation logs
type ListIrrigationLogUsecaseI interface {
	Execute(ctx context.Context, filter *entity.IrrigationLogFilter) (*entity.ListIrrigationLogsResponse, error)
}

type ListIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewListIrrigationLogUsecase creates a new instance of ListIrrigationLogUsecase
func NewListIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) ListIrrigationLogUsecaseI {
	return &ListIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

func (u *ListIrrigationLogUsecase) Execute(ctx context.Context, filter *entity.IrrigationLogFilter) (*entity.ListIrrigationLogsResponse, error) {
	u.validateRequest(filter)
	response, err := u.irrigationLogRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	if response.Total > 0 {
		response.TotalPages = (response.Total + filter.Limit - 1) / filter.Limit
	}

	return response, nil
}

func (u *ListIrrigationLogUsecase) validateRequest(filter *entity.IrrigationLogFilter) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}
	if filter.SortOrder == "" {
		filter.SortOrder = "desc"
	}
}
