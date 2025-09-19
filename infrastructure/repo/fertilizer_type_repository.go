package repo

import (
	"context"
	"fertigation-Service/domain/entity"
	"fertigation-Service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/common"
	"github.com/go-pg/pg/v10"
)

type fertilizerTypeRepository struct {
	db *pg.DB
}

// NewFertilizerTypeRepository creates a new fertilizer type repository
func NewFertilizerTypeRepository(db *pg.DB) repository.FertilizerTypeRepository {
	return &fertilizerTypeRepository{
		db: db,
	}
}

// Create creates a new fertilizer type
func (r *fertilizerTypeRepository) Create(ctx context.Context, req *entity.CreateFertilizerTypeRequest) (*entity.FertilizerType, error) {
	fertilizerType := &entity.FertilizerType{
		Name:                 req.Name,
		Type:                 req.Type,
		NPKRatio:             req.NPKRatio,
		NitrogenPercentage:   req.NitrogenPercentage,
		PhosphorusPercentage: req.PhosphorusPercentage,
		PotassiumPercentage:  req.PotassiumPercentage,
		TraceElements:        req.TraceElements,
		ApplicationMethod:    req.ApplicationMethod,
		DosagePerPlant:       req.DosagePerPlant,
		DosagePerM2:          req.DosagePerM2,
		Unit:                 req.Unit,
		Manufacturer:         req.Manufacturer,
		BatchNumber:          req.BatchNumber,
		ExpiryDate:           req.ExpiryDate,
		CostPerUnit:          req.CostPerUnit,
		Description:          req.Description,
		SafetyNotes:          req.SafetyNotes,
		Status:               req.Status,
		CreatedBy:            req.CreatedBy,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	_, err := r.db.ModelContext(ctx, fertilizerType).Insert()
	if err != nil {
		return nil, err
	}

	return fertilizerType, nil
}

// GetByID retrieves a fertilizer type by ID
func (r *fertilizerTypeRepository) GetByID(ctx context.Context, id string) (*entity.FertilizerType, error) {
	fertilizerType := &entity.FertilizerType{}
	err := r.db.ModelContext(ctx, fertilizerType).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return fertilizerType, nil
}

// Update updates an existing fertilizer type
func (r *fertilizerTypeRepository) Update(ctx context.Context, req *entity.UpdateFertilizerTypeRequest) (*entity.FertilizerType, error) {
	fertilizerType := &entity.FertilizerType{
		ID:                   req.ID,
		Name:                 req.Name,
		Type:                 req.Type,
		NPKRatio:             req.NPKRatio,
		NitrogenPercentage:   req.NitrogenPercentage,
		PhosphorusPercentage: req.PhosphorusPercentage,
		PotassiumPercentage:  req.PotassiumPercentage,
		TraceElements:        req.TraceElements,
		ApplicationMethod:    req.ApplicationMethod,
		DosagePerPlant:       req.DosagePerPlant,
		DosagePerM2:          req.DosagePerM2,
		Unit:                 req.Unit,
		Manufacturer:         req.Manufacturer,
		BatchNumber:          req.BatchNumber,
		ExpiryDate:           req.ExpiryDate,
		CostPerUnit:          req.CostPerUnit,
		Description:          req.Description,
		SafetyNotes:          req.SafetyNotes,
		Status:               req.Status,
		UpdatedAt:            time.Now(),
	}

	_, err := r.db.ModelContext(ctx, fertilizerType).Where("id = ?", req.ID).Update()
	if err != nil {
		return nil, err
	}

	return fertilizerType, nil
}

// Delete removes a fertilizer type by ID
func (r *fertilizerTypeRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).Where("id = ?", id).Delete()
	return err
}

