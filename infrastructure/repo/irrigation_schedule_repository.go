package repo

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
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
func (r *irrigationScheduleRepository) List(ctx context.Context, filter *entity.IrrigationScheduleFilter) ([]*entity.IrrigationSchedule, int64, error) {
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
		return nil, 0, err
	}

	// Apply pagination and sorting
	query = query.Order(filter.SortBy + " " + filter.SortOrder)
	query = query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}
	return schedules, int64(total), nil
}

func (r *irrigationScheduleRepository) GetByGrowingZoneID(ctx context.Context, growingZoneID string, filter *common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("growing_zone_id = ?", growingZoneID)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset((filter.Page - 1) * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetByPlantingCycleID retrieves irrigation schedules by planting cycle ID
func (r *irrigationScheduleRepository) GetByPlantingCycleID(ctx context.Context, plantingCycleID string) ([]*entity.IrrigationSchedule, error) {
	var schedules []*entity.IrrigationSchedule
	err := r.db.ModelContext(ctx, &schedules).Where("planting_cycle_id = ?", plantingCycleID).Select()
	return schedules, err
}

func (r *irrigationScheduleRepository) GetActiveSchedules(ctx context.Context, isActive bool, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("is_active = ?", isActive)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByType retrieves irrigation schedules by irrigation type
func (r *irrigationScheduleRepository) GetSchedulesByType(ctx context.Context, irrigationType string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("irrigation_type = ?", irrigationType)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByFrequency retrieves irrigation schedules by frequency
func (r *irrigationScheduleRepository) GetSchedulesByFrequency(ctx context.Context, frequency string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("frequency = ?", frequency)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesWithFertilizerMix retrieves irrigation schedules that include fertilizer mixing
func (r *irrigationScheduleRepository) GetSchedulesWithFertilizerMix(ctx context.Context, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("fertilizer_mix = ?", true)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByCreator retrieves irrigation schedules created by a specific user
func (r *irrigationScheduleRepository) GetSchedulesByCreator(ctx context.Context, createdBy string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("created_by = ?", createdBy)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

func (r *irrigationScheduleRepository) GetSchedulesForExecution(ctx context.Context, from, to string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("next_execution >= ? AND next_execution <= ?", from, to).
		Where("is_active = ?", true)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

func (r *irrigationScheduleRepository) UpdateNextExecution(ctx context.Context, id string, nextExecution string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("next_execution = ?", nextExecution).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Update()
	return err
}

func (r *irrigationScheduleRepository) UpdateLastExecuted(ctx context.Context, id string, lastExecuted string) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("last_executed = ?", lastExecuted).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Update()
	return err
}

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

func (r *irrigationScheduleRepository) CheckScheduleNameExists(ctx context.Context, scheduleName, growingZoneID string) (bool, error) {
	count, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Where("schedule_name = ? AND growing_zone_id = ?", scheduleName, growingZoneID).
		Count()
	return count > 0, err
}

func (r *irrigationScheduleRepository) GetSchedulesByDateRange(ctx context.Context, from, to string, request common.Pagination) ([]*entity.IrrigationSchedule, int64, error) {
	var schedules []*entity.IrrigationSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("created_at >= ? AND created_at <= ?", from, to)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(request.SortBy + " " + request.SortOrder).
		Limit(request.PageSize).
		Offset((request.Page - 1) * request.PageSize).
		Select()
	return schedules, int64(total), err
}

func (r *irrigationScheduleRepository) BulkUpdateStatus(ctx context.Context, ids []string, isActive bool) error {
	_, err := r.db.ModelContext(ctx, (*entity.IrrigationSchedule)(nil)).
		Set("is_active = ?", isActive).
		Set("updated_at = ?", time.Now()).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

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
