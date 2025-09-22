package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListFertilizerTypeUsecaseI interface {
	Execute(ctx context.Context, filter *entity.FertilizerTypeFilter) (common.PaginationResult[*entity.FertilizerType], error)
}

type ListFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
	helper             utils.Helper
}

// NewListFertilizerTypeUsecase creates a new instance of ListFertilizerTypeUsecase
func NewListFertilizerTypeUsecase(
	fertilizerTypeRepo repository.FertilizerTypeRepository,
	helper utils.Helper,
) ListFertilizerTypeUsecaseI {
	return &ListFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
		helper:             helper,
	}
}

func (u *ListFertilizerTypeUsecase) Execute(ctx context.Context, filter *entity.FertilizerTypeFilter) (common.PaginationResult[*entity.FertilizerType], error) {
	u.validateRequest(filter)

	response, total, err := u.fertilizerTypeRepo.List(ctx, filter)
	if err != nil {
		return common.PaginationResult[*entity.FertilizerType]{}, err
	}

	return common.PaginationResult[*entity.FertilizerType]{
		Data:       response,
		Total:      total,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(filter.PageSize)),
	}, nil
}

func (u *ListFertilizerTypeUsecase) validateRequest(filter *entity.FertilizerTypeFilter) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 {
		filter.PageSize = 10
	}
	if filter.PageSize > 100 {
		filter.PageSize = 100
	}
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}
	if filter.SortOrder == "" {
		filter.SortOrder = "desc"
	}
}
