package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetExpiredFertilizersUsecase handles retrieving expired fertilizer types
type GetExpiredFertilizersUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewGetExpiredFertilizersUsecase creates a new instance of GetExpiredFertilizersUsecase
func NewGetExpiredFertilizersUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *GetExpiredFertilizersUsecase {
	return &GetExpiredFertilizersUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute retrieves all expired fertilizer types
func (u *GetExpiredFertilizersUsecase) Execute(ctx context.Context) ([]*entity.FertilizerType, error) {
	fertilizerTypes, err := u.fertilizerTypeRepo.GetExpiredFertilizers(ctx)
	if err != nil {
		return nil, err
	}

	return fertilizerTypes, nil
}
