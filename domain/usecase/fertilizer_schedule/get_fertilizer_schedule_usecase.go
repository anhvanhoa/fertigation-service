package fertilizer_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

type GetFertilizerScheduleUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.FertilizerSchedule, error)
}

// GetFertilizerScheduleUsecase handles retrieving fertilizer schedules
type GetFertilizerScheduleUsecase struct {
	fertilizerScheduleRepo repository.FertilizerScheduleRepository
}

// NewGetFertilizerScheduleUsecase creates a new instance of GetFertilizerScheduleUsecase
func NewGetFertilizerScheduleUsecase(fertilizerScheduleRepo repository.FertilizerScheduleRepository) GetFertilizerScheduleUsecaseI {
	return &GetFertilizerScheduleUsecase{
		fertilizerScheduleRepo: fertilizerScheduleRepo,
	}
}

// Execute retrieves a fertilizer schedule by ID
func (u *GetFertilizerScheduleUsecase) Execute(ctx context.Context, id string) (*entity.FertilizerSchedule, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	if fertilizerSchedule, err := u.fertilizerScheduleRepo.GetByID(ctx, id); fertilizerSchedule == nil {
		return nil, ErrFertilizerScheduleNotFound
	} else if err != nil {
		return nil, err
	} else {
		return fertilizerSchedule, nil
	}
}
