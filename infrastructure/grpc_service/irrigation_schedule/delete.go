package irrigation_schedule_service

import (
	"context"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) DeleteIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.DeleteIrrigationScheduleRequest) (*irrigationScheduleP.DeleteIrrigationScheduleResponse, error) {
	err := s.deleteIrrigationScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.DeleteIrrigationScheduleResponse{
		Success: true,
		Message: "Irrigation schedule deleted successfully",
	}, nil
}