// List retrieves fertilizer types with filtering and pagination
func (r *fertilizerTypeRepository) List(ctx context.Context, filter *entity.FertilizerTypeFilter) ([]*entity.FertilizerType, int64, error) {
	var fertilizerTypes []*entity.FertilizerType
	query := r.db.ModelContext(ctx, &fertilizerTypes)

	// Apply filters
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.ApplicationMethod != "" {
		query = query.Where("application_method = ?", filter.ApplicationMethod)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Manufacturer != "" {
		query = query.Where("manufacturer ILIKE ?", "%"+filter.Manufacturer+"%")
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.ExpiryDateFrom != nil {
		query = query.Where("expiry_date >= ?", filter.ExpiryDateFrom)
	}
	if filter.ExpiryDateTo != nil {
		query = query.Where("expiry_date <= ?", filter.ExpiryDateTo)
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
	query = query.Limit(filter.PageSize).Offset((filter.Page - 1) * filter.PageSize)

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	// Convert to response format
	var responses []*entity.FertilizerType
	for _, ft := range fertilizerTypes {
		responses = append(responses, &entity.FertilizerType{
			ID:                   ft.ID,
			Name:                 ft.Name,
			Type:                 ft.Type,
			NPKRatio:             ft.NPKRatio,
			NitrogenPercentage:   ft.NitrogenPercentage,
			PhosphorusPercentage: ft.PhosphorusPercentage,
			PotassiumPercentage:  ft.PotassiumPercentage,
			TraceElements:        ft.TraceElements,
			ApplicationMethod:    ft.ApplicationMethod,
			DosagePerPlant:       ft.DosagePerPlant,
			DosagePerM2:          ft.DosagePerM2,
			Unit:                 ft.Unit,
			Manufacturer:         ft.Manufacturer,
			BatchNumber:          ft.BatchNumber,
			ExpiryDate:           ft.ExpiryDate,
			CostPerUnit:          ft.CostPerUnit,
			Description:          ft.Description,
			SafetyNotes:          ft.SafetyNotes,
			Status:               ft.Status,
			CreatedBy:            ft.CreatedBy,
			CreatedAt:            ft.CreatedAt,
			UpdatedAt:            ft.UpdatedAt,
		})
	}

	return responses, int64(total), nil
}

// GetByName retrieves fertilizer types by name (partial match)
func (r *fertilizerTypeRepository) GetByName(ctx context.Context, name string, filter common.Pagination) ([]*entity.FertilizerType, int64, error) {
	var fertilizerTypes []*entity.FertilizerType
	q := r.db.ModelContext(ctx, &fertilizerTypes).Where("name ILIKE ?", "%"+name+"%")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return fertilizerTypes, int64(total), err
}

// GetByType retrieves fertilizer types by type
func (r *fertilizerTypeRepository) GetByType(ctx context.Context, fertilizerType string, filter common.Pagination) ([]*entity.FertilizerType, int64, error) {
	var fertilizerTypes []*entity.FertilizerType
	q := r.db.ModelContext(ctx, &fertilizerTypes).Where("type = ?", fertilizerType)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.Order(filter.SortBy + " " + filter.SortOrder).
		Limit(filter.PageSize).
		Offset(filter.Page * filter.PageSize).
		Select()
	return fertilizerTypes, int64(total), err
}

// GetByApplicationMethod retrieves fertilizer types by application method
func (r *fertilizerTypeRepository) GetByApplicationMethod(ctx context.Context, method string, filter common.Pagination) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).Where("application_method = ?", method).Select()
	return fertilizerTypes, err
}

// GetByStatus retrieves fertilizer types by status
func (r *fertilizerTypeRepository) GetByStatus(ctx context.Context, status string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).Where("status = ?", status).Select()
	return fertilizerTypes, err
}

// GetByManufacturer retrieves fertilizer types by manufacturer
func (r *fertilizerTypeRepository) GetByManufacturer(ctx context.Context, manufacturer string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).Where("manufacturer ILIKE ?", "%"+manufacturer+"%").Select()
	return fertilizerTypes, err
}

// GetByCreator retrieves fertilizer types created by a specific user
func (r *fertilizerTypeRepository) GetByCreator(ctx context.Context, createdBy string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).Where("created_by = ?", createdBy).Select()
	return fertilizerTypes, err
}

// GetByNPKRatio retrieves fertilizer types by NPK ratio
func (r *fertilizerTypeRepository) GetByNPKRatio(ctx context.Context, npkRatio string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).Where("npk_ratio = ?", npkRatio).Select()
	return fertilizerTypes, err
}

