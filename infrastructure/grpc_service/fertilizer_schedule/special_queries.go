package fertilizer_schedule_service

import (
	"context"

	"github.com/anhvanhoa/service-core/common"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
)

func (s *FertilizerScheduleService) GetPendingSchedules(ctx context.Context, req *proto_fertilizer_schedule.Pagination) (*proto_fertilizer_schedule.ListFertilizerSchedulesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getPendingSchedulesUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerSchedulesResponse(response), nil
}

func (s *FertilizerScheduleService) GetUpcomingSchedules(ctx context.Context, req *proto_fertilizer_schedule.GetUpcomingSchedulesRequest) (*proto_fertilizer_schedule.ListFertilizerSchedulesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	response, err := s.getUpcomingSchedulesUsecase.Execute(ctx, int(req.Days), filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerSchedulesResponse(response), nil
}

func (s *FertilizerScheduleService) GetCompletedSchedules(ctx context.Context, req *proto_fertilizer_schedule.Pagination) (*proto_fertilizer_schedule.ListFertilizerSchedulesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getCompletedSchedulesUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerSchedulesResponse(response), nil
}

func (s *FertilizerScheduleService) GetSchedulesByPlantingCycle(ctx context.Context, req *proto_fertilizer_schedule.GetSchedulesByPlantingCycleRequest) (*proto_fertilizer_schedule.ListFertilizerSchedulesResponse, error) {
	filter := &common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	response, err := s.getSchedulesByPlantingCycleUsecase.Execute(ctx, req.PlantingCycleId, *filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerSchedulesResponse(response), nil
}
