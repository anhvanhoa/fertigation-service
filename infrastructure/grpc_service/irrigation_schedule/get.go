package irrigation_schedule_service

import (
	"context"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) GetIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.GetIrrigationScheduleRequest) (*irrigationScheduleP.IrrigationScheduleResponse, error) {
	irrigationSchedule, err := s.getIrrigationScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.createProtoIrrigationSchedule(irrigationSchedule), nil
}
