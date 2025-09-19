package irrigation_log_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IrrigationLogService) CreateIrrigationLog(ctx context.Context, req *irrigationLogP.CreateIrrigationLogRequest) (*irrigationLogP.CreateIrrigationLogResponse, error) {
	irrigationLogReq, err := s.createEntityIrrigationLogReq(req)
	if err != nil {
		return nil, err
	}
	irrigationLog, err := s.createIrrigationLogUsecase.Execute(ctx, irrigationLogReq)
	if err != nil {
		return nil, err
	}
	return &irrigationLogP.CreateIrrigationLogResponse{
		Success:       true,
		Message:       "Irrigation log created successfully",
		IrrigationLog: s.createProtoIrrigationLog(irrigationLog),
	}, nil
}

func (s *IrrigationLogService) createEntityIrrigationLogReq(req *irrigationLogP.CreateIrrigationLogRequest) (*entity.CreateIrrigationLogRequest, error) {
	irrigationLog := &entity.CreateIrrigationLogRequest{
		IrrigationScheduleID:   req.IrrigationScheduleId,
		DeviceID:               req.DeviceId,
		PlannedDurationMinutes: int(req.PlannedDurationMinutes),
		ActualDurationMinutes:  int(req.ActualDurationMinutes),
		WaterUsedLiters:        req.WaterUsedLiters,
		WaterPressure:          req.WaterPressure,
		Status:                 req.Status,
		FailureReason:          req.FailureReason,
		Notes:                  req.Notes,
		CreatedBy:              req.CreatedBy,
	}

	if req.StartedAt != nil {
		startedAt, err := time.Parse(time.RFC3339, req.StartedAt.String())
		if err != nil {
			return nil, err
		}
		irrigationLog.StartedAt = &startedAt
	}

	if req.EndedAt != nil {
		endedAt, err := time.Parse(time.RFC3339, req.EndedAt.String())
		if err != nil {
			return nil, err
		}
		irrigationLog.EndedAt = &endedAt
	}

	return irrigationLog, nil
}

func (s *IrrigationLogService) createProtoIrrigationLog(irrigationLog *entity.IrrigationLog) *irrigationLogP.IrrigationLog {
	response := &irrigationLogP.IrrigationLog{
		Id:                     irrigationLog.ID,
		IrrigationScheduleId:   irrigationLog.IrrigationScheduleID,
		DeviceId:               irrigationLog.DeviceID,
		PlannedDurationMinutes: int32(irrigationLog.PlannedDurationMinutes),
		ActualDurationMinutes:  int32(irrigationLog.ActualDurationMinutes),
		WaterUsedLiters:        irrigationLog.WaterUsedLiters,
		WaterPressure:          irrigationLog.WaterPressure,
		Status:                 irrigationLog.Status,
		FailureReason:          irrigationLog.FailureReason,
		Notes:                  irrigationLog.Notes,
		CreatedBy:              irrigationLog.CreatedBy,
	}

	if irrigationLog.StartedAt != nil {
		response.StartedAt = timestamppb.New(*irrigationLog.StartedAt)
	}

	if irrigationLog.EndedAt != nil {
		response.EndedAt = timestamppb.New(*irrigationLog.EndedAt)
	}

	return response
}
