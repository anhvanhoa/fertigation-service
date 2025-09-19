package fertilizer_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
)

func (s *FertilizerScheduleService) UpdateFertilizerSchedule(ctx context.Context, req *proto_fertilizer_schedule.UpdateFertilizerScheduleRequest) (*proto_fertilizer_schedule.FertilizerScheduleResponse, error) {
	fertilizerScheduleReq, err := s.createEntityUpdateFertilizerScheduleReq(req)
	if err != nil {
		return nil, err
	}
	fertilizerSchedule, err := s.updateFertilizerScheduleUsecase.Execute(ctx, fertilizerScheduleReq)
	if err != nil {
		return nil, err
	}
	return s.createProtoFertilizerSchedule(fertilizerSchedule), nil
}

func (s *FertilizerScheduleService) createEntityUpdateFertilizerScheduleReq(req *proto_fertilizer_schedule.UpdateFertilizerScheduleRequest) (*entity.UpdateFertilizerScheduleRequest, error) {
	fertilizerSchedule := &entity.UpdateFertilizerScheduleRequest{
		ID:                  req.Id,
		PlantingCycleID:     req.PlantingCycleId,
		FertilizerTypeID:    req.FertilizerTypeId,
		Dosage:              req.Dosage,
		Unit:                req.Unit,
		ApplicationMethod:   req.ApplicationMethod,
		GrowthStage:         req.GrowthStage,
		WeatherConditions:   req.WeatherConditions,
		SoilConditions:      req.SoilConditions,
		IsCompleted:         req.IsCompleted,
		ActualDosage:        req.ActualDosage,
		EffectivenessRating: int(req.EffectivenessRating),
		Notes:               req.Notes,
	}

	if req.ApplicationDate != nil {
		applicationDate, err := time.Parse(time.RFC3339, req.ApplicationDate.String())
		if err != nil {
			return nil, err
		}
		fertilizerSchedule.ApplicationDate = &applicationDate
	}

	if req.CompletedDate != nil {
		completedDate, err := time.Parse(time.RFC3339, req.CompletedDate.String())
		if err != nil {
			return nil, err
		}
		fertilizerSchedule.CompletedDate = &completedDate
	}

	return fertilizerSchedule, nil
}
