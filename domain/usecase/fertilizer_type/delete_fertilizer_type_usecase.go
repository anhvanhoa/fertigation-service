package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/repository"
)

// DeleteFertilizerTypeUsecase handles deleting fertilizer types
type DeleteFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewDeleteFertilizerTypeUsecase creates a new instance of DeleteFertilizerTypeUsecase
func NewDeleteFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *DeleteFertilizerTypeUsecase {
	return &DeleteFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute deletes a fertilizer type by ID
func (u *DeleteFertilizerTypeUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	// Check if fertilizer type exists
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
