package entity

import (
	"time"
)

type IrrigationLog struct {
	ID                     string
	IrrigationScheduleID   string
	DeviceID               string
	StartedAt              *time.Time
	EndedAt                *time.Time
	PlannedDurationMinutes int
	ActualDurationMinutes  int
	WaterUsedLiters        float64
	WaterPressure          float64
	Status                 string
	FailureReason          string
	Notes                  string
	CreatedBy              string
	CreatedAt              time.Time
}

type CreateIrrigationLogRequest struct {
	IrrigationScheduleID   string
	DeviceID               string
	StartedAt              *time.Time
	EndedAt                *time.Time
	PlannedDurationMinutes int
	ActualDurationMinutes  int
	WaterUsedLiters        float64
	WaterPressure          float64
	Status                 string
	FailureReason          string
	Notes                  string
	CreatedBy              string
}

type UpdateIrrigationLogRequest struct {
	ID                     string
	IrrigationScheduleID   string
	DeviceID               string
	StartedAt              *time.Time
	EndedAt                *time.Time
	PlannedDurationMinutes int
	ActualDurationMinutes  int
	WaterUsedLiters        float64
	WaterPressure          float64
	Status                 string
	FailureReason          string
	Notes                  string
}

type IrrigationLogFilter struct {
	IrrigationScheduleID string
	DeviceID             string
	Status               string
	CreatedBy            string
	StartedAtFrom        *time.Time
	StartedAtTo          *time.Time
	EndedAtFrom          *time.Time
	EndedAtTo            *time.Time
	CreatedAtFrom        *time.Time
	CreatedAtTo          *time.Time
	Page                 int
	Limit                int
	SortBy               string
	SortOrder            string
}

type IrrigationLogResponse struct {
	ID                     string
	IrrigationScheduleID   string
	DeviceID               string
	StartedAt              *time.Time
	EndedAt                *time.Time
	PlannedDurationMinutes int
	ActualDurationMinutes  int
	WaterUsedLiters        float64
	WaterPressure          float64
	Status                 string
	FailureReason          string
	Notes                  string
	CreatedBy              string
	CreatedAt              time.Time
}

type ListIrrigationLogsResponse struct {
	IrrigationLogs []IrrigationLogResponse
	Total          int
	Page           int
	Limit          int
	TotalPages     int
}

type IrrigationLogStatistics struct {
	TotalLogs          int
	CompletedLogs      int
	FailedLogs         int
	InterruptedLogs    int
	ManualOverrideLogs int
	TotalWaterUsed     float64
	AverageWaterUsed   float64
	TotalDuration      int
	AverageDuration    float64
	SuccessRate        float64
	ByStatus           map[string]int
	ByDevice           map[string]int
}

type WaterUsageReport struct {
	TotalWaterUsed     float64
	AverageWaterUsed   float64
	MaxWaterUsed       float64
	MinWaterUsed       float64
	TotalExecutions    int
	WaterUsageByDay    map[string]float64
	WaterUsageByDevice map[string]float64
}

type EfficiencyReport struct {
	TotalPlannedDuration int
	TotalActualDuration  int
	DurationEfficiency   float64
	TotalPlannedWater    float64
	TotalActualWater     float64
	WaterEfficiency      float64
	OverallEfficiency    float64
	EfficiencyByDevice   map[string]float64
	EfficiencyBySchedule map[string]float64
}
