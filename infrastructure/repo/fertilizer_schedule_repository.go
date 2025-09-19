package repo

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/go-pg/pg/v10"
)

type fertilizerScheduleRepository struct {
	db *pg.DB
}

// NewFertilizerScheduleRepository creates a new fertilizer schedule repository
func NewFertilizerScheduleRepository(db *pg.DB) repository.FertilizerScheduleRepository {
	return &fertilizerScheduleRepository{
		db: db,
	}
}

// Create creates a new fertilizer schedule
func (r *fertilizerScheduleRepository) Create(ctx context.Context, req *entity.CreateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error) {
	schedule := &entity.FertilizerSchedule{
		PlantingCycleID:     req.PlantingCycleID,
		FertilizerTypeID:    req.FertilizerTypeID,
		ApplicationDate:     req.ApplicationDate,
		Dosage:              req.Dosage,
		Unit:                req.Unit,
		ApplicationMethod:   req.ApplicationMethod,
		GrowthStage:         req.GrowthStage,
		WeatherConditions:   req.WeatherConditions,
		SoilConditions:      req.SoilConditions,
		IsCompleted:         req.IsCompleted,
		CompletedDate:       req.CompletedDate,
		ActualDosage:        req.ActualDosage,
		EffectivenessRating: req.EffectivenessRating,
		Notes:               req.Notes,
		CreatedBy:           req.CreatedBy,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	_, err := r.db.ModelContext(ctx, schedule).Insert()
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// GetByID retrieves a fertilizer schedule by ID
func (r *fertilizerScheduleRepository) GetByID(ctx context.Context, id string) (*entity.FertilizerSchedule, error) {
	schedule := &entity.FertilizerSchedule{}
	err := r.db.ModelContext(ctx, schedule).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return schedule, nil
}

// Update updates an existing fertilizer schedule
func (r *fertilizerScheduleRepository) Update(ctx context.Context, req *entity.UpdateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error) {
	schedule := &entity.FertilizerSchedule{
		ID:                  req.ID,
		PlantingCycleID:     req.PlantingCycleID,
		FertilizerTypeID:    req.FertilizerTypeID,
		ApplicationDate:     req.ApplicationDate,
		Dosage:              req.Dosage,
		Unit:                req.Unit,
		ApplicationMethod:   req.ApplicationMethod,
		GrowthStage:         req.GrowthStage,
		WeatherConditions:   req.WeatherConditions,
		SoilConditions:      req.SoilConditions,
		IsCompleted:         req.IsCompleted,
		CompletedDate:       req.CompletedDate,
		ActualDosage:        req.ActualDosage,
		EffectivenessRating: req.EffectivenessRating,
		Notes:               req.Notes,
		UpdatedAt:           time.Now(),
	}

	_, err := r.db.ModelContext(ctx, schedule).Where("id = ?", req.ID).Update()
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// Delete removes a fertilizer schedule by ID
func (r *fertilizerScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).Where("id = ?", id).Delete()
	return err
}

// List retrieves fertilizer schedules with filtering and pagination
func (r *fertilizerScheduleRepository) List(ctx context.Context, filter *entity.FertilizerScheduleFilter) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	query := r.db.ModelContext(ctx, &schedules)

	// Apply filters
	if filter.PlantingCycleID != "" {
		query = query.Where("planting_cycle_id = ?", filter.PlantingCycleID)
	}
	if filter.FertilizerTypeID != "" {
		query = query.Where("fertilizer_type_id = ?", filter.FertilizerTypeID)
	}
	if filter.ApplicationMethod != "" {
		query = query.Where("application_method = ?", filter.ApplicationMethod)
	}
	if filter.GrowthStage != "" {
		query = query.Where("growth_stage = ?", filter.GrowthStage)
	}
	if filter.IsCompleted {
		query = query.Where("is_completed = ?", filter.IsCompleted)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.ApplicationDateFrom != nil {
		query = query.Where("application_date >= ?", filter.ApplicationDateFrom)
	}
	if filter.ApplicationDateTo != nil {
		query = query.Where("application_date <= ?", filter.ApplicationDateTo)
	}
	if filter.CompletedDateFrom != nil {
		query = query.Where("completed_date >= ?", filter.CompletedDateFrom)
	}
	if filter.CompletedDateTo != nil {
		query = query.Where("completed_date <= ?", filter.CompletedDateTo)
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
		return nil, 0, err
	}

	// Apply pagination and sorting
	query = query.Order(filter.SortBy + " " + filter.SortOrder)
	query = query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	// Convert to response format
	var responses []entity.FertilizerScheduleResponse
	for _, schedule := range schedules {
		responses = append(responses, entity.FertilizerScheduleResponse{
			ID:                  schedule.ID,
			PlantingCycleID:     schedule.PlantingCycleID,
			FertilizerTypeID:    schedule.FertilizerTypeID,
			ApplicationDate:     schedule.ApplicationDate,
			Dosage:              schedule.Dosage,
			Unit:                schedule.Unit,
			ApplicationMethod:   schedule.ApplicationMethod,
			GrowthStage:         schedule.GrowthStage,
			WeatherConditions:   schedule.WeatherConditions,
			SoilConditions:      schedule.SoilConditions,
			IsCompleted:         schedule.IsCompleted,
			CompletedDate:       schedule.CompletedDate,
			ActualDosage:        schedule.ActualDosage,
			EffectivenessRating: schedule.EffectivenessRating,
			Notes:               schedule.Notes,
			CreatedBy:           schedule.CreatedBy,
			CreatedAt:           schedule.CreatedAt,
			UpdatedAt:           schedule.UpdatedAt,
		})
	}

	return schedules, int64(total), nil
}

// GetByPlantingCycleID retrieves fertilizer schedules by planting cycle ID
func (r *fertilizerScheduleRepository) GetByPlantingCycleID(ctx context.Context, plantingCycleID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("planting_cycle_id = ?", plantingCycleID)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetByFertilizerTypeID retrieves fertilizer schedules by fertilizer type ID
func (r *fertilizerScheduleRepository) GetByFertilizerTypeID(ctx context.Context, fertilizerTypeID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("fertilizer_type_id = ?", fertilizerTypeID)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetByApplicationMethod retrieves fertilizer schedules by application method
func (r *fertilizerScheduleRepository) GetByApplicationMethod(ctx context.Context, method string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("application_method = ?", method)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetByGrowthStage retrieves fertilizer schedules by growth stage
func (r *fertilizerScheduleRepository) GetByGrowthStage(ctx context.Context, growthStage string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("growth_stage = ?", growthStage)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetByCreator retrieves fertilizer schedules created by a specific user
func (r *fertilizerScheduleRepository) GetByCreator(ctx context.Context, createdBy string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("created_by = ?", createdBy)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetCompletedSchedules retrieves all completed fertilizer schedules
func (r *fertilizerScheduleRepository) GetCompletedSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("is_completed = ?", true)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy+" "+filter.SortOrder).
		Where("is_completed = ?", true).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetPendingSchedules retrieves all pending fertilizer schedules
func (r *fertilizerScheduleRepository) GetPendingSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("is_completed = ?", false)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByDateRange retrieves schedules within a specific date range
func (r *fertilizerScheduleRepository) GetSchedulesByDateRange(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("created_at >= ? AND created_at <= ?", from, to)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByApplicationDate retrieves schedules by application date range
func (r *fertilizerScheduleRepository) GetSchedulesByApplicationDate(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("application_date >= ? AND application_date <= ?", from, to)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByCompletionDate retrieves schedules by completion date range
func (r *fertilizerScheduleRepository) GetSchedulesByCompletionDate(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("completed_date >= ? AND completed_date <= ?", from, to)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByDosageRange retrieves schedules within a dosage range
func (r *fertilizerScheduleRepository) GetSchedulesByDosageRange(ctx context.Context, minDosage, maxDosage float64, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("dosage >= ? AND dosage <= ?", minDosage, maxDosage)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByEffectivenessRating retrieves schedules by effectiveness rating
func (r *fertilizerScheduleRepository) GetSchedulesByEffectivenessRating(ctx context.Context, rating int, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("effectiveness_rating = ?", rating)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByWeatherConditions retrieves schedules by weather conditions
func (r *fertilizerScheduleRepository) GetSchedulesByWeatherConditions(ctx context.Context, conditions string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("weather_conditions ILIKE ?", "%"+conditions+"%")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesBySoilConditions retrieves schedules by soil conditions
func (r *fertilizerScheduleRepository) GetSchedulesBySoilConditions(ctx context.Context, conditions string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("soil_conditions ILIKE ?", "%"+conditions+"%")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetUpcomingSchedules retrieves schedules that need to be executed soon
func (r *fertilizerScheduleRepository) GetUpcomingSchedules(ctx context.Context, days int, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	threshold := time.Now().AddDate(0, 0, days)
	q := r.db.ModelContext(ctx, &schedules).
		Where("application_date <= ? AND application_date > ?", threshold, time.Now()).
		Where("is_completed = ?", false)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetOverdueSchedules retrieves schedules that are overdue
func (r *fertilizerScheduleRepository) GetOverdueSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("application_date < ?", time.Now()).
		Where("is_completed = ?", false)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetSchedulesByUnit retrieves schedules by unit
func (r *fertilizerScheduleRepository) GetSchedulesByUnit(ctx context.Context, unit string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).Where("unit = ?", unit)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// Count returns the total number of fertilizer schedules matching the filter
func (r *fertilizerScheduleRepository) Count(ctx context.Context, filter *entity.FertilizerScheduleFilter) (int, error) {
	query := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil))

	// Apply same filters as List method
	if filter.PlantingCycleID != "" {
		query = query.Where("planting_cycle_id = ?", filter.PlantingCycleID)
	}
	if filter.FertilizerTypeID != "" {
		query = query.Where("fertilizer_type_id = ?", filter.FertilizerTypeID)
	}
	if filter.ApplicationMethod != "" {
		query = query.Where("application_method = ?", filter.ApplicationMethod)
	}
	if filter.GrowthStage != "" {
		query = query.Where("growth_stage = ?", filter.GrowthStage)
	}
	if filter.IsCompleted {
		query = query.Where("is_completed = ?", filter.IsCompleted)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.ApplicationDateFrom != nil {
		query = query.Where("application_date >= ?", filter.ApplicationDateFrom)
	}
	if filter.ApplicationDateTo != nil {
		query = query.Where("application_date <= ?", filter.ApplicationDateTo)
	}
	if filter.CompletedDateFrom != nil {
		query = query.Where("completed_date >= ?", filter.CompletedDateFrom)
	}
	if filter.CompletedDateTo != nil {
		query = query.Where("completed_date <= ?", filter.CompletedDateTo)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}

	return query.Count()
}

// GetScheduleStatistics returns statistics about fertilizer schedules
func (r *fertilizerScheduleRepository) GetScheduleStatistics(ctx context.Context) (*entity.FertilizerScheduleStatistics, error) {
	stats := &entity.FertilizerScheduleStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalSchedules = total

	completed, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("is_completed = ?", true).Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedSchedules = completed

	pending, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("is_completed = ?", false).Count()
	if err != nil {
		return nil, err
	}
	stats.PendingSchedules = pending

	return stats, nil
}

// GetPlantingCycleScheduleStatistics returns statistics for schedules of a specific planting cycle
func (r *fertilizerScheduleRepository) GetPlantingCycleScheduleStatistics(ctx context.Context, plantingCycleID string) (*entity.FertilizerScheduleStatistics, error) {
	stats := &entity.FertilizerScheduleStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("planting_cycle_id = ?", plantingCycleID).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalSchedules = total

	completed, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("planting_cycle_id = ? AND is_completed = ?", plantingCycleID, true).Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedSchedules = completed

	pending, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("planting_cycle_id = ? AND is_completed = ?", plantingCycleID, false).Count()
	if err != nil {
		return nil, err
	}
	stats.PendingSchedules = pending

	return stats, nil
}

// GetFertilizerTypeScheduleStatistics returns statistics for schedules of a specific fertilizer type
func (r *fertilizerScheduleRepository) GetFertilizerTypeScheduleStatistics(ctx context.Context, fertilizerTypeID string) (*entity.FertilizerScheduleStatistics, error) {
	stats := &entity.FertilizerScheduleStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("fertilizer_type_id = ?", fertilizerTypeID).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalSchedules = total

	completed, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("fertilizer_type_id = ? AND is_completed = ?", fertilizerTypeID, true).Count()
	if err != nil {
		return nil, err
	}
	stats.CompletedSchedules = completed

	pending, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("fertilizer_type_id = ? AND is_completed = ?", fertilizerTypeID, false).Count()
	if err != nil {
		return nil, err
	}
	stats.PendingSchedules = pending

	return stats, nil
}

// GetEffectivenessReport returns effectiveness report for fertilizer schedules
func (r *fertilizerScheduleRepository) GetEffectivenessReport(ctx context.Context, from, to string) (*entity.FertilizerEffectivenessReport, error) {
	report := &entity.FertilizerEffectivenessReport{}

	// Get total schedules in date range
	total, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Count()
	if err != nil {
		return nil, err
	}
	report.TotalSchedules = total

	// Get average effectiveness rating
	var avgRating float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("effectiveness_rating > 0").
		ColumnExpr("AVG(effectiveness_rating)").
		Select(&avgRating)
	if err != nil {
		return nil, err
	}
	report.AverageEffectiveness = avgRating

	return report, nil
}

// GetDosageReport returns dosage report for fertilizer schedules
func (r *fertilizerScheduleRepository) GetDosageReport(ctx context.Context, from, to string) (*entity.FertilizerDosageReport, error) {
	report := &entity.FertilizerDosageReport{}

	// Get total dosage used
	var totalDosage float64
	err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("is_completed = ?", true).
		ColumnExpr("SUM(actual_dosage)").
		Select(&totalDosage)
	if err != nil {
		return nil, err
	}
	report.TotalDosageUsed = totalDosage

	// Get average dosage
	var avgDosage float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("is_completed = ?", true).
		ColumnExpr("AVG(actual_dosage)").
		Select(&avgDosage)
	if err != nil {
		return nil, err
	}
	report.AverageDosageUsed = avgDosage

	return report, nil
}

// GetGrowthStageReport returns growth stage report for fertilizer schedules
func (r *fertilizerScheduleRepository) GetGrowthStageReport(ctx context.Context, from, to string) (*entity.FertilizerGrowthStageReport, error) {
	report := &entity.FertilizerGrowthStageReport{}

	// Get schedules by growth stage
	seedlingCount, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("growth_stage = ?", "seedling").
		Count()
	if err != nil {
		return nil, err
	}
	report.ByGrowthStage["seedling"] = seedlingCount

	vegetativeCount, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("growth_stage = ?", "vegetative").
		Count()
	if err != nil {
		return nil, err
	}
	report.ByGrowthStage["vegetative"] = vegetativeCount

	floweringCount, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("growth_stage = ?", "flowering").
		Count()
	if err != nil {
		return nil, err
	}
	report.ByGrowthStage["flowering"] = floweringCount

	fruitingCount, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Where("application_date >= ? AND application_date <= ?", from, to).
		Where("growth_stage = ?", "fruiting").
		Count()
	if err != nil {
		return nil, err
	}
	report.ByGrowthStage["fruiting"] = fruitingCount

	return report, nil
}

// BulkUpdateStatus updates the completion status of multiple schedules
func (r *fertilizerScheduleRepository) BulkUpdateStatus(ctx context.Context, ids []string, isCompleted bool) error {
	_, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Set("is_completed = ?", isCompleted).
		Set("updated_at = ?", time.Now()).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

// BulkUpdateEffectivenessRating updates the effectiveness rating of multiple schedules
func (r *fertilizerScheduleRepository) BulkUpdateEffectivenessRating(ctx context.Context, ids []string, rating int) error {
	_, err := r.db.ModelContext(ctx, (*entity.FertilizerSchedule)(nil)).
		Set("effectiveness_rating = ?", rating).
		Set("updated_at = ?", time.Now()).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

// GetRecentSchedules retrieves recently created fertilizer schedules
func (r *fertilizerScheduleRepository) GetRecentSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Order("created_at DESC")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Order("created_at DESC").
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetScheduleTimeline retrieves a timeline of fertilizer schedules for a planting cycle
func (r *fertilizerScheduleRepository) GetScheduleTimeline(ctx context.Context, plantingCycleID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("planting_cycle_id = ?", plantingCycleID).
		Order("application_date ASC")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return schedules, int64(total), err
}

// GetScheduleRecommendations returns fertilizer schedule recommendations based on growth stage
func (r *fertilizerScheduleRepository) GetScheduleRecommendations(ctx context.Context, growthStage string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error) {
	var schedules []*entity.FertilizerSchedule
	q := r.db.ModelContext(ctx, &schedules).
		Where("growth_stage = ?", growthStage).
		Where("is_completed = ?", true)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Order(filter.SortBy + " " + filter.SortOrder).
		Select()
	return schedules, int64(total), err
}
