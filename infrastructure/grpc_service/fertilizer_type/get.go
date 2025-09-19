package fertilizer_type_service

import (
	"context"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

func (s *FertilizerTypeService) GetFertilizerType(ctx context.Context, req *fertilizerTypeP.GetFertilizerTypeRequest) (*fertilizerTypeP.FertilizerTypeResponse, error) {
	fertilizerType, err := s.getFertilizerTypeUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.createProtoFertilizerType(fertilizerType), nil
}
