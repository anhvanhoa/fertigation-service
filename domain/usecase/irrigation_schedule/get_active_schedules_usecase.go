package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetActiveSchedulesUsecaseI interface {
	Execute(ctx context.Context, request common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error)
}

type GetActiveSchedulesUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
	helper                 utils.Helper
}

func NewGetActiveSchedulesUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository, helper utils.Helper) GetActiveSchedulesUsecaseI {
	return &GetActiveSchedulesUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
		helper:                 helper,
	}
}

func (u *GetActiveSchedulesUsecase) Execute(ctx context.Context, request common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error) {
	schedules, total, err := u.irrigationScheduleRepo.GetActiveSchedules(ctx, true, request)
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
