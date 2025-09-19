package repository

import (
	"context"
	"fertigation-Service/domain/entity"
)

// IrrigationScheduleRepository defines the interface for irrigation schedule data operations
type IrrigationScheduleRepository interface {
	// Create creates a new irrigation schedule
	Create(ctx context.Context, req *entity.CreateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error)

	// GetByID retrieves an irrigation schedule by ID
	GetByID(ctx context.Context, id string) (*entity.IrrigationSchedule, error)

	// Update updates an existing irrigation schedule
	Update(ctx context.Context, req *entity.UpdateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error)

	// Delete removes an irrigation schedule by ID
	Delete(ctx context.Context, id string) error

	// List retrieves irrigation schedules with filtering and pagination
	List(ctx context.Context, filter *entity.IrrigationScheduleFilter) (*entity.ListIrrigationSchedulesResponse, error)

	// GetByGrowingZoneID retrieves irrigation schedules by growing zone ID
	GetByGrowingZoneID(ctx context.Context, growingZoneID string) ([]*entity.IrrigationSchedule, error)

	// GetByPlantingCycleID retrieves irrigation schedules by planting cycle ID
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string) ([]*entity.IrrigationSchedule, error)

	// GetActiveSchedules retrieves all active irrigation schedules
	GetActiveSchedules(ctx context.Context) ([]*entity.IrrigationSchedule, error)

	// GetSchedulesByType retrieves irrigation schedules by irrigation type
	GetSchedulesByType(ctx context.Context, irrigationType string) ([]*entity.IrrigationSchedule, error)

	// GetSchedulesByFrequency retrieves irrigation schedules by frequency
	GetSchedulesByFrequency(ctx context.Context, frequency string) ([]*entity.IrrigationSchedule, error)

	// GetSchedulesWithFertilizerMix retrieves irrigation schedules that include fertilizer mixing
	GetSchedulesWithFertilizerMix(ctx context.Context) ([]*entity.IrrigationSchedule, error)

	// GetSchedulesByCreator retrieves irrigation schedules created by a specific user
	GetSchedulesByCreator(ctx context.Context, createdBy string) ([]*entity.IrrigationSchedule, error)

	// GetSchedulesForExecution retrieves schedules that need to be executed within a time range
	GetSchedulesForExecution(ctx context.Context, from, to string) ([]*entity.IrrigationSchedule, error)

	// UpdateNextExecution updates the next execution time for a schedule
	UpdateNextExecution(ctx context.Context, id string, nextExecution string) error

	// UpdateLastExecuted updates the last executed time for a schedule
	UpdateLastExecuted(ctx context.Context, id string, lastExecuted string) error

	// Count returns the total number of irrigation schedules matching the filter
	Count(ctx context.Context, filter *entity.IrrigationScheduleFilter) (int, error)

	// CheckScheduleNameExists checks if a schedule name already exists for a specific growing zone
	CheckScheduleNameExists(ctx context.Context, scheduleName, growingZoneID string) (bool, error)

	// GetSchedulesByDateRange retrieves schedules within a specific date range
	GetSchedulesByDateRange(ctx context.Context, from, to string) ([]*entity.IrrigationSchedule, error)

	// BulkUpdateStatus updates the status of multiple schedules
	BulkUpdateStatus(ctx context.Context, ids []string, isActive bool) error

	// GetScheduleStatistics returns statistics about irrigation schedules
	GetScheduleStatistics(ctx context.Context) (*entity.IrrigationScheduleStatistics, error)
}