// GetByExpiryDate retrieves fertilizer types by expiry date range
func (r *fertilizerTypeRepository) GetByExpiryDate(ctx context.Context, from, to string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("expiry_date >= ? AND expiry_date <= ?", from, to).
		Select()
	return fertilizerTypes, err
}

// GetExpiredFertilizers retrieves all expired fertilizer types
func (r *fertilizerTypeRepository) GetExpiredFertilizers(ctx context.Context) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("expiry_date < ?", time.Now()).
		Select()
	return fertilizerTypes, err
}

// GetExpiringSoon retrieves fertilizer types expiring within specified days
func (r *fertilizerTypeRepository) GetExpiringSoon(ctx context.Context, days int) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	expiryThreshold := time.Now().AddDate(0, 0, days)
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("expiry_date <= ? AND expiry_date > ?", expiryThreshold, time.Now()).
		Select()
	return fertilizerTypes, err
}

// GetByCostRange retrieves fertilizer types within a cost range
func (r *fertilizerTypeRepository) GetByCostRange(ctx context.Context, minCost, maxCost float64) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("cost_per_unit >= ? AND cost_per_unit <= ?", minCost, maxCost).
		Select()
	return fertilizerTypes, err
}

// GetByNitrogenRange retrieves fertilizer types within a nitrogen percentage range
func (r *fertilizerTypeRepository) GetByNitrogenRange(ctx context.Context, minNitrogen, maxNitrogen float64) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("nitrogen_percentage >= ? AND nitrogen_percentage <= ?", minNitrogen, maxNitrogen).
		Select()
	return fertilizerTypes, err
}

// GetByPhosphorusRange retrieves fertilizer types within a phosphorus percentage range
func (r *fertilizerTypeRepository) GetByPhosphorusRange(ctx context.Context, minPhosphorus, maxPhosphorus float64) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("phosphorus_percentage >= ? AND phosphorus_percentage <= ?", minPhosphorus, maxPhosphorus).
		Select()
	return fertilizerTypes, err
}

// GetByPotassiumRange retrieves fertilizer types within a potassium percentage range
func (r *fertilizerTypeRepository) GetByPotassiumRange(ctx context.Context, minPotassium, maxPotassium float64) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("potassium_percentage >= ? AND potassium_percentage <= ?", minPotassium, maxPotassium).
		Select()
	return fertilizerTypes, err
}

// GetByDosageRange retrieves fertilizer types within a dosage range
func (r *fertilizerTypeRepository) GetByDosageRange(ctx context.Context, minDosage, maxDosage float64) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("dosage_per_plant >= ? AND dosage_per_plant <= ?", minDosage, maxDosage).
		Select()
	return fertilizerTypes, err
}

// Count returns the total number of fertilizer types matching the filter
func (r *fertilizerTypeRepository) Count(ctx context.Context, filter *entity.FertilizerTypeFilter) (int, error) {
	query := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil))

	// Apply same filters as List method
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.ApplicationMethod != "" {
		query = query.Where("application_method = ?", filter.ApplicationMethod)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Manufacturer != "" {
		query = query.Where("manufacturer ILIKE ?", "%"+filter.Manufacturer+"%")
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.ExpiryDateFrom != nil {
		query = query.Where("expiry_date >= ?", filter.ExpiryDateFrom)
	}
	if filter.ExpiryDateTo != nil {
		query = query.Where("expiry_date <= ?", filter.ExpiryDateTo)
	}
	if filter.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedAtFrom)
	}
	if filter.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedAtTo)
	}

	return query.Count()
}

// CheckNameExists checks if a fertilizer type name already exists
func (r *fertilizerTypeRepository) CheckNameExists(ctx context.Context, name string) (bool, error) {
	count, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("name = ?", name).
		Count()
	return count > 0, err
}

// CheckBatchNumberExists checks if a batch number already exists
func (r *fertilizerTypeRepository) CheckBatchNumberExists(ctx context.Context, batchNumber string) (bool, error) {
	count, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("batch_number = ?", batchNumber).
		Count()
	return count > 0, err
}

