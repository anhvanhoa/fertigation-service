package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type CreateFertilizerScheduleUsecaseI interface {
	Execute(ctx context.Context, req *entity.CreateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error)
}

// CreateFertilizerScheduleUsecase handles the creation of fertilizer schedules
type CreateFertilizerScheduleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
}

// NewCreateFertilizerScheduleUsecase creates a new instance of CreateFertilizerScheduleUsecase
func NewCreateFertilizerScheduleUsecase(fertilizerScheduleRepo repository.FertilizerScheduleRepository) CreateFertilizerScheduleUsecaseI {
	return &CreateFertilizerScheduleUsecase{
		fertilizerScheduleRepo: fertilizerScheduleRepo,
	}
}

// Execute creates a new fertilizer schedule
func (u *CreateFertilizerScheduleUsecase) Execute(ctx context.Context, req *entity.CreateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}
	// Save to repository
	createdSchedule, err := u.fertilizerScheduleRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return createdSchedule, nil
}

// validateRequest validates the create request
func (u *CreateFertilizerScheduleUsecase) validateRequest(req *entity.CreateFertilizerScheduleRequest) error {
	if req.PlantingCycleID == "" {
		return ErrInvalidPlantingCycleID
	}

	if req.FertilizerTypeID == "" {
		return ErrInvalidFertilizerTypeID
	}

	if req.CreatedBy == "" {
		return ErrInvalidCreatedBy
	}

	if req.ApplicationMethod != "" {
		validMethods := []string{"foliar", "soil", "hydroponic", "fertigation"}
		methodValid := false
		for _, validMethod := range validMethods {
			if req.ApplicationMethod == validMethod {
				methodValid = true
				break
			}
		}
		if !methodValid {
			return ErrInvalidApplicationMethod
		}
	}

	// Validate growth stage if provided
	if req.GrowthStage != "" {
		validStages := []string{"seedling", "vegetative", "flowering", "fruiting", "pre_harvest"}
		stageValid := false
		for _, validStage := range validStages {
			if req.GrowthStage == validStage {
				stageValid = true
				break
			}
		}
		if !stageValid {
			return ErrInvalidGrowthStage
		}
	}

	if req.EffectivenessRating > 0 && (req.EffectivenessRating < 1 || req.EffectivenessRating > 5) {
		return ErrInvalidEffectivenessRating
	}

	return nil
}
