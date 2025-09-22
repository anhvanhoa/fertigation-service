package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type GetFertilizerTypeUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.FertilizerType, error)
}

type GetFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

func NewGetFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) GetFertilizerTypeUsecaseI {
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
