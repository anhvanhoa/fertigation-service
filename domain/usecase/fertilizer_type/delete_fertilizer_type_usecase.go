package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/repository"
)

type DeleteFertilizerTypeUsecaseI interface {
	Execute(ctx context.Context, id string) error
}

type DeleteFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

func NewDeleteFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) DeleteFertilizerTypeUsecaseI {
	return &DeleteFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

func (u *DeleteFertilizerTypeUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	existingFertilizerType, err := u.fertilizerTypeRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingFertilizerType == nil {
		return ErrFertilizerTypeNotFound
	}

	// Delete the fertilizer type
	err = u.fertilizerTypeRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
