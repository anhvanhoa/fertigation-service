package entity

import (
	"time"
)

type FertilizerType struct {
	ID                   string
	Name                 string
	Type                 string
	NPKRatio             string
	NitrogenPercentage   float64
	PhosphorusPercentage float64
	PotassiumPercentage  float64
	TraceElements        string
	ApplicationMethod    string
	DosagePerPlant       float64
	DosagePerM2          float64
	Unit                 string
	Manufacturer         string
	BatchNumber          string
	ExpiryDate           *time.Time
	CostPerUnit          float64
	Description          string
	SafetyNotes          string
	Status               string
	CreatedBy            string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CreateFertilizerTypeRequest struct {
	Name                 string
	Type                 string
	NPKRatio             string
	NitrogenPercentage   float64
	PhosphorusPercentage float64
	PotassiumPercentage  float64
	TraceElements        string
	ApplicationMethod    string
	DosagePerPlant       float64
	DosagePerM2          float64
	Unit                 string
	Manufacturer         string
	BatchNumber          string
	ExpiryDate           *time.Time
	CostPerUnit          float64
	Description          string
	SafetyNotes          string
	Status               string
	CreatedBy            string
}

type UpdateFertilizerTypeRequest struct {
	ID                   string
	Name                 string
	Type                 string
	NPKRatio             string
	NitrogenPercentage   float64
	PhosphorusPercentage float64
	PotassiumPercentage  float64
	TraceElements        string
	ApplicationMethod    string
	DosagePerPlant       float64
	DosagePerM2          float64
	Unit                 string
	Manufacturer         string
	BatchNumber          string
	ExpiryDate           *time.Time
	CostPerUnit          float64
	Description          string
	SafetyNotes          string
	Status               string
}

type FertilizerTypeFilter struct {
	Name              string
	Type              string
	ApplicationMethod string
	Status            string
	Manufacturer      string
	CreatedBy         string
	ExpiryDateFrom    *time.Time
	ExpiryDateTo      *time.Time
	CreatedAtFrom     *time.Time
	CreatedAtTo       *time.Time
	Page              int
	PageSize          int
	SortBy            string
	SortOrder         string
}

type FertilizerTypeStatistics struct {
	TotalFertilizerTypes    int
	ActiveFertilizerTypes   int
	InactiveFertilizerTypes int
	ExpiredFertilizerTypes  int
	ByType                  map[string]int
	ByStatus                map[string]int
	ByApplicationMethod     map[string]int
	ByManufacturer          map[string]int
	AverageCost             float64
	TotalCost               float64
}

type FertilizerExpiryReport struct {
	TotalFertilizerTypes int
	ExpiredCount         int
	ExpiringIn30Days     int
	ExpiringIn60Days     int
	ExpiringIn90Days     int
	NoExpiryDate         int
	ExpiryByMonth        map[string]int
	ExpiryByManufacturer map[string]int
}

type FertilizerCostAnalysis struct {
	TotalCost               float64
	AverageCost             float64
	MinCost                 float64
	MaxCost                 float64
	CostByType              map[string]float64
	CostByManufacturer      map[string]float64
	CostByApplicationMethod map[string]float64
}

type FertilizerNPKAnalysis struct {
	AverageNitrogen        float64
	AveragePhosphorus      float64
	AveragePotassium       float64
	NPKDistribution        map[string]int
	NPKByType              map[string]map[string]float64
	NPKByApplicationMethod map[string]map[string]float64
}
