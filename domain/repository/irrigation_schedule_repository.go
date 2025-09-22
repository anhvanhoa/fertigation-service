package repository

import (
	"context"
	"fertigation-Service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
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
	List(ctx context.Context, filter *entity.IrrigationScheduleFilter) ([]*entity.IrrigationSchedule, int64, error)

	GetByGrowingZoneID(ctx context.Context, growingZoneID string, filter *common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetByPlantingCycleID retrieves irrigation schedules by planting cycle ID
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string) ([]*entity.IrrigationSchedule, error)

	GetActiveSchedules(ctx context.Context, isActive bool, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetSchedulesByType retrieves irrigation schedules by irrigation type
	GetSchedulesByType(ctx context.Context, irrigationType string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetSchedulesByFrequency retrieves irrigation schedules by frequency
	GetSchedulesByFrequency(ctx context.Context, frequency string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetSchedulesWithFertilizerMix retrieves irrigation schedules that include fertilizer mixing
	GetSchedulesWithFertilizerMix(ctx context.Context, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetSchedulesByCreator retrieves irrigation schedules created by a specific user
	GetSchedulesByCreator(ctx context.Context, createdBy string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// GetSchedulesForExecution retrieves schedules that need to be executed within a time range
	GetSchedulesForExecution(ctx context.Context, from, to string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// UpdateNextExecution updates the next execution time for a schedule
	UpdateNextExecution(ctx context.Context, id string, nextExecution string) error

	// UpdateLastExecuted updates the last executed time for a schedule
	UpdateLastExecuted(ctx context.Context, id string, lastExecuted string) error

	// Count returns the total number of irrigation schedules matching the filter
	Count(ctx context.Context, filter *entity.IrrigationScheduleFilter) (int, error)

	// CheckScheduleNameExists checks if a schedule name already exists for a specific growing zone
	CheckScheduleNameExists(ctx context.Context, scheduleName, growingZoneID string) (bool, error)

	// GetSchedulesByDateRange retrieves schedules within a specific date range
	GetSchedulesByDateRange(ctx context.Context, from, to string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error)

	// BulkUpdateStatus updates the status of multiple schedules
	BulkUpdateStatus(ctx context.Context, ids []string, isActive bool) error

	GetScheduleStatistics(ctx context.Context) (*entity.IrrigationScheduleStatistics, error)
}
