package fertilizer_type_service

import (
	"context"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

func (s *FertilizerTypeService) DeleteFertilizerType(ctx context.Context, req *fertilizerTypeP.DeleteFertilizerTypeRequest) (*fertilizerTypeP.DeleteFertilizerTypeResponse, error) {
	err := s.deleteFertilizerTypeUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &fertilizerTypeP.DeleteFertilizerTypeResponse{
		Message: "Fertilizer type deleted successfully",
	}, nil
}
