package fertilizer_schedule_service

import (
	"context"

	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
)

func (s *FertilizerScheduleService) GetFertilizerSchedule(ctx context.Context, req *proto_fertilizer_schedule.GetFertilizerScheduleRequest) (*proto_fertilizer_schedule.FertilizerScheduleResponse, error) {
	fertilizerSchedule, err := s.getFertilizerScheduleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.createProtoFertilizerSchedule(fertilizerSchedule), nil
}
