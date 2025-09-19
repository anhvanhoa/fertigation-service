package fertilizer_schedule_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/fertilizer_schedule"

	"github.com/anhvanhoa/service-core/utils"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
)

type FertilizerScheduleService struct {
	proto_fertilizer_schedule.UnsafeFertilizerScheduleServiceServer
	createFertilizerScheduleUsecase    fertilizer_schedule.CreateFertilizerScheduleUsecaseI
	getFertilizerScheduleUsecase       fertilizer_schedule.GetFertilizerScheduleUsecaseI
	updateFertilizerScheduleUsecase    fertilizer_schedule.UpdateFertilizerScheduleUsecaseI
	deleteFertilizerScheduleUsecase    fertilizer_schedule.DeleteFertilizerScheduleUsecaseI
	listFertilizerScheduleUsecase      fertilizer_schedule.ListFertilizerScheduleUsecaseI
	getPendingSchedulesUsecase         fertilizer_schedule.GetPendingSchedulesUsecaseI
	getUpcomingSchedulesUsecase        fertilizer_schedule.GetUpcomingSchedulesUsecaseI
	getCompletedSchedulesUsecase       fertilizer_schedule.GetCompletedSchedulesUsecaseI
	getSchedulesByPlantingCycleUsecase fertilizer_schedule.GetSchedulesByPlantingCycleUsecaseI
}

func NewFertilizerScheduleService(fertilizerScheduleRepo repository.FertilizerScheduleRepository) proto_fertilizer_schedule.FertilizerScheduleServiceServer {
	helper := utils.NewHelper()
	return &FertilizerScheduleService{
		createFertilizerScheduleUsecase:    fertilizer_schedule.NewCreateFertilizerScheduleUsecase(fertilizerScheduleRepo),
		getFertilizerScheduleUsecase:       fertilizer_schedule.NewGetFertilizerScheduleUsecase(fertilizerScheduleRepo),
		updateFertilizerScheduleUsecase:    fertilizer_schedule.NewUpdateFertilizerScheduleUsecase(fertilizerScheduleRepo),
		deleteFertilizerScheduleUsecase:    fertilizer_schedule.NewDeleteFertilizerScheduleUsecase(fertilizerScheduleRepo),
		listFertilizerScheduleUsecase:      fertilizer_schedule.NewListFertilizerScheduleUsecase(fertilizerScheduleRepo, helper),
		getPendingSchedulesUsecase:         fertilizer_schedule.NewGetPendingSchedulesUsecase(fertilizerScheduleRepo, helper),
		getUpcomingSchedulesUsecase:        fertilizer_schedule.NewGetUpcomingSchedulesUsecase(fertilizerScheduleRepo, helper),
		getCompletedSchedulesUsecase:       fertilizer_schedule.NewGetCompletedSchedulesUsecase(fertilizerScheduleRepo, helper),
		getSchedulesByPlantingCycleUsecase: fertilizer_schedule.NewGetSchedulesByPlantingCycleUsecase(fertilizerScheduleRepo, helper),
	}
}
