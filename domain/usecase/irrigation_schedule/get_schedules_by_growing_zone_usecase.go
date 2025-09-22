package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetSchedulesByGrowingZoneUsecaseI interface {
	Execute(ctx context.Context, growingZoneID string, filter *common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error)
}

type GetSchedulesByGrowingZoneUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
	helper                 utils.Helper
}

func NewGetSchedulesByGrowingZoneUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository, helper utils.Helper) GetSchedulesByGrowingZoneUsecaseI {
	return &GetSchedulesByGrowingZoneUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
		helper:                 helper,
	}
}

func (u *GetSchedulesByGrowingZoneUsecase) Execute(ctx context.Context, growingZoneID string, filter *common.Pagination) (common.PaginationResult[*entity.IrrigationSchedule], error) {
	if growingZoneID == "" {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, ErrInvalidGrowingZoneID
	}

	schedules, total, err := u.irrigationScheduleRepo.GetByGrowingZoneID(ctx, growingZoneID, filter)
	if err != nil {
		return common.PaginationResult[*entity.IrrigationSchedule]{}, err
	}

	return common.PaginationResult[*entity.IrrigationSchedule]{
		Data:       schedules,
		Total:      total,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(filter.PageSize)),
	}, nil
}
