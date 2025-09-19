package irrigation_log

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// CreateIrrigationLogUsecase handles the creation of irrigation logs
type CreateIrrigationLogUsecase struct {
	irrigationLogRepo repository.IrrigationLogRepository
}

// NewCreateIrrigationLogUsecase creates a new instance of CreateIrrigationLogUsecase
func NewCreateIrrigationLogUsecase(irrigationLogRepo repository.IrrigationLogRepository) *CreateIrrigationLogUsecase {
	return &CreateIrrigationLogUsecase{
		irrigationLogRepo: irrigationLogRepo,
	}
}

// Execute creates a new irrigation log
func (u *CreateIrrigationLogUsecase) Execute(ctx context.Context, req *entity.CreateIrrigationLogRequest) (*entity.IrrigationLog, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Save to repository
	createdLog, err := u.irrigationLogRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return createdLog, nil
}

// validateRequest validates the create request
func (u *CreateIrrigationLogUsecase) validateRequest(req *entity.CreateIrrigationLogRequest) error {
	if req.IrrigationScheduleID == "" {
		return ErrInvalidIrrigationScheduleID
	}

	if req.DeviceID == "" {
		return ErrInvalidDeviceID
	}

	if req.Status == "" {
		return ErrInvalidStatus
	}

	if req.CreatedBy == "" {
		return ErrInvalidCreatedBy
	}

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

	return nil
}
