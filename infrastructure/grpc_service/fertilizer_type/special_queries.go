package fertilizer_type_service

import (
	"context"

	"github.com/anhvanhoa/service-core/common"
	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

func (s *FertilizerTypeService) GetExpiredFertilizers(ctx context.Context, req *fertilizerTypeP.Pagination) (*fertilizerTypeP.ListFertilizerTypesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getExpiredFertilizersUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerTypesResponse(&response), nil
}

func (s *FertilizerTypeService) GetExpiringSoon(ctx context.Context, req *fertilizerTypeP.GetExpiringSoonRequest) (*fertilizerTypeP.ListFertilizerTypesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	response, err := s.getExpiringSoonUsecase.Execute(ctx, int(req.Days), filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerTypesResponse(&response), nil
}

func (s *FertilizerTypeService) GetFertilizerTypesByType(ctx context.Context, req *fertilizerTypeP.GetFertilizerTypesByTypeRequest) (*fertilizerTypeP.ListFertilizerTypesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	response, err := s.getFertilizerTypesByTypeUsecase.Execute(ctx, req.Type, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerTypesResponse(&response), nil
}
