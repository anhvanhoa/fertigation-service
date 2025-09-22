package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type GetExpiringSoonUsecaseI interface {
	Execute(ctx context.Context, days int, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error)
}

type GetExpiringSoonUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
	helper             utils.Helper
}

// NewGetExpiringSoonUsecase creates a new instance of GetExpiringSoonUsecase
func NewGetExpiringSoonUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository, helper utils.Helper) GetExpiringSoonUsecaseI {
	return &GetExpiringSoonUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
		helper:             helper,
	}
}

// Execute retrieves fertilizer types expiring within specified days
func (u *GetExpiringSoonUsecase) Execute(ctx context.Context, days int, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error) {
	if days <= 0 {
		days = 30 // Default to 30 days
	}

	fertilizerTypes, total, err := u.fertilizerTypeRepo.GetExpiringSoon(ctx, days, filter)
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
