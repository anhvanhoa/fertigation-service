package entity

import (
	"time"
)

type FertilizerSchedule struct {
	nameTable           struct{} `pg:"fertilizer_schedules"`
	ID                  string
	PlantingCycleID     string
	FertilizerTypeID    string
	ApplicationDate     *time.Time
	Dosage              float64
	Unit                string
	ApplicationMethod   string
	GrowthStage         string
	WeatherConditions   string
	SoilConditions      string
	IsCompleted         bool
	CompletedDate       *time.Time
	ActualDosage        float64
	EffectivenessRating int
	Notes               string
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (f *FertilizerSchedule) TableName() any {
	return f.nameTable
}

type CreateFertilizerScheduleRequest struct {
	PlantingCycleID     string
	FertilizerTypeID    string
	ApplicationDate     *time.Time
	Dosage              float64
	Unit                string
	ApplicationMethod   string
	GrowthStage         string
	WeatherConditions   string
	SoilConditions      string
	IsCompleted         bool
	CompletedDate       *time.Time
	ActualDosage        float64
	EffectivenessRating int
	Notes               string
	CreatedBy           string
}

type UpdateFertilizerScheduleRequest struct {
	ID                  string
	PlantingCycleID     string
	FertilizerTypeID    string
	ApplicationDate     *time.Time
	Dosage              float64
	Unit                string
	ApplicationMethod   string
	GrowthStage         string
	WeatherConditions   string
	SoilConditions      string
	IsCompleted         bool
	CompletedDate       *time.Time
	ActualDosage        float64
	EffectivenessRating int
	Notes               string
}

type FertilizerScheduleFilter struct {
	PlantingCycleID     string
	FertilizerTypeID    string
	ApplicationMethod   string
	GrowthStage         string
	IsCompleted         bool
	CreatedBy           string
	ApplicationDateFrom *time.Time
	ApplicationDateTo   *time.Time
	CompletedDateFrom   *time.Time
	CompletedDateTo     *time.Time
	CreatedAtFrom       *time.Time
	CreatedAtTo         *time.Time
	Page                int
	Limit               int
	SortBy              string
	SortOrder           string
}

type FertilizerScheduleResponse struct {
	ID                  string
	PlantingCycleID     string
	FertilizerTypeID    string
	ApplicationDate     *time.Time
	Dosage              float64
	Unit                string
	ApplicationMethod   string
	GrowthStage         string
	WeatherConditions   string
	SoilConditions      string
	IsCompleted         bool
	CompletedDate       *time.Time
	ActualDosage        float64
	EffectivenessRating int
	Notes               string
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type FertilizerScheduleStatistics struct {
	TotalSchedules        int
	CompletedSchedules    int
	PendingSchedules      int
	OverdueSchedules      int
	AverageEffectiveness  float64
	TotalDosageUsed       float64
	AverageDosageUsed     float64
	ByGrowthStage         map[string]int
	ByApplicationMethod   map[string]int
	ByEffectivenessRating map[int]int
	ByWeatherConditions   map[string]int
	BySoilConditions      map[string]int
}

type FertilizerEffectivenessReport struct {
	TotalSchedules             int
	AverageEffectiveness       float64
	MaxEffectiveness           int
	MinEffectiveness           int
	EffectivenessByGrowthStage map[string]float64
	EffectivenessByMethod      map[string]float64
	EffectivenessByWeather     map[string]float64
	EffectivenessBySoil        map[string]float64
}

type FertilizerDosageReport struct {
	TotalDosageUsed        float64
	AverageDosageUsed      float64
	MaxDosageUsed          float64
	MinDosageUsed          float64
	DosageByGrowthStage    map[string]float64
	DosageByMethod         map[string]float64
	DosageByFertilizerType map[string]float64
	DosageByUnit           map[string]float64
}

type FertilizerGrowthStageReport struct {
	TotalSchedules        int
	ByGrowthStage         map[string]int
	AverageDosageByStage  map[string]float64
	EffectivenessByStage  map[string]float64
	CompletionRateByStage map[string]float64
}
