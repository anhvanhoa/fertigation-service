package irrigation_log_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
)

func (s *IrrigationLogService) UpdateIrrigationLog(ctx context.Context, req *irrigationLogP.UpdateIrrigationLogRequest) (*irrigationLogP.IrrigationLogResponse, error) {
	irrigationLogReq, err := s.createEntityUpdateIrrigationLogReq(req)
	if err != nil {
		return nil, err
	}
	irrigationLog, err := s.updateIrrigationLogUsecase.Execute(ctx, irrigationLogReq)
	if err != nil {
		return nil, err
	}
	return s.createProtoIrrigationLog(irrigationLog), nil
}

func (s *IrrigationLogService) createEntityUpdateIrrigationLogReq(req *irrigationLogP.UpdateIrrigationLogRequest) (*entity.UpdateIrrigationLogRequest, error) {
	irrigationLog := &entity.UpdateIrrigationLogRequest{
		ID:                     req.Id,
		IrrigationScheduleID:   req.IrrigationScheduleId,
		DeviceID:               req.DeviceId,
		PlannedDurationMinutes: int(req.PlannedDurationMinutes),
		ActualDurationMinutes:  int(req.ActualDurationMinutes),
		WaterUsedLiters:        req.WaterUsedLiters,
		WaterPressure:          req.WaterPressure,
		Status:                 req.Status,
		FailureReason:          req.FailureReason,
		Notes:                  req.Notes,
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
