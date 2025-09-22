package irrigation_schedule_service

import (
	"context"
	"time"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) GetActiveSchedules(ctx context.Context, req *common_proto.PaginationRequest) (*irrigationScheduleP.ListIrrigationSchedulesResponse, error) {
	filter := &common.Pagination{
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getActiveSchedulesUsecase.Execute(ctx, *filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListIrrigationSchedulesResponse(response), nil
}

func (s *IrrigationScheduleService) GetSchedulesByGrowingZone(ctx context.Context, req *irrigationScheduleP.GetSchedulesByGrowingZoneRequest) (*irrigationScheduleP.ListIrrigationSchedulesResponse, error) {
	filter := &common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	response, err := s.getSchedulesByGrowingZoneUsecase.Execute(ctx, req.GrowingZoneId, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListIrrigationSchedulesResponse(response), nil
}

func (s *IrrigationScheduleService) GetSchedulesForExecution(ctx context.Context, req *irrigationScheduleP.GetSchedulesForExecutionRequest) (*irrigationScheduleP.ListIrrigationSchedulesResponse, error) {
	filter := common.Pagination{
		Page:      int(req.Pagination.Page),
		PageSize:  int(req.Pagination.PageSize),
		SortBy:    req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	fromTime := req.FromTime.AsTime().Format(time.RFC3339)
	toTime := req.ToTime.AsTime().Format(time.RFC3339)
	response, err := s.getSchedulesForExecutionUsecase.Execute(ctx, fromTime, toTime, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListIrrigationSchedulesResponse(response), nil
}
