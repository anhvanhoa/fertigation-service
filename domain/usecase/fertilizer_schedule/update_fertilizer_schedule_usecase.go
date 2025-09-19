package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type UpdateFertilizerScheduleUsecaseI interface {
	Execute(ctx context.Context, req *entity.UpdateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error)
}

// UpdateFertilizerScheduleUsecase handles updating fertilizer schedules
type UpdateFertilizerScheduleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
}

// NewUpdateFertilizerScheduleUsecase creates a new instance of UpdateFertilizerScheduleUsecase
func NewUpdateFertilizerScheduleUsecase(fertilizerScheduleRepo repository.FertilizerScheduleRepository) UpdateFertilizerScheduleUsecaseI {
	return &UpdateFertilizerScheduleUsecase{
		fertilizerScheduleRepo: fertilizerScheduleRepo,
	}
}

// Execute updates an existing fertilizer schedule
func (u *UpdateFertilizerScheduleUsecase) Execute(ctx context.Context, req *entity.UpdateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if fertilizer schedule exists
	existingSchedule, err := u.fertilizerScheduleRepo.GetByID(ctx, req.ID)
	if existingSchedule == nil {
		return nil, ErrFertilizerScheduleNotFound
	}
	if err != nil {
		return nil, err
	}

	// Update the schedule
	updatedSchedule, err := u.fertilizerScheduleRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return updatedSchedule, nil
}

// validateRequest validates the update request
func (u *UpdateFertilizerScheduleUsecase) validateRequest(req *entity.UpdateFertilizerScheduleRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	if req.PlantingCycleID != "" {
		return ErrInvalidPlantingCycleID
	}

	if req.FertilizerTypeID != "" {
		return ErrInvalidFertilizerTypeID
	}

	// Validate application method if provided
	if req.ApplicationMethod != "" {
		validMethods := [4]string{"foliar", "soil", "hydroponic", "fertigation"}
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
		validStages := [5]string{"seedling", "vegetative", "flowering", "fruiting", "pre_harvest"}
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

	// Validate effectiveness rating if provided
	if req.EffectivenessRating > 0 && (req.EffectivenessRating < 1 || req.EffectivenessRating > 5) {
		return ErrInvalidEffectivenessRating
	}

	return nil
}
