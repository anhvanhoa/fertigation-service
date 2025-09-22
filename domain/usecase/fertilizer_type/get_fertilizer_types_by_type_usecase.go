package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

// GetFertilizerTypesByTypeUsecase handles retrieving fertilizer types by type
type GetFertilizerTypesByTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
	helper             utils.Helper
}

type GetFertilizerTypesByTypeUsecaseI interface {
	Execute(ctx context.Context, fertilizerType string, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error)
}

func NewGetFertilizerTypesByTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository, helper utils.Helper) GetFertilizerTypesByTypeUsecaseI {
	return &GetFertilizerTypesByTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
		helper:             helper,
	}
}

func (u *GetFertilizerTypesByTypeUsecase) Execute(ctx context.Context, fertilizerType string, filter common.Pagination) (common.PaginationResult[*entity.FertilizerType], error) {
	if fertilizerType == "" {
		return common.PaginationResult[*entity.FertilizerType]{}, ErrInvalidFertilizerType
	}

	// Validate fertilizer type
	validTypes := []string{"organic", "chemical", "liquid", "granular", "powder"}
	typeValid := false
	for _, validType := range validTypes {
		if fertilizerType == validType {
			typeValid = true
			break
		}
	}
	if !typeValid {
		return common.PaginationResult[*entity.FertilizerType]{}, ErrInvalidFertilizerType
	}

	fertilizerTypes, total, err := u.fertilizerTypeRepo.GetByType(ctx, fertilizerType, filter)
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
