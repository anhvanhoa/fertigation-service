package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type UpdateFertilizerTypeUsecaseI interface {
	Execute(ctx context.Context, req *entity.UpdateFertilizerTypeRequest) (*entity.FertilizerType, error)
}

type UpdateFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

func NewUpdateFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) UpdateFertilizerTypeUsecaseI {
	return &UpdateFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

func (u *UpdateFertilizerTypeUsecase) Execute(ctx context.Context, req *entity.UpdateFertilizerTypeRequest) (*entity.FertilizerType, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if fertilizer type exists
	existingFertilizerType, err := u.fertilizerTypeRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingFertilizerType == nil {
		return nil, ErrFertilizerTypeNotFound
	}

	// Check if fertilizer type name already exists (if name is being updated)
	if req.Name != "" && req.Name != existingFertilizerType.Name {
		exists, err := u.fertilizerTypeRepo.CheckNameExists(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrFertilizerTypeNameAlreadyExists
		}
	}

	// Check if batch number already exists (if batch number is being updated)
	if req.BatchNumber != "" && req.BatchNumber != existingFertilizerType.BatchNumber {
		exists, err := u.fertilizerTypeRepo.CheckBatchNumberExists(ctx, req.BatchNumber)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrBatchNumberAlreadyExists
		}
	}

	// Update the fertilizer type
	updatedFertilizerType, err := u.fertilizerTypeRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return updatedFertilizerType, nil
}

// validateRequest validates the update request
func (u *UpdateFertilizerTypeUsecase) validateRequest(req *entity.UpdateFertilizerTypeRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	if req.Name != "" {
		return ErrFertilizerTypeNameRequired
	}

	// Validate type if provided
	if req.Type != "" {
		validTypes := []string{"organic", "chemical", "liquid", "granular", "powder"}
		typeValid := false
		for _, validType := range validTypes {
			if req.Type == validType {
				typeValid = true
				break
			}
		}
		if !typeValid {
			return ErrInvalidFertilizerType
		}
	}

	// Validate application method if provided
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

	return nil
}
