package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetCompletedSchedulesUsecaseI interface {
	Execute(ctx context.Context, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error)
}

// GetCompletedSchedulesUsecase handles retrieving completed fertilizer schedules
type GetCompletedSchedulesUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
	helper                 utils.Helper
}

// NewGetCompletedSchedulesUsecase creates a new instance of GetCompletedSchedulesUsecase
func NewGetCompletedSchedulesUsecase(fertilizerScheduleRepo repository.FertilizerScheduleRepository, helper utils.Helper) GetCompletedSchedulesUsecaseI {
	return &GetCompletedSchedulesUsecase{
		fertilizerScheduleRepo,
		helper,
	}
}

// Execute retrieves all completed fertilizer schedules
func (u *GetCompletedSchedulesUsecase) Execute(ctx context.Context, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error) {
	schedules, total, err := u.fertilizerScheduleRepo.GetCompletedSchedules(ctx, request)
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
