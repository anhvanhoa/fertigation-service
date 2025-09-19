package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// GetSchedulesByGrowingZoneUsecase handles retrieving irrigation schedules by growing zone
type GetSchedulesByGrowingZoneUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewGetSchedulesByGrowingZoneUsecase creates a new instance of GetSchedulesByGrowingZoneUsecase
func NewGetSchedulesByGrowingZoneUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *GetSchedulesByGrowingZoneUsecase {
	return &GetSchedulesByGrowingZoneUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute retrieves irrigation schedules by growing zone ID
func (u *GetSchedulesByGrowingZoneUsecase) Execute(ctx context.Context, growingZoneID string) ([]*entity.IrrigationSchedule, error) {
	if growingZoneID == "" {
		return nil, ErrInvalidGrowingZoneID
	}

	schedules, err := u.irrigationScheduleRepo.GetByGrowingZoneID(ctx, growingZoneID)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
