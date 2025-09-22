package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListIrrigationScheduleUsecaseI interface {
	Execute(ctx context.Context, filter *entity.IrrigationScheduleFilter) (common.PaginationResult[*entity.IrrigationSchedule], error)
}

type ListIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
	helper                 utils.Helper
}

func NewListIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository, helper utils.Helper) ListIrrigationScheduleUsecaseI {
	return &ListIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
		helper:                 helper,
	}
}

func (u *ListIrrigationScheduleUsecase) Execute(ctx context.Context, filter *entity.IrrigationScheduleFilter) (common.PaginationResult[*entity.IrrigationSchedule], error) {
	// Set default pagination values
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

	response, total, err := u.irrigationScheduleRepo.List(ctx, filter)
	if err != nil {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, err
	}

	return common.PaginationResult[*entity.IrrigationSchedule]{
		Data:       response,
		Total:      total,
		Page:       filter.Page,
		PageSize:   filter.Limit,
		TotalPages: u.helper.CalculateTotalPages(total, int64(filter.Limit)),
	}, nil
}
