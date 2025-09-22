package irrigation_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IrrigationScheduleService) ListIrrigationSchedules(ctx context.Context, req *irrigationScheduleP.ListIrrigationSchedulesRequest) (*irrigationScheduleP.ListIrrigationSchedulesResponse, error) {
	filter := s.createEntityIrrigationScheduleFilter(req)
	response, err := s.listIrrigationScheduleUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListIrrigationSchedulesResponse(response), nil
}

func (s *IrrigationScheduleService) createEntityIrrigationScheduleFilter(req *irrigationScheduleP.ListIrrigationSchedulesRequest) *entity.IrrigationScheduleFilter {
	filter := &entity.IrrigationScheduleFilter{
		GrowingZoneID:   req.GrowingZoneId,
		PlantingCycleID: req.PlantingCycleId,
		ScheduleName:    req.ScheduleName,
		IrrigationType:  req.IrrigationType,
		Frequency:       req.Frequency,
		IsActive:        req.IsActive,
		FertilizerMix:   req.FertilizerMix,
		CreatedBy:       req.CreatedBy,
		Page:            int(req.Page),
		Limit:           int(req.Limit),
		SortBy:          req.SortBy,
		SortOrder:       req.SortOrder,
	}

	if req.CreatedAtFrom != nil {
		createdAtFrom, err := time.Parse(time.RFC3339, req.CreatedAtFrom.String())
		if err == nil {
			filter.CreatedAtFrom = &createdAtFrom
		}
	}

	if req.CreatedAtTo != nil {
		createdAtTo, err := time.Parse(time.RFC3339, req.CreatedAtTo.String())
		if err == nil {
			filter.CreatedAtTo = &createdAtTo
		}
	}

	if req.NextExecutionFrom != nil {
		nextExecutionFrom, err := time.Parse(time.RFC3339, req.NextExecutionFrom.String())
		if err == nil {
			filter.NextExecutionFrom = &nextExecutionFrom
		}
	}

	if req.NextExecutionTo != nil {
		nextExecutionTo, err := time.Parse(time.RFC3339, req.NextExecutionTo.String())
		if err == nil {
			filter.NextExecutionTo = &nextExecutionTo
		}
	}

	return filter
}

func (s *IrrigationScheduleService) createProtoListIrrigationSchedulesResponse(response common.PaginationResult[*entity.IrrigationSchedule]) *irrigationScheduleP.ListIrrigationSchedulesResponse {
	protoSchedules := make([]*irrigationScheduleP.IrrigationScheduleResponse, len(response.Data))
	for i, schedule := range response.Data {
		protoSchedules[i] = s.createProtoIrrigationScheduleFromResponse(schedule)
	}

	return &irrigationScheduleP.ListIrrigationSchedulesResponse{
		IrrigationSchedules: protoSchedules,
		Pagination: &proto_common.PaginationResponse{
			Total:      int32(response.Total),
			Page:       int32(response.Page),
			Limit:      int32(response.PageSize),
			TotalPages: int32(response.TotalPages),
		},
	}
}

func (s *IrrigationScheduleService) createProtoIrrigationScheduleFromResponse(schedule *entity.IrrigationSchedule) *irrigationScheduleP.IrrigationScheduleResponse {
	response := &irrigationScheduleP.IrrigationScheduleResponse{
		Id:                schedule.ID,
		GrowingZoneId:     schedule.GrowingZoneID,
		PlantingCycleId:   schedule.PlantingCycleID,
		ScheduleName:      schedule.ScheduleName,
		IrrigationType:    schedule.IrrigationType,
		StartTime:         schedule.StartTime,
		DurationMinutes:   int32(schedule.DurationMinutes),
		Frequency:         schedule.Frequency,
		DaysOfWeek:        schedule.DaysOfWeek,
		WaterAmountLiters: schedule.WaterAmountLiters,
		FertilizerMix:     schedule.FertilizerMix,
		IsActive:          schedule.IsActive,
		CreatedBy:         schedule.CreatedBy,
	}

	if schedule.LastExecuted != nil {
		response.LastExecuted = timestamppb.New(*schedule.LastExecuted)
	}

	if schedule.NextExecution != nil {
		response.NextExecution = timestamppb.New(*schedule.NextExecution)
	}

	return response
}
