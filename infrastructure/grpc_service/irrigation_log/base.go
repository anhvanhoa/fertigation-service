package irrigation_log_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/irrigation_log"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
)

type IrrigationLogService struct {
	irrigationLogP.UnsafeIrrigationLogServiceServer
	createIrrigationLogUsecase *irrigation_log.CreateIrrigationLogUsecase
	getIrrigationLogUsecase    *irrigation_log.GetIrrigationLogUsecase
	updateIrrigationLogUsecase *irrigation_log.UpdateIrrigationLogUsecase
	deleteIrrigationLogUsecase *irrigation_log.DeleteIrrigationLogUsecase
	listIrrigationLogUsecase   *irrigation_log.ListIrrigationLogUsecase
}

func NewIrrigationLogService(irrigationLogRepo repository.IrrigationLogRepository) irrigationLogP.IrrigationLogServiceServer {
	return &IrrigationLogService{
		createIrrigationLogUsecase: irrigation_log.NewCreateIrrigationLogUsecase(irrigationLogRepo),
		getIrrigationLogUsecase:    irrigation_log.NewGetIrrigationLogUsecase(irrigationLogRepo),
		updateIrrigationLogUsecase: irrigation_log.NewUpdateIrrigationLogUsecase(irrigationLogRepo),
		deleteIrrigationLogUsecase: irrigation_log.NewDeleteIrrigationLogUsecase(irrigationLogRepo),
		listIrrigationLogUsecase:   irrigation_log.NewListIrrigationLogUsecase(irrigationLogRepo),
	}
}
