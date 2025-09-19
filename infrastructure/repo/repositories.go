package repo

import (
	"fertigation-Service/domain/repository"

	"github.com/go-pg/pg/v10"
)

// Repositories holds all repository implementations
type Repositories struct {
	IrrigationScheduleRepository repository.IrrigationScheduleRepository
	IrrigationLogRepository      repository.IrrigationLogRepository
	FertilizerTypeRepository     repository.FertilizerTypeRepository
	FertilizerScheduleRepository repository.FertilizerScheduleRepository
}

// NewRepositories creates a new instance of all repositories
func NewRepositories(db *pg.DB) *Repositories {
	return &Repositories{
		IrrigationScheduleRepository: NewIrrigationScheduleRepository(db),
		IrrigationLogRepository:      NewIrrigationLogRepository(db),
		FertilizerTypeRepository:     NewFertilizerTypeRepository(db),
		FertilizerScheduleRepository: NewFertilizerScheduleRepository(db),
	}
}
