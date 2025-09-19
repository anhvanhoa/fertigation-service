package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetExpiringSoonUsecase handles retrieving fertilizer types expiring soon
type GetExpiringSoonUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewGetExpiringSoonUsecase creates a new instance of GetExpiringSoonUsecase
func NewGetExpiringSoonUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *GetExpiringSoonUsecase {
	return &GetExpiringSoonUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute retrieves fertilizer types expiring within specified days
func (u *GetExpiringSoonUsecase) Execute(ctx context.Context, days int) ([]*entity.FertilizerType, error) {
	if days <= 0 {
		days = 30 // Default to 30 days
	}

	fertilizerTypes, err := u.fertilizerTypeRepo.GetExpiringSoon(ctx, days)
	if err != nil {
		return nil, err
	}

	return fertilizerTypes, nil
}
