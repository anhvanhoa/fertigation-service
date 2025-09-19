package fertilizer_schedule_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	"github.com/anhvanhoa/service-core/common"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *FertilizerScheduleService) ListFertilizerSchedules(ctx context.Context, req *proto_fertilizer_schedule.FilterFertilizerSchedulesRequest) (*proto_fertilizer_schedule.ListFertilizerSchedulesResponse, error) {
	filter := s.createEntityFertilizerScheduleFilter(req)
	response, err := s.listFertilizerScheduleUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerSchedulesResponse(response), nil
}

func (s *FertilizerScheduleService) createEntityFertilizerScheduleFilter(req *proto_fertilizer_schedule.FilterFertilizerSchedulesRequest) *entity.FertilizerScheduleFilter {
	filter := &entity.FertilizerScheduleFilter{
		PlantingCycleID:   req.PlantingCycleId,
		FertilizerTypeID:  req.FertilizerTypeId,
		ApplicationMethod: req.ApplicationMethod,
		GrowthStage:       req.GrowthStage,
		IsCompleted:       req.IsCompleted,
		CreatedBy:         req.CreatedBy,
		Page:              int(req.Page),
		Limit:             int(req.Limit),
		SortBy:            req.SortBy,
		SortOrder:         req.SortOrder,
	}

	if req.ApplicationDateFrom != nil {
		applicationDateFrom, err := time.Parse(time.RFC3339, req.ApplicationDateFrom.String())
		if err == nil {
			filter.ApplicationDateFrom = &applicationDateFrom
		}
	}

	if req.ApplicationDateTo != nil {
		applicationDateTo, err := time.Parse(time.RFC3339, req.ApplicationDateTo.String())
		if err == nil {
			filter.ApplicationDateTo = &applicationDateTo
		}
	}

	if req.CompletedDateFrom != nil {
		completedDateFrom, err := time.Parse(time.RFC3339, req.CompletedDateFrom.String())
		if err == nil {
			filter.CompletedDateFrom = &completedDateFrom
		}
	}

	if req.CompletedDateTo != nil {
		completedDateTo, err := time.Parse(time.RFC3339, req.CompletedDateTo.String())
		if err == nil {
			filter.CompletedDateTo = &completedDateTo
		}
	}

	if req.CreatedAtFrom != nil {
		createdAtFrom, err := time.Parse(time.RFC3339, req.CreatedAtFrom.String())
		if err == nil {
			filter.CreatedAtFrom = &createdAtFrom
		}
	}

	if req.CreatedAtTo != nil {
		createdAtTo, err := time.Parse(time.RFC3339, req.CreatedAtTo.String())
		if err == nil {
			filter.CreatedAtTo = &createdAtTo
		}
	}

	return filter
}

func (s *FertilizerScheduleService) createProtoListFertilizerSchedulesResponse(response *common.PaginationResult[*entity.FertilizerSchedule]) *proto_fertilizer_schedule.ListFertilizerSchedulesResponse {
	protoSchedules := make([]*proto_fertilizer_schedule.FertilizerScheduleResponse, len(response.Data))
	for i, schedule := range response.Data {
		protoSchedules[i] = s.createProtoFertilizerScheduleFromResponse(schedule)
	}

	return &proto_fertilizer_schedule.ListFertilizerSchedulesResponse{
		FertilizerSchedules: protoSchedules,
		Total:               int32(response.Total),
		Page:                int32(response.Page),
		TotalPages:          int32(response.TotalPages),
	}
}

func (s *FertilizerScheduleService) createProtoFertilizerScheduleFromResponse(schedule *entity.FertilizerSchedule) *proto_fertilizer_schedule.FertilizerScheduleResponse {
	response := &proto_fertilizer_schedule.FertilizerScheduleResponse{
		Id:                  schedule.ID,
		PlantingCycleId:     schedule.PlantingCycleID,
		FertilizerTypeId:    schedule.FertilizerTypeID,
		Dosage:              schedule.Dosage,
		Unit:                schedule.Unit,
		ApplicationMethod:   schedule.ApplicationMethod,
		GrowthStage:         schedule.GrowthStage,
		WeatherConditions:   schedule.WeatherConditions,
		SoilConditions:      schedule.SoilConditions,
		IsCompleted:         schedule.IsCompleted,
		ActualDosage:        schedule.ActualDosage,
		EffectivenessRating: int32(schedule.EffectivenessRating),
		Notes:               schedule.Notes,
		CreatedBy:           schedule.CreatedBy,
	}

	if schedule.ApplicationDate != nil {
		response.ApplicationDate = timestamppb.New(*schedule.ApplicationDate)
	}

	if schedule.CompletedDate != nil {
		response.CompletedDate = timestamppb.New(*schedule.CompletedDate)
	}

	return response
}
