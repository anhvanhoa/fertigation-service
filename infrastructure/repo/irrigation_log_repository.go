package repo

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/go-pg/pg/v10"
)

type irrigationLogRepository struct {
	db *pg.DB
}

// NewIrrigationLogRepository creates a new irrigation log repository
func NewIrrigationLogRepository(db *pg.DB) repository.IrrigationLogRepository {
	return &irrigationLogRepository{
		db: db,
	}
}

// Create creates a new irrigation log
func (r *irrigationLogRepository) Create(ctx context.Context, req *entity.CreateIrrigationLogRequest) (*entity.IrrigationLog, error) {
	log := &entity.IrrigationLog{
		IrrigationScheduleID:   req.IrrigationScheduleID,
		DeviceID:               req.DeviceID,
		StartedAt:              req.StartedAt,
		EndedAt:                req.EndedAt,
		PlannedDurationMinutes: req.PlannedDurationMinutes,
		ActualDurationMinutes:  req.ActualDurationMinutes,
		WaterUsedLiters:        req.WaterUsedLiters,
		WaterPressure:          req.WaterPressure,
		Status:                 req.Status,
		FailureReason:          req.FailureReason,
		Notes:                  req.Notes,
		CreatedBy:              req.CreatedBy,
		CreatedAt:              time.Now(),
	}

	_, err := r.db.ModelContext(ctx, log).Insert()
	if err != nil {
		return nil, err
	}

	return log, nil
}

// GetByID retrieves an irrigation log by ID
func (r *irrigationLogRepository) GetByID(ctx context.Context, id string) (*entity.IrrigationLog, error) {
	log := &entity.IrrigationLog{}
	err := r.db.ModelContext(ctx, log).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return log, nil
}

// Update updates an existing irrigation log
func (r *irrigationLogRepository) Update(ctx context.Context, req *entity.UpdateIrrigationLogRequest) (*entity.IrrigationLog, error) {
	log := &entity.IrrigationLog{
		ID:                     req.ID,
		IrrigationScheduleID:   req.IrrigationScheduleID,
		DeviceID:               req.DeviceID,
		StartedAt:              req.StartedAt,
		EndedAt:                req.EndedAt,
		PlannedDurationMinutes: req.PlannedDurationMinutes,
		ActualDurationMinutes:  req.ActualDurationMinutes,
		WaterUsedLiters:        req.WaterUsedLiters,
		WaterPressure:          req.WaterPressure,
		Status:                 req.Status,
		FailureReason:          req.FailureReason,
		Notes:                  req.Notes,
	}

	_, err := r.db.ModelContext(ctx, log).Where("id = ?", req.ID).Update()
	if err != nil {
		return nil, err
	}

	return log, nil
}

// Delete removes an irrigation log by ID
func (r *irrigationLogRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).Where("id = ?", id).Delete()
	return err
}