// GetFertilizerTypeStatistics returns statistics about fertilizer types
func (r *fertilizerTypeRepository) GetFertilizerTypeStatistics(ctx context.Context) (*entity.FertilizerTypeStatistics, error) {
	stats := &entity.FertilizerTypeStatistics{}

	total, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalFertilizerTypes = total

	active, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("status = ?", "active").Count()
	if err != nil {
		return nil, err
	}
	stats.ActiveFertilizerTypes = active

	expired, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("expiry_date < ?", time.Now()).Count()
	if err != nil {
		return nil, err
	}
	stats.ExpiredFertilizerTypes = expired

	return stats, nil
}

// GetExpiryReport returns expiry report for fertilizer types
func (r *fertilizerTypeRepository) GetExpiryReport(ctx context.Context) (*entity.FertilizerExpiryReport, error) {
	report := &entity.FertilizerExpiryReport{}

	expired, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("expiry_date < ?", time.Now()).Count()
	if err != nil {
		return nil, err
	}
	report.ExpiredCount = expired

	expiringSoon, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Where("expiry_date <= ? AND expiry_date > ?", time.Now().AddDate(0, 0, 30), time.Now()).Count()
	if err != nil {
		return nil, err
	}
	report.ExpiringIn30Days = expiringSoon

	return report, nil
}

// GetCostAnalysis returns cost analysis for fertilizer types
func (r *fertilizerTypeRepository) GetCostAnalysis(ctx context.Context) (*entity.FertilizerCostAnalysis, error) {
	analysis := &entity.FertilizerCostAnalysis{}

	var avgCost float64
	err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("AVG(cost_per_unit)").
		Select(&avgCost)
	if err != nil {
		return nil, err
	}
	analysis.AverageCost = avgCost

	var minCost float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("MIN(cost_per_unit)").
		Select(&minCost)
	if err != nil {
		return nil, err
	}
	analysis.MinCost = minCost

	var maxCost float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("MAX(cost_per_unit)").
		Select(&maxCost)
	if err != nil {
		return nil, err
	}
	analysis.MaxCost = maxCost

	return analysis, nil
}

// GetNPKAnalysis returns NPK analysis for fertilizer types
func (r *fertilizerTypeRepository) GetNPKAnalysis(ctx context.Context) (*entity.FertilizerNPKAnalysis, error) {
	analysis := &entity.FertilizerNPKAnalysis{}

	var avgNitrogen float64
	err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("AVG(nitrogen_percentage)").
		Select(&avgNitrogen)
	if err != nil {
		return nil, err
	}
	analysis.AverageNitrogen = avgNitrogen

	var avgPhosphorus float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("AVG(phosphorus_percentage)").
		Select(&avgPhosphorus)
	if err != nil {
		return nil, err
	}
	analysis.AveragePhosphorus = avgPhosphorus

	var avgPotassium float64
	err = r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		ColumnExpr("AVG(potassium_percentage)").
		Select(&avgPotassium)
	if err != nil {
		return nil, err
	}
	analysis.AveragePotassium = avgPotassium

	return analysis, nil
}

// BulkUpdateStatus updates the status of multiple fertilizer types
func (r *fertilizerTypeRepository) BulkUpdateStatus(ctx context.Context, ids []string, status string) error {
	_, err := r.db.ModelContext(ctx, (*entity.FertilizerType)(nil)).
		Set("status = ?", status).
		Set("updated_at = ?", time.Now()).
		Where("id IN (?)", pg.In(ids)).
		Update()
	return err
}

// GetRecentFertilizerTypes retrieves recently created fertilizer types
func (r *fertilizerTypeRepository) GetRecentFertilizerTypes(ctx context.Context, limit int) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Order("created_at DESC").
		Limit(limit).
		Select()
	return fertilizerTypes, err
}

// SearchFertilizerTypes performs full-text search on fertilizer types
func (r *fertilizerTypeRepository) SearchFertilizerTypes(ctx context.Context, query string) ([]*entity.FertilizerType, error) {
	var fertilizerTypes []*entity.FertilizerType
	err := r.db.ModelContext(ctx, &fertilizerTypes).
		Where("name ILIKE ? OR description ILIKE ? OR manufacturer ILIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Select()
	return fertilizerTypes, err
}
