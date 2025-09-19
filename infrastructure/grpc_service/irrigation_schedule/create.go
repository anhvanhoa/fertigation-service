package irrigation_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IrrigationScheduleService) CreateIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.CreateIrrigationScheduleRequest) (*irrigationScheduleP.CreateIrrigationScheduleResponse, error) {
	irrigationScheduleReq, err := s.createEntityIrrigationScheduleReq(req)
	if err != nil {
		return nil, err
	}
	irrigationSchedule, err := s.createIrrigationScheduleUsecase.Execute(ctx, irrigationScheduleReq)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.CreateIrrigationScheduleResponse{
		Success:            true,
		Message:            "Irrigation schedule created successfully",
		IrrigationSchedule: s.createProtoIrrigationSchedule(irrigationSchedule),
	}, nil
}

func (s *IrrigationScheduleService) createEntityIrrigationScheduleReq(req *irrigationScheduleP.CreateIrrigationScheduleRequest) (*entity.CreateIrrigationScheduleRequest, error) {
	irrigationSchedule := &entity.CreateIrrigationScheduleRequest{
		GrowingZoneID:     req.GrowingZoneId,
		PlantingCycleID:   req.PlantingCycleId,
		ScheduleName:      req.ScheduleName,
		IrrigationType:    req.IrrigationType,
		StartTime:         req.StartTime,
		DurationMinutes:   int(req.DurationMinutes),
		Frequency:         req.Frequency,
		DaysOfWeek:        req.DaysOfWeek,
		WaterAmountLiters: req.WaterAmountLiters,
		FertilizerMix:     req.FertilizerMix,
		IsActive:          req.IsActive,
		CreatedBy:         req.CreatedBy,
	}

	return irrigationSchedule, nil
}

func (s *IrrigationScheduleService) createProtoIrrigationSchedule(irrigationSchedule *entity.IrrigationSchedule) *irrigationScheduleP.IrrigationSchedule {
	response := &irrigationScheduleP.IrrigationSchedule{
		Id:                irrigationSchedule.ID,
		GrowingZoneId:     irrigationSchedule.GrowingZoneID,
		PlantingCycleId:   irrigationSchedule.PlantingCycleID,
		ScheduleName:      irrigationSchedule.ScheduleName,
		IrrigationType:    irrigationSchedule.IrrigationType,
		StartTime:         irrigationSchedule.StartTime,
		DurationMinutes:   int32(irrigationSchedule.DurationMinutes),
		Frequency:         irrigationSchedule.Frequency,
		DaysOfWeek:        irrigationSchedule.DaysOfWeek,
		WaterAmountLiters: irrigationSchedule.WaterAmountLiters,
		FertilizerMix:     irrigationSchedule.FertilizerMix,
		IsActive:          irrigationSchedule.IsActive,
		CreatedBy:         irrigationSchedule.CreatedBy,
	}

	if irrigationSchedule.LastExecuted != nil {
		response.LastExecuted = timestamppb.New(*irrigationSchedule.LastExecuted)
	}

	if irrigationSchedule.NextExecution != nil {
		response.NextExecution = timestamppb.New(*irrigationSchedule.NextExecution)
	}

	return response
}
