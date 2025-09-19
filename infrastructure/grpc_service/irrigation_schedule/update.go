package irrigation_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	irrigationScheduleP "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
)

func (s *IrrigationScheduleService) UpdateIrrigationSchedule(ctx context.Context, req *irrigationScheduleP.UpdateIrrigationScheduleRequest) (*irrigationScheduleP.UpdateIrrigationScheduleResponse, error) {
	irrigationScheduleReq, err := s.createEntityUpdateIrrigationScheduleReq(req)
	if err != nil {
		return nil, err
	}
	irrigationSchedule, err := s.updateIrrigationScheduleUsecase.Execute(ctx, irrigationScheduleReq)
	if err != nil {
		return nil, err
	}
	return &irrigationScheduleP.UpdateIrrigationScheduleResponse{
		Success:            true,
		Message:            "Irrigation schedule updated successfully",
		IrrigationSchedule: s.createProtoIrrigationSchedule(irrigationSchedule),
	}, nil
}

func (s *IrrigationScheduleService) createEntityUpdateIrrigationScheduleReq(req *irrigationScheduleP.UpdateIrrigationScheduleRequest) (*entity.UpdateIrrigationScheduleRequest, error) {
	irrigationSchedule := &entity.UpdateIrrigationScheduleRequest{
		ID:                req.Id,
		GrowingZoneID:     req.GrowingZoneId,
		PlantingCycleId:   req.PlantingCycleId,
		ScheduleName:      req.ScheduleName,
		IrrigationType:    req.IrrigationType,
		StartTime:         req.StartTime,
		DurationMinutes:   int(req.DurationMinutes),
		Frequency:         req.Frequency,
		DaysOfWeek:        req.DaysOfWeek,
		WaterAmountLiters: req.WaterAmountLiters,
		FertilizerMix:     req.FertilizerMix,
		IsActive:          req.IsActive,
	}

	if req.LastExecuted != nil {
		lastExecuted, err := time.Parse(time.RFC3339, req.LastExecuted.String())
		if err != nil {
			return nil, err
		}
		irrigationSchedule.LastExecuted = &lastExecuted
	}

	if req.NextExecution != nil {
		nextExecution, err := time.Parse(time.RFC3339, req.NextExecution.String())
		if err != nil {
			return nil, err
		}
		irrigationSchedule.NextExecution = &nextExecution
	}

	return irrigationSchedule, nil
}
