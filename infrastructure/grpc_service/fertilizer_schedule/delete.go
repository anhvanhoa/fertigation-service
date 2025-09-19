package fertilizer_schedule_service

import (
	"context"

	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
)

func (s *FertilizerScheduleService) DeleteFertilizerSchedule(ctx context.Context, req *proto_fertilizer_schedule.DeleteFertilizerScheduleRequest) (*proto_fertilizer_schedule.DeleteFertilizerScheduleResponse, error) {
	err := s.deleteFertilizerScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_fertilizer_schedule.DeleteFertilizerScheduleResponse{
		Message: "Fertilizer schedule deleted successfully",
	}, nil
}