// List retrieves irrigation logs with filtering and pagination
func (r *irrigationLogRepository) List(ctx context.Context, filter *entity.IrrigationLogFilter) (*entity.ListIrrigationLogsResponse, error) {
	var logs []*entity.IrrigationLog
	query := r.db.ModelContext(ctx, &logs)

	// Apply filters
	if filter.IrrigationScheduleID != "" {
		query = query.Where("irrigation_schedule_id = ?", filter.IrrigationScheduleID)
	}
	if filter.DeviceID != "" {
		query = query.Where("device_id = ?", filter.DeviceID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.StartedAtFrom != nil {
		query = query.Where("started_at >= ?", filter.StartedAtFrom)
	}
	if filter.StartedAtTo != nil {
		query = query.Where("started_at <= ?", filter.StartedAtTo)
	}
	if filter.EndedAtFrom != nil {
		query = query.Where("ended_at >= ?", filter.EndedAtFrom)
	}
	if filter.EndedAtTo != nil {
		query = query.Where("ended_at <= ?", filter.EndedAtTo)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}

	// Get total count
	total, err := query.Count()
	if err != nil {
		return nil, err
	}

	// Apply pagination and sorting
	query = query.Order(filter.SortBy + " " + filter.SortOrder)
	query = query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit)

	err = query.Select()
	if err != nil {
		return nil, err
	}

	// Convert to response format
	var responses []entity.IrrigationLogResponse
	for _, log := range logs {
		responses = append(responses, entity.IrrigationLogResponse{
			ID:                     log.ID,
			IrrigationScheduleID:   log.IrrigationScheduleID,
			DeviceID:               log.DeviceID,
			StartedAt:              log.StartedAt,
			EndedAt:                log.EndedAt,
			PlannedDurationMinutes: log.PlannedDurationMinutes,
			ActualDurationMinutes:  log.ActualDurationMinutes,
			WaterUsedLiters:        log.WaterUsedLiters,
			WaterPressure:          log.WaterPressure,
			Status:                 log.Status,
			FailureReason:          log.FailureReason,
			Notes:                  log.Notes,
			CreatedBy:              log.CreatedBy,
			CreatedAt:              log.CreatedAt,
		})
	}

	return &entity.ListIrrigationLogsResponse{
		IrrigationLogs: responses,
		Total:          total,
		Page:           filter.Page,
		Limit:          filter.Limit,
	}, nil
}

// GetByScheduleID retrieves irrigation logs by schedule ID
func (r *irrigationLogRepository) GetByScheduleID(ctx context.Context, scheduleID string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("irrigation_schedule_id = ?", scheduleID).Select()
	return logs, err
}

// GetByDeviceID retrieves irrigation logs by device ID
func (r *irrigationLogRepository) GetByDeviceID(ctx context.Context, deviceID string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("device_id = ?", deviceID).Select()
	return logs, err
}

// GetByStatus retrieves irrigation logs by status
func (r *irrigationLogRepository) GetByStatus(ctx context.Context, status string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("status = ?", status).Select()
	return logs, err
}

// GetByCreator retrieves irrigation logs created by a specific user
func (r *irrigationLogRepository) GetByCreator(ctx context.Context, createdBy string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("created_by = ?", createdBy).Select()
	return logs, err
}

// GetLogsByDateRange retrieves logs within a specific date range
func (r *irrigationLogRepository) GetLogsByDateRange(ctx context.Context, from, to string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("created_at >= ? AND created_at <= ?", from, to).
		Select()
	return logs, err
}

// GetLogsByExecutionDate retrieves logs by execution date range
func (r *irrigationLogRepository) GetLogsByExecutionDate(ctx context.Context, startedAtFrom, startedAtTo string) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("started_at >= ? AND started_at <= ?", startedAtFrom, startedAtTo).
		Select()
	return logs, err
}

// GetFailedLogs retrieves all failed irrigation logs
func (r *irrigationLogRepository) GetFailedLogs(ctx context.Context) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("status = ?", "failed").Select()
	return logs, err
}

// GetCompletedLogs retrieves all completed irrigation logs
func (r *irrigationLogRepository) GetCompletedLogs(ctx context.Context) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("status = ?", "completed").Select()
	return logs, err
}

// GetInterruptedLogs retrieves all interrupted irrigation logs
func (r *irrigationLogRepository) GetInterruptedLogs(ctx context.Context) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("status = ?", "interrupted").Select()
	return logs, err
}

// GetManualOverrideLogs retrieves all manual override irrigation logs
func (r *irrigationLogRepository) GetManualOverrideLogs(ctx context.Context) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).Where("status = ?", "manual_override").Select()
	return logs, err
}

// GetLogsByWaterUsage retrieves logs within a water usage range
func (r *irrigationLogRepository) GetLogsByWaterUsage(ctx context.Context, minUsage, maxUsage float64) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("water_used_liters >= ? AND water_used_liters <= ?", minUsage, maxUsage).
		Select()
	return logs, err
}

// GetLogsByDuration retrieves logs within a duration range
func (r *irrigationLogRepository) GetLogsByDuration(ctx context.Context, minDuration, maxDuration int) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("actual_duration_minutes >= ? AND actual_duration_minutes <= ?", minDuration, maxDuration).
		Select()
	return logs, err
}

// GetLogsByPressure retrieves logs within a water pressure range
func (r *irrigationLogRepository) GetLogsByPressure(ctx context.Context, minPressure, maxPressure float64) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("water_pressure >= ? AND water_pressure <= ?", minPressure, maxPressure).
		Select()
	return logs, err
}

