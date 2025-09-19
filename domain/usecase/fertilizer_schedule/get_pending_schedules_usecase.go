package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetPendingSchedulesUsecaseI interface {
	Execute(ctx context.Context, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error)
}

// GetPendingSchedulesUsecase handles retrieving pending fertilizer schedules
type GetPendingSchedulesUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
	helper                 utils.Helper
}

// NewGetPendingSchedulesUsecase creates a new instance of GetPendingSchedulesUsecase
func NewGetPendingSchedulesUsecase(
	fertilizerScheduleRepo repository.FertilizerScheduleRepository,
	helper utils.Helper,
) GetPendingSchedulesUsecaseI {
	return &GetPendingSchedulesUsecase{
		fertilizerScheduleRepo,
		helper,
	}
}

// Execute retrieves all pending fertilizer schedules
func (u *GetPendingSchedulesUsecase) Execute(ctx context.Context, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error) {
	schedules, total, err := u.fertilizerScheduleRepo.GetPendingSchedules(ctx, request)
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
