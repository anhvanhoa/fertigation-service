package irrigation_log_service

import (
	"context"

	irrigationLogP "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
)

func (s *IrrigationLogService) DeleteIrrigationLog(ctx context.Context, req *irrigationLogP.DeleteIrrigationLogRequest) (*irrigationLogP.DeleteIrrigationLogResponse, error) {
	err := s.deleteIrrigationLogUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &irrigationLogP.DeleteIrrigationLogResponse{
		Success: true,
		Message: "Irrigation log deleted successfully",
	}, nil
}
