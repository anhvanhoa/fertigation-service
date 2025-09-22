package irrigation_schedule_service

import (
	"context"

	commonP "github.com/anhvanhoa/sf-proto/gen/common/v1"
	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) DeleteIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.DeleteIrrigationScheduleRequest) (*commonP.CommonResponse, error) {
	err := s.deleteIrrigationScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &commonP.CommonResponse{
		Message: "Irrigation schedule deleted successfully",
	}, nil
}
