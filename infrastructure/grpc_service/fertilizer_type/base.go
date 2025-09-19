package fertilizer_type_service

import (
	"fertigation-Service/domain/repository"
	"fertigation-Service/domain/usecase/fertilizer_type"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

type FertilizerTypeService struct {
	fertilizerTypeP.UnsafeFertilizerTypeServiceServer
	createFertilizerTypeUsecase     *fertilizer_type.CreateFertilizerTypeUsecase
	getFertilizerTypeUsecase        *fertilizer_type.GetFertilizerTypeUsecase
	updateFertilizerTypeUsecase     *fertilizer_type.UpdateFertilizerTypeUsecase
	deleteFertilizerTypeUsecase     *fertilizer_type.DeleteFertilizerTypeUsecase
	listFertilizerTypeUsecase       *fertilizer_type.ListFertilizerTypeUsecase
	getExpiredFertilizersUsecase    *fertilizer_type.GetExpiredFertilizersUsecase
	getExpiringSoonUsecase          *fertilizer_type.GetExpiringSoonUsecase
	getFertilizerTypesByTypeUsecase *fertilizer_type.GetFertilizerTypesByTypeUsecase
}

func NewFertilizerTypeService(fertilizerTypeRepo repository.FertilizerTypeRepository) fertilizerTypeP.FertilizerTypeServiceServer {
	return &FertilizerTypeService{
		createFertilizerTypeUsecase:     fertilizer_type.NewCreateFertilizerTypeUsecase(fertilizerTypeRepo),
		getFertilizerTypeUsecase:        fertilizer_type.NewGetFertilizerTypeUsecase(fertilizerTypeRepo),
		updateFertilizerTypeUsecase:     fertilizer_type.NewUpdateFertilizerTypeUsecase(fertilizerTypeRepo),
		deleteFertilizerTypeUsecase:     fertilizer_type.NewDeleteFertilizerTypeUsecase(fertilizerTypeRepo),
		listFertilizerTypeUsecase:       fertilizer_type.NewListFertilizerTypeUsecase(fertilizerTypeRepo),
		getExpiredFertilizersUsecase:    fertilizer_type.NewGetExpiredFertilizersUsecase(fertilizerTypeRepo),
		getExpiringSoonUsecase:          fertilizer_type.NewGetExpiringSoonUsecase(fertilizerTypeRepo),
		getFertilizerTypesByTypeUsecase: fertilizer_type.NewGetFertilizerTypesByTypeUsecase(fertilizerTypeRepo),
	}
}
