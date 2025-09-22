package irrigation_schedule_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/irrigation_schedule"

	"github.com/anhvanhoa/service-core/utils"
	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

type IrrigationScheduleService struct {
	irrigationScheduleP.UnsafeIrrigationScheduleServiceServer
	createIrrigationScheduleUsecase  irrigation_schedule.CreateIrrigationScheduleUsecaseI
	getIrrigationScheduleUsecase     irrigation_schedule.GetIrrigationScheduleUsecaseI
	updateIrrigationScheduleUsecase  irrigation_schedule.UpdateIrrigationScheduleUsecaseI
	deleteIrrigationScheduleUsecase  irrigation_schedule.DeleteIrrigationScheduleUsecaseI
	listIrrigationScheduleUsecase    irrigation_schedule.ListIrrigationScheduleUsecaseI
	getActiveSchedulesUsecase        irrigation_schedule.GetActiveSchedulesUsecaseI
	getSchedulesByGrowingZoneUsecase irrigation_schedule.GetSchedulesByGrowingZoneUsecaseI
	getSchedulesForExecutionUsecase  irrigation_schedule.GetSchedulesForExecutionUsecaseI
}

func NewIrrigationScheduleService(irrigationScheduleRepo repository.IrrigationScheduleRepository) irrigationScheduleP.IrrigationScheduleServiceServer {
	helper := utils.NewHelper()
	return &IrrigationScheduleService{
		createIrrigationScheduleUsecase:  irrigation_schedule.NewCreateIrrigationScheduleUsecase(irrigationScheduleRepo),
		getIrrigationScheduleUsecase:     irrigation_schedule.NewGetIrrigationScheduleUsecase(irrigationScheduleRepo),
		updateIrrigationScheduleUsecase:  irrigation_schedule.NewUpdateIrrigationScheduleUsecase(irrigationScheduleRepo),
		deleteIrrigationScheduleUsecase:  irrigation_schedule.NewDeleteIrrigationScheduleUsecase(irrigationScheduleRepo),
		listIrrigationScheduleUsecase:    irrigation_schedule.NewListIrrigationScheduleUsecase(irrigationScheduleRepo, helper),
		getActiveSchedulesUsecase:        irrigation_schedule.NewGetActiveSchedulesUsecase(irrigationScheduleRepo, helper),
		getSchedulesByGrowingZoneUsecase: irrigation_schedule.NewGetSchedulesByGrowingZoneUsecase(irrigationScheduleRepo, helper),
		getSchedulesForExecutionUsecase:  irrigation_schedule.NewGetSchedulesForExecutionUsecase(irrigationScheduleRepo, helper),
	}
}
