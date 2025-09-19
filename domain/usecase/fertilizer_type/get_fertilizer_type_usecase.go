package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetFertilizerTypeUsecase handles retrieving fertilizer types
type GetFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewGetFertilizerTypeUsecase creates a new instance of GetFertilizerTypeUsecase
func NewGetFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *GetFertilizerTypeUsecase {
	return &GetFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute retrieves a fertilizer type by ID
func (u *GetFertilizerTypeUsecase) Execute(ctx context.Context, id string) (*entity.FertilizerType, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	fertilizerType, err := u.fertilizerTypeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if fertilizerType == nil {
		return nil, ErrFertilizerTypeNotFound
	}

	return fertilizerType, nil
}
