package irrigation_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) GetActiveSchedules(ctx context.Context, req *irrigationScheduleP.GetActiveSchedulesRequest) (*irrigationScheduleP.GetActiveSchedulesResponse, error) {
	filter := &entity.IrrigationScheduleFilter{
		IsActive:  true,
		Page:      int(req.Page),
		Limit:     int(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getActiveSchedulesUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.GetActiveSchedulesResponse{
		Success: true,
		Message: "Active schedules retrieved successfully",
		Data:    s.createProtoListIrrigationSchedulesResponse(response),
	}, nil
}

func (s *IrrigationScheduleService) GetSchedulesByGrowingZone(ctx context.Context, req *irrigationScheduleP.GetSchedulesByGrowingZoneRequest) (*irrigationScheduleP.GetSchedulesByGrowingZoneResponse, error) {
	filter := &entity.IrrigationScheduleFilter{
		GrowingZoneID: req.GrowingZoneId,
		Page:          int(req.Page),
		Limit:         int(req.Limit),
		SortBy:        req.SortBy,
		SortOrder:     req.SortOrder,
	}
	response, err := s.getSchedulesByGrowingZoneUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.GetSchedulesByGrowingZoneResponse{
		Success: true,
		Message: "Schedules by growing zone retrieved successfully",
		Data:    s.createProtoListIrrigationSchedulesResponse(response),
	}, nil
}

func (s *IrrigationScheduleService) GetSchedulesForExecution(ctx context.Context, req *irrigationScheduleP.GetSchedulesForExecutionRequest) (*irrigationScheduleP.GetSchedulesForExecutionResponse, error) {
	filter := &entity.IrrigationScheduleFilter{
		IsActive:  true,
		Page:      int(req.Page),
		Limit:     int(req.Limit),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
	response, err := s.getSchedulesForExecutionUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.GetSchedulesForExecutionResponse{
		Success: true,
		Message: "Schedules for execution retrieved successfully",
		Data:    s.createProtoListIrrigationSchedulesResponse(response),
	}, nil
}
