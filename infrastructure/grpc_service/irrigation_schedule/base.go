package irrigation_schedule_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/irrigation_schedule"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

type IrrigationScheduleService struct {
	irrigationScheduleP.UnsafeIrrigationScheduleServiceServer
	createIrrigationScheduleUsecase  *irrigation_schedule.CreateIrrigationScheduleUsecase
	getIrrigationScheduleUsecase     *irrigation_schedule.GetIrrigationScheduleUsecase
	updateIrrigationScheduleUsecase  *irrigation_schedule.UpdateIrrigationScheduleUsecase
	deleteIrrigationScheduleUsecase  *irrigation_schedule.DeleteIrrigationScheduleUsecase
	listIrrigationScheduleUsecase    *irrigation_schedule.ListIrrigationScheduleUsecase
	getActiveSchedulesUsecase        *irrigation_schedule.GetActiveSchedulesUsecase
	getSchedulesByGrowingZoneUsecase *irrigation_schedule.GetSchedulesByGrowingZoneUsecase
	getSchedulesForExecutionUsecase  *irrigation_schedule.GetSchedulesForExecutionUsecase
}

func NewIrrigationScheduleService(irrigationScheduleRepo repository.IrrigationScheduleRepository) irrigationScheduleP.IrrigationScheduleServiceServer {
	return &IrrigationScheduleService{
		createIrrigationScheduleUsecase:  irrigation_schedule.NewCreateIrrigationScheduleUsecase(irrigationScheduleRepo),
		getIrrigationScheduleUsecase:     irrigation_schedule.NewGetIrrigationScheduleUsecase(irrigationScheduleRepo),
		updateIrrigationScheduleUsecase:  irrigation_schedule.NewUpdateIrrigationScheduleUsecase(irrigationScheduleRepo),
		deleteIrrigationScheduleUsecase:  irrigation_schedule.NewDeleteIrrigationScheduleUsecase(irrigationScheduleRepo),
		listIrrigationScheduleUsecase:    irrigation_schedule.NewListIrrigationScheduleUsecase(irrigationScheduleRepo),
		getActiveSchedulesUsecase:        irrigation_schedule.NewGetActiveSchedulesUsecase(irrigationScheduleRepo),
		getSchedulesByGrowingZoneUsecase: irrigation_schedule.NewGetSchedulesByGrowingZoneUsecase(irrigationScheduleRepo),
		getSchedulesForExecutionUsecase:  irrigation_schedule.NewGetSchedulesForExecutionUsecase(irrigationScheduleRepo),
	}
}
