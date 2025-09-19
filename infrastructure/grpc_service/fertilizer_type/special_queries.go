package fertilizer_type_service

import (
	"context"
	"fertigation-Service/domain/entity"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

func (s *FertilizerTypeService) GetExpiredFertilizers(ctx context.Context, req *fertilizerTypeP.GetExpiredFertilizersRequest) (*fertilizerTypeP.GetExpiredFertilizersResponse, error) {
	filter := &entity.FertilizerTypeFilter{
		Page:      int(req.Page),
		Limit:     int(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getExpiredFertilizersUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &fertilizerTypeP.GetExpiredFertilizersResponse{
		Success: true,
		Message: "Expired fertilizers retrieved successfully",
		Data:    s.createProtoListFertilizerTypesResponse(response),
	}, nil
}

func (s *FertilizerTypeService) GetExpiringSoon(ctx context.Context, req *fertilizerTypeP.GetExpiringSoonRequest) (*fertilizerTypeP.GetExpiringSoonResponse, error) {
	filter := &entity.FertilizerTypeFilter{
		Page:      int(req.Page),
		Limit:     int(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getExpiringSoonUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &fertilizerTypeP.GetExpiringSoonResponse{
		Success: true,
		Message: "Expiring soon fertilizers retrieved successfully",
		Data:    s.createProtoListFertilizerTypesResponse(response),
	}, nil
}

func (s *FertilizerTypeService) GetFertilizerTypesByType(ctx context.Context, req *fertilizerTypeP.GetFertilizerTypesByTypeRequest) (*fertilizerTypeP.GetFertilizerTypesByTypeResponse, error) {
	filter := &entity.FertilizerTypeFilter{
		Type:      req.Type,
		Page:      int(req.Page),
		Limit:     int(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getFertilizerTypesByTypeUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &fertilizerTypeP.GetFertilizerTypesByTypeResponse{
		Success: true,
		Message: "Fertilizer types by type retrieved successfully",
		Data:    s.createProtoListFertilizerTypesResponse(response),
	}, nil
}
