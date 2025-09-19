package irrigation_schedule

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
)

// ListIrrigationScheduleUsecase handles listing irrigation schedules
type ListIrrigationScheduleUsecase struct {
	irrigationScheduleRepo repository.IrrigationScheduleRepository
}

// NewListIrrigationScheduleUsecase creates a new instance of ListIrrigationScheduleUsecase
func NewListIrrigationScheduleUsecase(irrigationScheduleRepo repository.IrrigationScheduleRepository) *ListIrrigationScheduleUsecase {
	return &ListIrrigationScheduleUsecase{
		irrigationScheduleRepo: irrigationScheduleRepo,
	}
}

// Execute retrieves a list of irrigation schedules with filtering and pagination
func (u *ListIrrigationScheduleUsecase) Execute(ctx context.Context, filter *entity.IrrigationScheduleFilter) (*entity.ListIrrigationSchedulesResponse, error) {
	// Set default pagination values
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	// Set default sort values
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}
	if filter.SortOrder == "" {
		filter.SortOrder = "desc"
	}

	// Get irrigation schedules from repository
	response, err := u.irrigationScheduleRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	if response.Total > 0 {
		response.TotalPages = (response.Total + filter.Limit - 1) / filter.Limit
	}

	return response, nil
}
