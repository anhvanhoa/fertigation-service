package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetSchedulesForExecutionUsecaseI interface {
	Execute(ctx context.Context, from, to string, request common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error)
}

type GetSchedulesForExecutionUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
	helper                 utils.Helper
}

func NewGetSchedulesForExecutionUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository, helper utils.Helper) GetSchedulesForExecutionUsecaseI {
	return &GetSchedulesForExecutionUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
		helper:                 helper,
	}
}

func (u *GetSchedulesForExecutionUsecase) Execute(ctx context.Context, from, to string, request common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error) {
	if from == "" || to == "" {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, ErrInvalidTimeRange
	}

	// Validate time format
	_, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, ErrInvalidTimeFormat
	}

	_, err = time.Parse(time.RFC3339, to)
	if err != nil {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, ErrInvalidTimeFormat
	}

	schedules, total, err := u.irrigationScheduleRepo.GetSchedulesForExecution(ctx, from, to, request)
	if err != nil {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, err
	}

	return common.PaginationResult[*entity.IrrigationSchedule]{
		Total:      total,
		Page:       request.Page,
		PageSize:   request.PageSize,
		Data:       schedules,
		TotalPages: u.helper.CalculateTotalPages(total, int64(request.PageSize)),
	}, nil
}
