package irrigation_log_service

import (
	"context"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
)

func (s *IrrigationLogService) GetIrrigationLog(ctx context.Context, req *irrigationLogP.GetIrrigationLogRequest) (*irrigationLogP.GetIrrigationLogResponse, error) {
	irrigationLog, err := s.getIrrigationLogUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &irrigationLogP.GetIrrigationLogResponse{
		Success:       true,
		Message:       "Irrigation log retrieved successfully",
		IrrigationLog: s.createProtoIrrigationLog(irrigationLog),
	}, nil
}
