package repository

import (
	"context"
	"fertigation-Service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

// FertilizerTypeRepository defines the interface for fertilizer type data operations
type FertilizerTypeRepository interface {
	// Create creates a new fertilizer type
	Create(ctx context.Context, req *entity.CreateFertilizerTypeRequest) (*entity.FertilizerType, error)

	// GetByID retrieves a fertilizer type by ID
	GetByID(ctx context.Context, id string) (*entity.FertilizerType, error)

	// Update updates an existing fertilizer type
	Update(ctx context.Context, req *entity.UpdateFertilizerTypeRequest) (*entity.FertilizerType, error)

	// Delete removes a fertilizer type by ID
	Delete(ctx context.Context, id string) error

	// List retrieves fertilizer types with filtering and pagination
	List(ctx context.Context, filter *entity.FertilizerTypeFilter) ([]*entity.FertilizerType, int64, error)

	// GetByName retrieves fertilizer types by name (partial match)
	GetByName(ctx context.Context, name string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByType retrieves fertilizer types by type
	GetByType(ctx context.Context, fertilizerType string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByApplicationMethod retrieves fertilizer types by application method
	GetByApplicationMethod(ctx context.Context, method string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByStatus retrieves fertilizer types by status
	GetByStatus(ctx context.Context, status string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByManufacturer retrieves fertilizer types by manufacturer
	GetByManufacturer(ctx context.Context, manufacturer string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByCreator retrieves fertilizer types created by a specific user
	GetByCreator(ctx context.Context, createdBy string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByNPKRatio retrieves fertilizer types by NPK ratio
	GetByNPKRatio(ctx context.Context, npkRatio string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByExpiryDate retrieves fertilizer types by expiry date range
	GetByExpiryDate(ctx context.Context, from, to string, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetExpiredFertilizers retrieves all expired fertilizer types
	GetExpiredFertilizers(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetExpiringSoon retrieves fertilizer types expiring within specified days
	GetExpiringSoon(ctx context.Context, days int, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByCostRange retrieves fertilizer types within a cost range
	GetByCostRange(ctx context.Context, minCost, maxCost float64, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByNitrogenRange retrieves fertilizer types within a nitrogen percentage range
	GetByNitrogenRange(ctx context.Context, minNitrogen, maxNitrogen float64, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByPhosphorusRange retrieves fertilizer types within a phosphorus percentage range
	GetByPhosphorusRange(ctx context.Context, minPhosphorus, maxPhosphorus float64, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByPotassiumRange retrieves fertilizer types within a potassium percentage range
	GetByPotassiumRange(ctx context.Context, minPotassium, maxPotassium float64, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// GetByDosageRange retrieves fertilizer types within a dosage range
	GetByDosageRange(ctx context.Context, minDosage, maxDosage float64, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// Count returns the total number of fertilizer types matching the filter
	Count(ctx context.Context, filter *entity.FertilizerTypeFilter) (int, error)

	// CheckNameExists checks if a fertilizer type name already exists
	CheckNameExists(ctx context.Context, name string) (bool, error)

	// CheckBatchNumberExists checks if a batch number already exists
	CheckBatchNumberExists(ctx context.Context, batchNumber string) (bool, error)

	// GetFertilizerTypeStatistics returns statistics about fertilizer types
	GetFertilizerTypeStatistics(ctx context.Context) (*entity.FertilizerTypeStatistics, error)

	// GetExpiryReport returns expiry report for fertilizer types
	GetExpiryReport(ctx context.Context) (*entity.FertilizerExpiryReport, error)

	// GetCostAnalysis returns cost analysis for fertilizer types
	GetCostAnalysis(ctx context.Context) (*entity.FertilizerCostAnalysis, error)

	// GetNPKAnalysis returns NPK analysis for fertilizer types
	GetNPKAnalysis(ctx context.Context) (*entity.FertilizerNPKAnalysis, error)

	// BulkUpdateStatus updates the status of multiple fertilizer types
	BulkUpdateStatus(ctx context.Context, ids []string, status string) error

	// GetRecentFertilizerTypes retrieves recently created fertilizer types
	GetRecentFertilizerTypes(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerType, int64, error)

	// SearchFertilizerTypes performs full-text search on fertilizer types
	SearchFertilizerTypes(ctx context.Context, filter common.Pagination) ([]*entity.FertilizerType, int64, error)
}
