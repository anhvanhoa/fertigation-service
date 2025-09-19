package repository

import (
	"context"
	"fertigation-Service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

// FertilizerScheduleRepository defines the interface for fertilizer schedule data operations
type FertilizerScheduleRepository interface {
	// Create creates a new fertilizer schedule
	Create(ctx context.Context, req *entity.CreateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error)

	// GetByID retrieves a fertilizer schedule by ID
	GetByID(ctx context.Context, id string) (*entity.FertilizerSchedule, error)

	// Update updates an existing fertilizer schedule
	Update(ctx context.Context, req *entity.UpdateFertilizerScheduleRequest) (*entity.FertilizerSchedule, error)

	// Delete removes a fertilizer schedule by ID
	Delete(ctx context.Context, id string) error

	// List retrieves fertilizer schedules with filtering and pagination
	List(ctx context.Context, filter *entity.FertilizerScheduleFilter) ([]*entity.FertilizerSchedule, int64, error)

	// GetByPlantingCycleID retrieves fertilizer schedules by planting cycle ID
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetByFertilizerTypeID retrieves fertilizer schedules by fertilizer type ID
	GetByFertilizerTypeID(ctx context.Context, fertilizerTypeID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetByApplicationMethod retrieves fertilizer schedules by application method
	GetByApplicationMethod(ctx context.Context, method string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetByGrowthStage retrieves fertilizer schedules by growth stage
	GetByGrowthStage(ctx context.Context, growthStage string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetByCreator retrieves fertilizer schedules created by a specific user
	GetByCreator(ctx context.Context, createdBy string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetCompletedSchedules retrieves all completed fertilizer schedules
	GetCompletedSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetPendingSchedules retrieves all pending fertilizer schedules
	GetPendingSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByDateRange retrieves schedules within a specific date range
	GetSchedulesByDateRange(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByApplicationDate retrieves schedules by application date range
	GetSchedulesByApplicationDate(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByCompletionDate retrieves schedules by completion date range
	GetSchedulesByCompletionDate(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByDosageRange retrieves schedules within a dosage range
	GetSchedulesByDosageRange(ctx context.Context, minDosage, maxDosage float64, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByEffectivenessRating retrieves schedules by effectiveness rating
	GetSchedulesByEffectivenessRating(ctx context.Context, rating int, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByWeatherConditions retrieves schedules by weather conditions
	GetSchedulesByWeatherConditions(ctx context.Context, conditions string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesBySoilConditions retrieves schedules by soil conditions
	GetSchedulesBySoilConditions(ctx context.Context, conditions string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetUpcomingSchedules retrieves schedules that need to be executed soon
	GetUpcomingSchedules(ctx context.Context, days int, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetOverdueSchedules retrieves schedules that are overdue
	GetOverdueSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetSchedulesByUnit retrieves schedules by unit
	GetSchedulesByUnit(ctx context.Context, unit string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// Count returns the total number of fertilizer schedules matching the filter
	Count(ctx context.Context, filter *entity.FertilizerScheduleFilter) (int, error)

	// GetScheduleStatistics returns statistics about fertilizer schedules
	GetScheduleStatistics(ctx context.Context) (*entity.FertilizerScheduleStatistics, error)

	// GetPlantingCycleScheduleStatistics returns statistics for schedules of a specific planting cycle
	GetPlantingCycleScheduleStatistics(ctx context.Context, plantingCycleID string) (*entity.FertilizerScheduleStatistics, error)

	// GetFertilizerTypeScheduleStatistics returns statistics for schedules of a specific fertilizer type
	GetFertilizerTypeScheduleStatistics(ctx context.Context, fertilizerTypeID string) (*entity.FertilizerScheduleStatistics, error)

	// GetEffectivenessReport returns effectiveness report for fertilizer schedules
	GetEffectivenessReport(ctx context.Context, from, to string) (*entity.FertilizerEffectivenessReport, error)

	// GetDosageReport returns dosage report for fertilizer schedules
	GetDosageReport(ctx context.Context, from, to string) (*entity.FertilizerDosageReport, error)

	// GetGrowthStageReport returns growth stage report for fertilizer schedules
	GetGrowthStageReport(ctx context.Context, from, to string) (*entity.FertilizerGrowthStageReport, error)

	// BulkUpdateStatus updates the completion status of multiple schedules
	BulkUpdateStatus(ctx context.Context, ids []string, isCompleted bool) error

	// BulkUpdateEffectivenessRating updates the effectiveness rating of multiple schedules
	BulkUpdateEffectivenessRating(ctx context.Context, ids []string, rating int) error

	// GetRecentSchedules retrieves recently created fertilizer schedules
	GetRecentSchedules(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetScheduleTimeline retrieves a timeline of fertilizer schedules for a planting cycle
	GetScheduleTimeline(ctx context.Context, plantingCycleID string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)

	// GetScheduleRecommendations returns fertilizer schedule recommendations based on growth stage
	GetScheduleRecommendations(ctx context.Context, growthStage string, filter common.Pagination) ([]*entity.FertilizerSchedule, int64, error)
}
