package fertilizer_type_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/fertilizer_type"

	"github.com/anhvanhoa/service-core/utils"
	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

type FertilizerTypeService struct {
	fertilizerTypeP.UnsafeFertilizerTypeServiceServer
	createFertilizerTypeUsecase     fertilizer_type.CreateFertilizerTypeUsecaseI
	getFertilizerTypeUsecase        fertilizer_type.GetFertilizerTypeUsecaseI
	updateFertilizerTypeUsecase     fertilizer_type.UpdateFertilizerTypeUsecaseI
	deleteFertilizerTypeUsecase     fertilizer_type.DeleteFertilizerTypeUsecaseI
	listFertilizerTypeUsecase       fertilizer_type.ListFertilizerTypeUsecaseI
	getExpiredFertilizersUsecase    fertilizer_type.GetExpiredFertilizersUsecaseI
	getExpiringSoonUsecase          fertilizer_type.GetExpiringSoonUsecaseI
	getFertilizerTypesByTypeUsecase fertilizer_type.GetFertilizerTypesByTypeUsecaseI
}

func NewFertilizerTypeService(fertilizerTypeRepo repository.FertilizerTypeRepository) fertilizerTypeP.FertilizerTypeServiceServer {
	helper := utils.NewHelper()
	return &FertilizerTypeService{
		createFertilizerTypeUsecase:     fertilizer_type.NewCreateFertilizerTypeUsecase(fertilizerTypeRepo),
		getFertilizerTypeUsecase:        fertilizer_type.NewGetFertilizerTypeUsecase(fertilizerTypeRepo),
		updateFertilizerTypeUsecase:     fertilizer_type.NewUpdateFertilizerTypeUsecase(fertilizerTypeRepo),
		deleteFertilizerTypeUsecase:     fertilizer_type.NewDeleteFertilizerTypeUsecase(fertilizerTypeRepo),
		listFertilizerTypeUsecase:       fertilizer_type.NewListFertilizerTypeUsecase(fertilizerTypeRepo, helper),
		getExpiredFertilizersUsecase:    fertilizer_type.NewGetExpiredFertilizersUsecase(fertilizerTypeRepo, helper),
		getExpiringSoonUsecase:          fertilizer_type.NewGetExpiringSoonUsecase(fertilizerTypeRepo, helper),
		getFertilizerTypesByTypeUsecase: fertilizer_type.NewGetFertilizerTypesByTypeUsecase(fertilizerTypeRepo, helper),
	}
}
