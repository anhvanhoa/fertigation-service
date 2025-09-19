package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetUpcomingSchedulesUsecaseI interface {
	Execute(ctx context.Context, days int, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error)
}

// GetUpcomingSchedulesUsecase handles retrieving upcoming fertilizer schedules
type GetUpcomingSchedulesUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
	helper                 utils.Helper
}

// NewGetUpcomingSchedulesUsecase creates a new instance of GetUpcomingSchedulesUsecase
func NewGetUpcomingSchedulesUsecase(
	fertilizerScheduleRepo repository.FertilizerScheduleRepository,
	helper utils.Helper,
) GetUpcomingSchedulesUsecaseI {
	return &GetUpcomingSchedulesUsecase{
		fertilizerScheduleRepo,
		helper,
	}
}

// Execute retrieves schedules that need to be executed soon
func (u *GetUpcomingSchedulesUsecase) Execute(ctx context.Context, days int, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error) {
	if days <= 0 {
		days = 7 // Default to 7 days
	}

	schedules, total, err := u.fertilizerScheduleRepo.GetUpcomingSchedules(ctx, days, request)
	if err != nil {
		return nil, err
	}

	totalPages := u.helper.CalculateTotalPages(total, int64(request.PageSize))

	return &common.PaginationResult[*entity.FertilizerSchedule]{
		Total:      total,
		Page:       request.Page,
		PageSize:   request.PageSize,
		Data:       schedules,
		TotalPages: totalPages,
	}, nil
}
