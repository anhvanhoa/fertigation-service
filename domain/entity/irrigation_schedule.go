package entity

import (
	"time"
)

type IrrigationSchedule struct {
	nameTable         struct{} `pg:"irrigation_schedules"`
	ID                string
	GrowingZoneID     string
	PlantingCycleID   string
	ScheduleName      string
	IrrigationType    string
	StartTime         string
	DurationMinutes   int
	Frequency         string
	DaysOfWeek        string
	WaterAmountLiters float64
	FertilizerMix     bool
	IsActive          bool
	LastExecuted      *time.Time
	NextExecution     *time.Time
	CreatedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (i *IrrigationSchedule) TableName() any {
	return i.nameTable
}

type CreateIrrigationScheduleRequest struct {
	GrowingZoneID     string
	PlantingCycleID   string
	ScheduleName      string
	IrrigationType    string
	StartTime         string
	DurationMinutes   int
	Frequency         string
	DaysOfWeek        string
	WaterAmountLiters float64
	FertilizerMix     bool
	IsActive          bool
	CreatedBy         string
}

type UpdateIrrigationScheduleRequest struct {
	ID                string
	GrowingZoneID     string
	PlantingCycleID   string
	ScheduleName      string
	IrrigationType    string
	StartTime         string
	DurationMinutes   int
	Frequency         string
	DaysOfWeek        string
	WaterAmountLiters float64
	FertilizerMix     bool
	IsActive          bool
	LastExecuted      *time.Time
	NextExecution     *time.Time
}

type IrrigationScheduleFilter struct {
	GrowingZoneID     string
	PlantingCycleID   string
	ScheduleName      string
	IrrigationType    string
	Frequency         string
	IsActive          bool
	FertilizerMix     bool
	CreatedBy         string
	CreatedAtFrom     *time.Time
	CreatedAtTo       *time.Time
	NextExecutionFrom *time.Time
	NextExecutionTo   *time.Time
	Page              int
	Limit             int
	SortBy            string
	SortOrder         string
}

type IrrigationScheduleResponse struct {
	ID                string
	GrowingZoneID     string
	PlantingCycleID   string
	ScheduleName      string
	IrrigationType    string
	StartTime         string
	DurationMinutes   int
	Frequency         string
	DaysOfWeek        string
	WaterAmountLiters float64
	FertilizerMix     bool
	IsActive          bool
	LastExecuted      *time.Time
	NextExecution     *time.Time
	CreatedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type IrrigationScheduleStatistics struct {
	TotalSchedules         int
	ActiveSchedules        int
	InactiveSchedules      int
	FertilizerMixSchedules int
	ByType                 map[string]int
	ByFrequency            map[string]int
}
