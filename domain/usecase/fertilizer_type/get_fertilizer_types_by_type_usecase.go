package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetFertilizerTypesByTypeUsecase handles retrieving fertilizer types by type
type GetFertilizerTypesByTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewGetFertilizerTypesByTypeUsecase creates a new instance of GetFertilizerTypesByTypeUsecase
func NewGetFertilizerTypesByTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *GetFertilizerTypesByTypeUsecase {
	return &GetFertilizerTypesByTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute retrieves fertilizer types by type
func (u *GetFertilizerTypesByTypeUsecase) Execute(ctx context.Context, fertilizerType string) ([]*entity.FertilizerType, error) {
	if fertilizerType == "" {
		return nil, ErrInvalidFertilizerType
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
		return nil, ErrInvalidFertilizerType
	}

	fertilizerTypes, err := u.fertilizerTypeRepo.GetByType(ctx, fertilizerType)
	if err != nil {
		return nil, err
	}

	return fertilizerTypes, nil
}
