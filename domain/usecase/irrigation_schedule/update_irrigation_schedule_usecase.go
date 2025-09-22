package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type UpdateIrrigationScheduleUsecaseI interface {
	Execute(ctx context.Context, req *entity.UpdateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error)
}

type UpdateIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

func NewUpdateIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) UpdateIrrigationScheduleUsecaseI {
	return &UpdateIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

func (u *UpdateIrrigationScheduleUsecase) Execute(ctx context.Context, req *entity.UpdateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if irrigation schedule exists
	existingSchedule, err := u.irrigationScheduleRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingSchedule == nil {
		return nil, ErrIrrigationScheduleNotFound
	}

	// Check if schedule name already exists for the growing zone (if name is being updated)
	if req.ScheduleName != "" && req.ScheduleName != existingSchedule.ScheduleName {
		if existingSchedule.GrowingZoneID != "" {
			exists, err := u.irrigationScheduleRepo.CheckScheduleNameExists(ctx, req.ScheduleName, existingSchedule.GrowingZoneID)
			if err != nil {
				return nil, err
			}
			if exists {
				return nil, ErrScheduleNameAlreadyExists
			}
		}
	}

	// Update the schedule
	updatedSchedule, err := u.irrigationScheduleRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return updatedSchedule, nil
}

// validateRequest validates the update request
func (u *UpdateIrrigationScheduleUsecase) validateRequest(req *entity.UpdateIrrigationScheduleRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	if req.ScheduleName == "" {
		return ErrScheduleNameRequired
	}

	if req.GrowingZoneID == "" {
		return ErrInvalidGrowingZoneID
	}

	if req.PlantingCycleID == "" {
		return ErrInvalidPlantingCycleID
	}

	return nil
}