// Count returns the total number of irrigation logs matching the filter
func (r *irrigationLogRepository) Count(ctx context.Context, filter *entity.IrrigationLogFilter) (int, error) {
	query := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil))

	// Apply same filters as List method
	if filter.IrrigationScheduleID != "" {
		query = query.Where("irrigation_schedule_id = ?", filter.IrrigationScheduleID)
	}
	if filter.DeviceID != "" {
		query = query.Where("device_id = ?", filter.DeviceID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.StartedAtFrom != nil {
		query = query.Where("started_at >= ?", filter.StartedAtFrom)
	}
	if filter.StartedAtTo != nil {
		query = query.Where("started_at <= ?", filter.StartedAtTo)
	}
	if filter.EndedAtFrom != nil {
		query = query.Where("ended_at >= ?", filter.EndedAtFrom)
	}
	if filter.EndedAtTo != nil {
		query = query.Where("ended_at <= ?", filter.EndedAtTo)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}

	return query.Count()
}

// GetLogStatistics returns statistics about irrigation logs
func (r *irrigationLogRepository) GetLogStatistics(ctx context.Context) (*entity.IrrigationLogStatistics, error) {
	stats := &entity.IrrigationLogStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalLogs = total

	completed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("status = ?", "completed").Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedLogs = completed

	failed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("status = ?", "failed").Count()
	if err != nil {
		return nil, err
	}
	stats.FailedLogs = failed

	return stats, nil
}

// GetScheduleLogStatistics returns statistics for logs of a specific schedule
func (r *irrigationLogRepository) GetScheduleLogStatistics(ctx context.Context, scheduleID string) (*entity.IrrigationLogStatistics, error) {
	stats := &entity.IrrigationLogStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("irrigation_schedule_id = ?", scheduleID).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalLogs = total

	completed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("irrigation_schedule_id = ? AND status = ?", scheduleID, "completed").Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedLogs = completed

	failed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("irrigation_schedule_id = ? AND status = ?", scheduleID, "failed").Count()
	if err != nil {
		return nil, err
	}
	stats.FailedLogs = failed

	return stats, nil
}

// GetDeviceLogStatistics returns statistics for logs of a specific device
func (r *irrigationLogRepository) GetDeviceLogStatistics(ctx context.Context, deviceID string) (*entity.IrrigationLogStatistics, error) {
	stats := &entity.IrrigationLogStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("device_id = ?", deviceID).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalLogs = total

	completed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("device_id = ? AND status = ?", deviceID, "completed").Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedLogs = completed

	failed, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("device_id = ? AND status = ?", deviceID, "failed").Count()
	if err != nil {
		return nil, err
	}
	stats.FailedLogs = failed

	return stats, nil
}

// GetWaterUsageReport returns water usage report for a date range
func (r *irrigationLogRepository) GetWaterUsageReport(ctx context.Context, from, to string) (*entity.WaterUsageReport, error) {
	report := &entity.WaterUsageReport{}

	// Get total water usage
	var totalUsage float64
	err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("started_at >= ? AND started_at <= ?", from, to).
		Where("status = ?", "completed").
		ColumnExpr("SUM(water_used_liters)").
		Select(&totalUsage)
	if err != nil {
		return nil, err
	}
	report.TotalWaterUsed = totalUsage

	// Get average water usage
	var avgUsage float64
	err = r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("started_at >= ? AND started_at <= ?", from, to).
		Where("status = ?", "completed").
		ColumnExpr("AVG(water_used_liters)").
		Select(&avgUsage)
	if err != nil {
		return nil, err
	}
	report.AverageWaterUsed = avgUsage

	return report, nil
}

// GetEfficiencyReport returns efficiency report for irrigation logs
func (r *irrigationLogRepository) GetEfficiencyReport(ctx context.Context, from, to string) (*entity.EfficiencyReport, error) {
	report := &entity.EfficiencyReport{}

	// Get total logs
	total, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("started_at >= ? AND started_at <= ?", from, to).
		Count()
	if err != nil {
		return nil, err
	}
	report.TotalPlannedDuration = total

	// Get successful executions
	successful, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Where("started_at >= ? AND started_at <= ?", from, to).
		Where("status = ?", "completed").
		Count()
	if err != nil {
		return nil, err
	}
	report.TotalActualDuration = successful

	// Calculate efficiency percentage
	if total > 0 {
		report.DurationEfficiency = float64(successful) / float64(total) * 100
	}

	return report, nil
}

// BulkUpdateStatus updates the status of multiple logs
func (r *irrigationLogRepository) BulkUpdateStatus(ctx context.Context, ids []string, status string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationLog)(nil)).
		Set("status = ?", status).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

// GetRecentLogs retrieves recent irrigation logs
func (r *irrigationLogRepository) GetRecentLogs(ctx context.Context, limit int) ([]*entity.IrrigationLog, error) {
	var logs []*entity.IrrigationLog
	err := r.db.ModelContext(ctx, &logs).
		Order("created_at DESC").
		Limit(limit).
		Select()
	return logs, err
}
