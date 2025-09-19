package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
)

// ListFertilizerTypeUsecase handles listing fertilizer types
type ListFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
	helper             utils.Helper
}

// NewListFertilizerTypeUsecase creates a new instance of ListFertilizerTypeUsecase
func NewListFertilizerTypeUsecase(
	fertilizerTypeRepo repository.FertilizerTypeRepository,
	helper utils.Helper,
) *ListFertilizerTypeUsecase {
	return &ListFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
		helper:             helper,
	}
}

// Execute retrieves a list of fertilizer types with filtering and pagination
func (u *ListFertilizerTypeUsecase) Execute(ctx context.Context, filter *entity.FertilizerTypeFilter) (*entity.ListFertilizerTypesResponse, error) {
	u.validateRequest(filter)

	response, err := u.fertilizerTypeRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	if response.Total > 0 {
		response.TotalPages = int(u.helper.CalculateTotalPages(int64(response.Total), int64(filter.PageSize)))
	}

	return response, nil
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
