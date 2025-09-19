package repo

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/go-pg/pg/v10"
)

type irrigationScheduleRepository struct {
	db *pg.DB
}

// NewIrrigationScheduleRepository creates a new irrigation schedule repository
func NewIrrigationScheduleRepository(db *pg.DB) repository.IrrigationScheduleRepository {
	return &irrigationScheduleRepository{
		db: db,
	}
}

// Create creates a new irrigation schedule
func (r *irrigationScheduleRepository) Create(ctx context.Context, req *entity.CreateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error) {
	schedule := &entity.IrrigationSchedule{
		GrowingZoneID:     req.GrowingZoneID,
		PlantingCycleID:   req.PlantingCycleID,
		ScheduleName:      req.ScheduleName,
		IrrigationType:    req.IrrigationType,
		StartTime:         req.StartTime,
		DurationMinutes:   req.DurationMinutes,
		Frequency:         req.Frequency,
		DaysOfWeek:        req.DaysOfWeek,
		WaterAmountLiters: req.WaterAmountLiters,
		FertilizerMix:     req.FertilizerMix,
		IsActive:          req.IsActive,
		CreatedBy:         req.CreatedBy,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	_, err := r.db.ModelContext(ctx, schedule).Insert()
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// GetByID retrieves an irrigation schedule by ID
func (r *irrigationScheduleRepository) GetByID(ctx context.Context, id string) (*entity.IrrigationSchedule, error) {
	schedule := &entity.IrrigationSchedule{}
	err := r.db.ModelContext(ctx, schedule).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return schedule, nil
}

// Update updates an existing irrigation schedule
func (r *irrigationScheduleRepository) Update(ctx context.Context, req *entity.UpdateIrrigationScheduleRequest) (*entity.IrrigationSchedule, error) {
	schedule := &entity.IrrigationSchedule{
		ID:                req.ID,
		GrowingZoneID:     req.GrowingZoneID,
		PlantingCycleID:   req.PlantingCycleID,
		ScheduleName:      req.ScheduleName,
		IrrigationType:    req.IrrigationType,
		StartTime:         req.StartTime,
		DurationMinutes:   req.DurationMinutes,
		Frequency:         req.Frequency,
		DaysOfWeek:        req.DaysOfWeek,
		WaterAmountLiters: req.WaterAmountLiters,
		FertilizerMix:     req.FertilizerMix,
		IsActive:          req.IsActive,
		LastExecuted:      req.LastExecuted,
		NextExecution:     req.NextExecution,
		UpdatedAt:         time.Now(),
	}

	_, err := r.db.ModelContext(ctx, schedule).Where("id = ?", req.ID).Update()
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// Delete removes an irrigation schedule by ID
func (r *irrigationScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).Where("id = ?", id).Delete()
	return err
}

// List retrieves irrigation schedules with filtering and pagination
func (r *irrigationScheduleRepository) List(ctx context.Context, filter *entity.IrrigationScheduleFilter) (*entity.ListIrrigationSchedulesResponse, error) {
	var schedules []*entity.IrrigationSchedule
	query := r.db.ModelContext(ctx, &schedules)

	// Apply filters
	if filter.GrowingZoneID != "" {
		query = query.Where("growing_zone_id = ?", filter.GrowingZoneID)
	}
	if filter.PlantingCycleID != "" {
		query = query.Where("planting_cycle_id = ?", filter.PlantingCycleID)
	}
	if filter.ScheduleName != "" {
		query = query.Where("schedule_name ILIKE ?", "%"+filter.ScheduleName+"%")
	}
	if filter.IrrigationType != "" {
		query = query.Where("irrigation_type = ?", filter.IrrigationType)
	}
	if filter.Frequency != "" {
		query = query.Where("frequency = ?", filter.Frequency)
	}
	if filter.IsActive {
		query = query.Where("is_active = ?", filter.IsActive)
	}
	if filter.FertilizerMix {
		query = query.Where("fertilizer_mix = ?", filter.FertilizerMix)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}
	if filter.NextExecutionFrom != nil {
		query = query.Where("next_execution >= ?", filter.NextExecutionFrom)
	}
	if filter.NextExecutionTo != nil {
		query = query.Where("next_execution <= ?", filter.NextExecutionTo)
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
	var responses []entity.IrrigationScheduleResponse
	for _, schedule := range schedules {
		responses = append(responses, entity.IrrigationScheduleResponse{
			ID:                schedule.ID,
			GrowingZoneID:     schedule.GrowingZoneID,
			PlantingCycleID:   schedule.PlantingCycleID,
			ScheduleName:      schedule.ScheduleName,
			IrrigationType:    schedule.IrrigationType,
			StartTime:         schedule.StartTime,
			DurationMinutes:   schedule.DurationMinutes,
			Frequency:         schedule.Frequency,
			DaysOfWeek:        schedule.DaysOfWeek,
			WaterAmountLiters: schedule.WaterAmountLiters,
			FertilizerMix:     schedule.FertilizerMix,
			IsActive:          schedule.IsActive,
			LastExecuted:      schedule.LastExecuted,
			NextExecution:     schedule.NextExecution,
			CreatedBy:         schedule.CreatedBy,
			CreatedAt:         schedule.CreatedAt,
			UpdatedAt:         schedule.UpdatedAt,
		})
	}

	return &entity.ListIrrigationSchedulesResponse{
		IrrigationSchedules: responses,
		Total:               total,
		Page:                filter.Page,
		Limit:               filter.Limit,
	}, nil
}

// GetByGrowingZoneID retrieves irrigation schedules by growing zone ID
func (r *irrigationScheduleRepository) GetByGrowingZoneID(ctx context.Context, growingZoneID string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("growing_zone_id = ?", growingZoneID).Select()
	return schedules, err
}

// GetByPlantingCycleID retrieves irrigation schedules by planting cycle ID
func (r *irrigationScheduleRepository) GetByPlantingCycleID(ctx context.Context, plantingCycleID string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("planting_cycle_id = ?", plantingCycleID).Select()
	return schedules, err
}

// GetActiveSchedules retrieves all active irrigation schedules
func (r *irrigationScheduleRepository) GetActiveSchedules(ctx context.Context) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("is_active = ?", true).Select()
	return schedules, err
}

// GetSchedulesByType retrieves irrigation schedules by irrigation type
func (r *irrigationScheduleRepository) GetSchedulesByType(ctx context.Context, irrigationType string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("irrigation_type = ?", irrigationType).Select()
	return schedules, err
}

// GetSchedulesByFrequency retrieves irrigation schedules by frequency
func (r *irrigationScheduleRepository) GetSchedulesByFrequency(ctx context.Context, frequency string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("frequency = ?", frequency).Select()
	return schedules, err
}

// GetSchedulesWithFertilizerMix retrieves irrigation schedules that include fertilizer mixing
func (r *irrigationScheduleRepository) GetSchedulesWithFertilizerMix(ctx context.Context) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("fertilizer_mix = ?", true).Select()
	return schedules, err
}

// GetSchedulesByCreator retrieves irrigation schedules created by a specific user
func (r *irrigationScheduleRepository) GetSchedulesByCreator(ctx context.Context, createdBy string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("created_by = ?", createdBy).Select()
	return schedules, err
}

// GetSchedulesForExecution retrieves schedules that need to be executed within a time range
func (r *irrigationScheduleRepository) GetSchedulesForExecution(ctx context.Context, from, to string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).
		Where("next_execution >= ? AND next_execution <= ?", from, to).
		Where("is_active = ?", true).
		Select()
	return schedules, err
}

