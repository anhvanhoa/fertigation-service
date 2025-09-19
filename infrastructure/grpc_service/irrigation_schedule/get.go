package irrigation_schedule_service

import (
	"context"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) GetIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.GetIrrigationScheduleRequest) (*irrigationScheduleP.GetIrrigationScheduleResponse, error) {
	irrigationSchedule, err := s.getIrrigationScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.GetIrrigationScheduleResponse{
		Success:            true,
		Message:            "Irrigation schedule retrieved successfully",
		IrrigationSchedule: s.createProtoIrrigationSchedule(irrigationSchedule),
	}, nil
}
