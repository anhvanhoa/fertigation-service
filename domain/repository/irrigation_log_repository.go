package repository

import (
	"context"
	"fertigation-Service/domain/entity"
)

// IrrigationLogRepository defines the interface for irrigation log data operations
type IrrigationLogRepository interface {
	// Create creates a new irrigation log
	Create(ctx context.Context, req *entity.CreateIrrigationLogRequest) (*entity.IrrigationLog, error)

	// GetByID retrieves an irrigation log by ID
	GetByID(ctx context.Context, id string) (*entity.IrrigationLog, error)

	// Update updates an existing irrigation log
	Update(ctx context.Context, req *entity.UpdateIrrigationLogRequest) (*entity.IrrigationLog, error)

	// Delete removes an irrigation log by ID
	Delete(ctx context.Context, id string) error

	// List retrieves irrigation logs with filtering and pagination
	List(ctx context.Context, filter *entity.IrrigationLogFilter) (*entity.ListIrrigationLogsResponse, error)

	// GetByScheduleID retrieves irrigation logs by schedule ID
	GetByScheduleID(ctx context.Context, scheduleID string) ([]*entity.IrrigationLog, error)

	// GetByDeviceID retrieves irrigation logs by device ID
	GetByDeviceID(ctx context.Context, deviceID string) ([]*entity.IrrigationLog, error)

	// GetByStatus retrieves irrigation logs by status
	GetByStatus(ctx context.Context, status string) ([]*entity.IrrigationLog, error)

	// GetByCreator retrieves irrigation logs created by a specific user
	GetByCreator(ctx context.Context, createdBy string) ([]*entity.IrrigationLog, error)

	// GetLogsByDateRange retrieves logs within a specific date range
	GetLogsByDateRange(ctx context.Context, from, to string) ([]*entity.IrrigationLog, error)

	// GetLogsByExecutionDate retrieves logs by execution date range
	GetLogsByExecutionDate(ctx context.Context, startedAtFrom, startedAtTo string) ([]*entity.IrrigationLog, error)

	// GetFailedLogs retrieves all failed irrigation logs
	GetFailedLogs(ctx context.Context) ([]*entity.IrrigationLog, error)

	// GetCompletedLogs retrieves all completed irrigation logs
	GetCompletedLogs(ctx context.Context) ([]*entity.IrrigationLog, error)

	// GetInterruptedLogs retrieves all interrupted irrigation logs
	GetInterruptedLogs(ctx context.Context) ([]*entity.IrrigationLog, error)

	// GetManualOverrideLogs retrieves all manual override irrigation logs
	GetManualOverrideLogs(ctx context.Context) ([]*entity.IrrigationLog, error)

	// GetLogsByWaterUsage retrieves logs within a water usage range
	GetLogsByWaterUsage(ctx context.Context, minUsage, maxUsage float64) ([]*entity.IrrigationLog, error)

	// GetLogsByDuration retrieves logs within a duration range
	GetLogsByDuration(ctx context.Context, minDuration, maxDuration int) ([]*entity.IrrigationLog, error)

	// GetLogsByPressure retrieves logs within a water pressure range
	GetLogsByPressure(ctx context.Context, minPressure, maxPressure float64) ([]*entity.IrrigationLog, error)

	// Count returns the total number of irrigation logs matching the filter
	Count(ctx context.Context, filter *entity.IrrigationLogFilter) (int, error)

	// GetLogStatistics returns statistics about irrigation logs
	GetLogStatistics(ctx context.Context) (*entity.IrrigationLogStatistics, error)

	// GetScheduleLogStatistics returns statistics for logs of a specific schedule
	GetScheduleLogStatistics(ctx context.Context, scheduleID string) (*entity.IrrigationLogStatistics, error)

	// GetDeviceLogStatistics returns statistics for logs of a specific device
	GetDeviceLogStatistics(ctx context.Context, deviceID string) (*entity.IrrigationLogStatistics, error)

	// GetWaterUsageReport returns water usage report for a date range
	GetWaterUsageReport(ctx context.Context, from, to string) (*entity.WaterUsageReport, error)

	// GetEfficiencyReport returns efficiency report for irrigation logs
	GetEfficiencyReport(ctx context.Context, from, to string) (*entity.EfficiencyReport, error)

	// BulkUpdateStatus updates the status of multiple logs
	BulkUpdateStatus(ctx context.Context, ids []string, status string) error

	// GetRecentLogs retrieves recent irrigation logs
	GetRecentLogs(ctx context.Context, limit int) ([]*entity.IrrigationLog, error)
}
