package irrigation_log_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IrrigationLogService) ListIrrigationLogs(ctx context.Context, req *irrigationLogP.ListIrrigationLogsRequest) (*irrigationLogP.ListIrrigationLogsResponse, error) {
	filter := s.createEntityIrrigationLogFilter(req)
	response, err := s.listIrrigationLogUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &irrigationLogP.ListIrrigationLogsResponse{
		Success: true,
		Message: "Irrigation logs retrieved successfully",
		Data:    s.createProtoListIrrigationLogsResponse(response),
	}, nil
}

func (s *IrrigationLogService) createEntityIrrigationLogFilter(req *irrigationLogP.ListIrrigationLogsRequest) *entity.IrrigationLogFilter {
	filter := &entity.IrrigationLogFilter{
		IrrigationScheduleID: req.IrrigationScheduleId,
		DeviceID:             req.DeviceId,
		Status:               req.Status,
		CreatedBy:            req.CreatedBy,
		Page:                 int(req.Page),
		Limit:                int(req.Limit),
		SortBy:               req.SortBy,
		SortOrder:            req.SortOrder,
	}

	if req.StartedAtFrom != nil {
		startedAtFrom, err := time.Parse(time.RFC3339, req.StartedAtFrom.String())
		if err == nil {
			filter.StartedAtFrom = &startedAtFrom
		}
	}

	if req.StartedAtTo != nil {
		startedAtTo, err := time.Parse(time.RFC3339, req.StartedAtTo.String())
		if err == nil {
			filter.StartedAtTo = &startedAtTo
		}
	}

	if req.EndedAtFrom != nil {
		endedAtFrom, err := time.Parse(time.RFC3339, req.EndedAtFrom.String())
		if err == nil {
			filter.EndedAtFrom = &endedAtFrom
		}
	}

	if req.EndedAtTo != nil {
		endedAtTo, err := time.Parse(time.RFC3339, req.EndedAtTo.String())
		if err == nil {
			filter.EndedAtTo = &endedAtTo
		}
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

	return filter
}

func (s *IrrigationLogService) createProtoListIrrigationLogsResponse(response *entity.ListIrrigationLogsResponse) *irrigationLogP.ListIrrigationLogsData {
	protoLogs := make([]*irrigationLogP.IrrigationLog, len(response.IrrigationLogs))
	for i, log := range response.IrrigationLogs {
		protoLogs[i] = s.createProtoIrrigationLogFromResponse(&log)
	}

	return &irrigationLogP.ListIrrigationLogsData{
		IrrigationLogs: protoLogs,
		Total:          int32(response.Total),
		Page:           int32(response.Page),
		Limit:          int32(response.Limit),
		TotalPages:     int32(response.TotalPages),
	}
}

func (s *IrrigationLogService) createProtoIrrigationLogFromResponse(log *entity.IrrigationLogResponse) *irrigationLogP.IrrigationLog {
	response := &irrigationLogP.IrrigationLog{
		Id:                     log.ID,
		IrrigationScheduleId:   log.IrrigationScheduleID,
		DeviceId:               log.DeviceID,
		PlannedDurationMinutes: int32(log.PlannedDurationMinutes),
		ActualDurationMinutes:  int32(log.ActualDurationMinutes),
		WaterUsedLiters:        log.WaterUsedLiters,
		WaterPressure:          log.WaterPressure,
		Status:                 log.Status,
		FailureReason:          log.FailureReason,
		Notes:                  log.Notes,
		CreatedBy:              log.CreatedBy,
	}

	if log.StartedAt != nil {
		response.StartedAt = timestamppb.New(*log.StartedAt)
	}

	if log.EndedAt != nil {
		response.EndedAt = timestamppb.New(*log.EndedAt)
	}

	return response
}
