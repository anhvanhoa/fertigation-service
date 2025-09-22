package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetExpiredFertilizersUsecaseI interface {
	Execute(ctx context.Context, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error)
}

type GetExpiredFertilizersUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
	helper             utils.Helper
}

// NewGetExpiredFertilizersUsecase creates a new instance of GetExpiredFertilizersUsecase
func NewGetExpiredFertilizersUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository, helper utils.Helper) GetExpiredFertilizersUsecaseI {
	return &GetExpiredFertilizersUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
		helper:             helper,
	}
}

// Execute retrieves all expired fertilizer types
func (u *GetExpiredFertilizersUsecase) Execute(ctx context.Context, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error) {
	fertilizerTypes, total, err := u.fertilizerTypeRepo.GetExpiredFertilizers(ctx, filter)
	if err != nil {
		return common.PaginationResult[*entity.FertilizerType]{}, err
	}

	return common.PaginationResult[*entity.FertilizerType]{
		Data:       fertilizerTypes,
		Total:      total,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(filter.PageSize)),
	}, nil
}
