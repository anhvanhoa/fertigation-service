package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetSchedulesByPlantingCycleUsecaseI interface {
	Execute(ctx context.Context, plantingCycleID string, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error)
}

// GetSchedulesByPlantingCycleUsecase handles retrieving fertilizer schedules by planting cycle
type GetSchedulesByPlantingCycleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
	helper                 utils.Helper
}

// NewGetSchedulesByPlantingCycleUsecase creates a new instance of GetSchedulesByPlantingCycleUsecase
func NewGetSchedulesByPlantingCycleUsecase(
	fertilizerScheduleRepo repository.FertilizerScheduleRepository,
	helper utils.Helper,
) GetSchedulesByPlantingCycleUsecaseI {
	return &GetSchedulesByPlantingCycleUsecase{
		fertilizerScheduleRepo,
		helper,
	}
}

// Execute retrieves fertilizer schedules by planting cycle ID
func (u *GetSchedulesByPlantingCycleUsecase) Execute(ctx context.Context, plantingCycleID string, request common.Pagination) (*common.PaginationResult[*entity.FertilizerSchedule], error) {
	if plantingCycleID == "" {
		return nil, ErrInvalidPlantingCycleID
	}

	if schedules, total, err := u.fertilizerScheduleRepo.GetByPlantingCycleID(ctx, plantingCycleID, request); schedules == nil {
		return nil, ErrFertilizerScheduleNotFound
	} else if err != nil {
		return nil, err
	} else {
		totalPages := u.helper.CalculateTotalPages(total, int64(request.PageSize))
		return &common.PaginationResult[*entity.FertilizerSchedule]{
			Total:      total,
			Page:       request.Page,
			PageSize:   request.PageSize,
			Data:       schedules,
			TotalPages: totalPages,
		}, nil
	}
}
