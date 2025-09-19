package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// CreateIrrigationScheduleUsecase handles the creation of irrigation schedules
type CreateIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewCreateIrrigationScheduleUsecase creates a new instance of CreateIrrigationScheduleUsecase
func NewCreateIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *CreateIrrigationScheduleUsecase {
	return &CreateIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute creates a new irrigation schedule
func (u *CreateIrrigationScheduleUsecase) Execute(ctx context.Context, req *entity.CreateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if schedule name already exists for the growing zone
	if req.GrowingZoneID != "" {
		exists, err := u.irrigationScheduleRepo.CheckScheduleNameExists(ctx, req.ScheduleName, req.GrowingZoneID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrScheduleNameAlreadyExists
		}
	}

	// Save to repository
	createdSchedule, err := u.irrigationScheduleRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return createdSchedule, nil
}

// validateRequest validates the create request
func (u *CreateIrrigationScheduleUsecase) validateRequest(req *entity.CreateIrrigationScheduleRequest) error {
	if req.ScheduleName == "" {
		return ErrScheduleNameRequired
	}

	if req.GrowingZoneID == "" {
		return ErrInvalidGrowingZoneID
	}

	if req.PlantingCycleID == "" {
		return ErrInvalidPlantingCycleID
	}

	if req.CreatedBy == "" {
		return ErrInvalidCreatedBy
	}

	return nil
}
