package irrigation_log

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// UpdateIrrigationLogUsecase handles updating irrigation logs
type UpdateIrrigationLogUsecaseI interface {
	Execute(ctx context.Context, req *entity.UpdateIrrigationLogRequest) (*entity.IrrigationLog, error)
}

type UpdateIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewUpdateIrrigationLogUsecase creates a new instance of UpdateIrrigationLogUsecase
func NewUpdateIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) UpdateIrrigationLogUsecaseI {
	return &UpdateIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

func (u *UpdateIrrigationLogUsecase) Execute(ctx context.Context, req *entity.UpdateIrrigationLogRequest) (*entity.IrrigationLog, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if irrigation log exists
	existingLog, err := u.irrigationLogRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingLog == nil {
		return nil, ErrIrrigationLogNotFound
	}

	// Update the log
	updatedLog, err := u.irrigationLogRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return updatedLog, nil
}

// validateRequest validates the update request
func (u *UpdateIrrigationLogUsecase) validateRequest(req *entity.UpdateIrrigationLogRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	if req.Status != "" {
		// Validate status values
		validStatuses := map[string]bool{
			"completed":       true,
			"failed":          true,
			"interrupted":     true,
			"manual_override": true,
		}
		statusValid := false
		for validStatus := range validStatuses {
			if req.Status == validStatus {
				statusValid = true
				break
			}
		}
		if !statusValid {
			return ErrInvalidStatus
		}
	}

	return nil
}
