package fertilizer_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *FertilizerScheduleService) CreateFertilizerSchedule(ctx context.Context, req *proto_fertilizer_schedule.CreateFertilizerScheduleRequest) (*proto_fertilizer_schedule.FertilizerScheduleResponse, error) {
	fertilizerScheduleReq, err := s.createEntityFertilizerScheduleReq(req)
	if err != nil {
		return nil, err
	}
	fertilizerSchedule, err := s.createFertilizerScheduleUsecase.Execute(ctx, fertilizerScheduleReq)
	if err != nil {
		return nil, err
	}
	return s.createProtoFertilizerSchedule(fertilizerSchedule), nil
}

func (s *FertilizerScheduleService) createEntityFertilizerScheduleReq(req *proto_fertilizer_schedule.CreateFertilizerScheduleRequest) (*entity.CreateFertilizerScheduleRequest, error) {
	fertilizerSchedule := &entity.CreateFertilizerScheduleRequest{
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
		CreatedBy:           req.CreatedBy,
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

func (s *FertilizerScheduleService) createProtoFertilizerSchedule(fertilizerSchedule *entity.FertilizerSchedule) *proto_fertilizer_schedule.FertilizerScheduleResponse {
	response := &proto_fertilizer_schedule.FertilizerScheduleResponse{
		Id:                  fertilizerSchedule.ID,
		PlantingCycleId:     fertilizerSchedule.PlantingCycleID,
		FertilizerTypeId:    fertilizerSchedule.FertilizerTypeID,
		Dosage:              fertilizerSchedule.Dosage,
		Unit:                fertilizerSchedule.Unit,
		ApplicationMethod:   fertilizerSchedule.ApplicationMethod,
		GrowthStage:         fertilizerSchedule.GrowthStage,
		WeatherConditions:   fertilizerSchedule.WeatherConditions,
		SoilConditions:      fertilizerSchedule.SoilConditions,
		IsCompleted:         fertilizerSchedule.IsCompleted,
		ActualDosage:        fertilizerSchedule.ActualDosage,
		EffectivenessRating: int32(fertilizerSchedule.EffectivenessRating),
		Notes:               fertilizerSchedule.Notes,
		CreatedBy:           fertilizerSchedule.CreatedBy,
	}

	if fertilizerSchedule.ApplicationDate != nil {
		response.ApplicationDate = timestamppb.New(*fertilizerSchedule.ApplicationDate)
	}

	if fertilizerSchedule.CompletedDate != nil {
		response.CompletedDate = timestamppb.New(*fertilizerSchedule.CompletedDate)
	}

	return response
}
