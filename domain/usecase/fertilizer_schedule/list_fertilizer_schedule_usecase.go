package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListFertilizerScheduleUsecaseI interface {
	Execute(ctx context.Context, filter *entity.FertilizerScheduleFilter) (*common.PaginationResult[*entity.FertilizerSchedule], error)
}

// ListFertilizerScheduleUsecase handles listing fertilizer schedules
type ListFertilizerScheduleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
	helper                 utils.Helper
}

// NewListFertilizerScheduleUsecase creates a new instance of ListFertilizerScheduleUsecase
func NewListFertilizerScheduleUsecase(
	fertilizerScheduleRepo repository.FertilizerScheduleRepository,
	helper utils.Helper,
) *ListFertilizerScheduleUsecase {
	return &ListFertilizerScheduleUsecase{
		fertilizerScheduleRepo: fertilizerScheduleRepo,
		helper:                 helper,
	}
}

// Execute retrieves a list of fertilizer schedules with filtering and pagination
func (u *ListFertilizerScheduleUsecase) Execute(ctx context.Context, filter *entity.FertilizerScheduleFilter) (*common.PaginationResult[*entity.FertilizerSchedule], error) {
	u.validateRequest(filter)
	response, total, err := u.fertilizerScheduleRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	totalPages := u.helper.CalculateTotalPages(total, int64(filter.Limit))
	return &common.PaginationResult[*entity.FertilizerSchedule]{
		Data:       response,
		Total:      total,
		Page:       filter.Page,
		PageSize:   filter.Limit,
		TotalPages: totalPages,
	}, nil
}

func (u *ListFertilizerScheduleUsecase) validateRequest(filter *entity.FertilizerScheduleFilter) {
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

	// Set default sort values
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}
	if filter.SortOrder == "" {
		filter.SortOrder = "desc"
	}
}