// UpdateNextExecution updates the next execution time for a schedule
func (r *irrigationScheduleRepository) UpdateNextExecution(ctx context.Context, id string, nextExecution string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("next_execution = ?", nextExecution).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Update()
	return err
}

// UpdateLastExecuted updates the last executed time for a schedule
func (r *irrigationScheduleRepository) UpdateLastExecuted(ctx context.Context, id string, lastExecuted string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("last_executed = ?", lastExecuted).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Update()
	return err
}

// Count returns the total number of irrigation schedules matching the filter
func (r *irrigationScheduleRepository) Count(ctx context.Context, filter *entity.IrrigationScheduleFilter) (int, error) {
	query := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil))

	// Apply same filters as List method
	if filter.GrowingZoneID != "" {
		query = query.Where("growing_zone_id = ?", filter.GrowingZoneID)
	}
	if filter.PlantingCycleID != "" {
		query = query.Where("planting_cycle_id = ?", filter.PlantingCycleID)
	}
	if filter.ScheduleName != "" {
		query = query.Where("schedule_name ILIKE ?", "%"+filter.ScheduleName+"%")
	}
	if filter.IrrigationType != "" {
		query = query.Where("irrigation_type = ?", filter.IrrigationType)
	}
	if filter.Frequency != "" {
		query = query.Where("frequency = ?", filter.Frequency)
	}
	if filter.IsActive {
		query = query.Where("is_active = ?", filter.IsActive)
	}
	if filter.FertilizerMix {
		query = query.Where("fertilizer_mix = ?", filter.FertilizerMix)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}
	if filter.NextExecutionFrom != nil {
		query = query.Where("next_execution >= ?", filter.NextExecutionFrom)
	}
	if filter.NextExecutionTo != nil {
		query = query.Where("next_execution <= ?", filter.NextExecutionTo)
	}

	return query.Count()
}

// CheckScheduleNameExists checks if a schedule name already exists for a specific growing zone
func (r *irrigationScheduleRepository) CheckScheduleNameExists(ctx context.Context, scheduleName, growingZoneID string) (bool, error) {
	count, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Where("schedule_name = ? AND growing_zone_id = ?", scheduleName, growingZoneID).
		Count()
	return count > 0, err
}

// GetSchedulesByDateRange retrieves schedules within a specific date range
func (r *irrigationScheduleRepository) GetSchedulesByDateRange(ctx context.Context, from, to string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).
		Where("created_at >= ? AND created_at <= ?", from, to).
		Select()
	return schedules, err
}

// BulkUpdateStatus updates the status of multiple schedules
func (r *irrigationScheduleRepository) BulkUpdateStatus(ctx context.Context, ids []string, isActive bool) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", time.Now()).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

// GetScheduleStatistics returns statistics about irrigation schedules
func (r *irrigationScheduleRepository) GetScheduleStatistics(ctx context.Context) (*entity.IrrigationScheduleStatistics, error) {
	stats := &entity.IrrigationScheduleStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalSchedules = total

	active, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Where("is_active = ?", true).Count()
	if err != nil {
		return nil, err
	}
	stats.ActiveSchedules = active

	return stats, nil
}
