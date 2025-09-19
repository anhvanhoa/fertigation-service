package fertilizer_type

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// CreateFertilizerTypeUsecase handles the creation of fertilizer types
type CreateFertilizerTypeUsecase struct {
	fertilizerTypeRepo repository.FertilizerTypeRepository
}

// NewCreateFertilizerTypeUsecase creates a new instance of CreateFertilizerTypeUsecase
func NewCreateFertilizerTypeUsecase(fertilizerTypeRepo repository.FertilizerTypeRepository) *CreateFertilizerTypeUsecase {
	return &CreateFertilizerTypeUsecase{
		fertilizerTypeRepo: fertilizerTypeRepo,
	}
}

// Execute creates a new fertilizer type
func (u *CreateFertilizerTypeUsecase) Execute(ctx context.Context, req *entity.CreateFertilizerTypeRequest) (*entity.FertilizerType, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if fertilizer type name already exists
	exists, err := u.fertilizerTypeRepo.CheckNameExists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrFertilizerTypeNameAlreadyExists
	}

	// Check if batch number already exists (if provided)
	if req.BatchNumber != "" {
		exists, err := u.fertilizerTypeRepo.CheckBatchNumberExists(ctx, req.BatchNumber)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrBatchNumberAlreadyExists
		}
	}

	// Set default status if not provided
	if req.Status == "" {
		req.Status = "active"
	}

	// Save to repository
	createdFertilizerType, err := u.fertilizerTypeRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return createdFertilizerType, nil
}

// validateRequest validates the create request
func (u *CreateFertilizerTypeUsecase) validateRequest(req *entity.CreateFertilizerTypeRequest) error {
	if req.Name == "" {
		return ErrFertilizerTypeNameRequired
	}

	if req.CreatedBy == "" {
		return ErrInvalidCreatedBy
	}

	// Validate type if provided
	if req.Type != "" {
		validTypes := map[string]bool{
			"organic":  true,
			"chemical": true,
			"liquid":   true,
			"granular": true,
			"powder":   true,
		}
		typeValid := false
		for validType := range validTypes {
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
		validMethods := map[string]bool{
			"foliar":      true,
			"soil":        true,
			"hydroponic":  true,
			"fertigation": true,
		}
		methodValid := false
		for validMethod := range validMethods {
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
