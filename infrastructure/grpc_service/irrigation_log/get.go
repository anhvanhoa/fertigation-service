package irrigation_log_service

import (
	"context"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
)

func (s *IrrigationLogService) GetIrrigationLog(ctx context.Context, req *irrigationLogP.GetIrrigationLogRequest) (*irrigationLogP.IrrigationLogResponse, error) {
	irrigationLog, err := s.getIrrigationLogUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.createProtoIrrigationLog(irrigationLog), nil
}
