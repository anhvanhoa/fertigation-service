package fertilizer_type_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *FertilizerTypeService) ListFertilizerTypes(ctx context.Context, req *fertilizerTypeP.ListFertilizerTypesRequest) (*fertilizerTypeP.ListFertilizerTypesResponse, error) {
	filter := s.createEntityFertilizerTypeFilter(req)
	response, err := s.listFertilizerTypeUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}
	return s.createProtoListFertilizerTypesResponse(response), nil
}

func (s *FertilizerTypeService) createEntityFertilizerTypeFilter(req *fertilizerTypeP.ListFertilizerTypesRequest) *entity.FertilizerTypeFilter {
	filter := &entity.FertilizerTypeFilter{
		Name:              req.Name,
		Type:              req.Type,
		ApplicationMethod: req.ApplicationMethod,
		Status:            req.Status,
		Manufacturer:      req.Manufacturer,
		CreatedBy:         req.CreatedBy,
		Page:              int(req.Pagination.Page),
		PageSize:          int(req.Pagination.PageSize),
		SortBy:            req.Pagination.SortBy,
		SortOrder:         req.Pagination.SortOrder,
	}

	if req.ExpiryDateFrom != nil {
		expiryDateFrom, err := time.Parse(time.RFC3339, req.ExpiryDateFrom.String())
		if err == nil {
			filter.ExpiryDateFrom = &expiryDateFrom
		}
	}

	if req.ExpiryDateTo != nil {
		expiryDateTo, err := time.Parse(time.RFC3339, req.ExpiryDateTo.String())
		if err == nil {
			filter.ExpiryDateTo = &expiryDateTo
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

func (s *FertilizerTypeService) createProtoListFertilizerTypesResponse(response *entity.ListFertilizerTypesResponse) *fertilizerTypeP.ListFertilizerTypesResponse {
	protoTypes := make([]*fertilizerTypeP.FertilizerTypeResponse, len(response.FertilizerTypes))
	for i, fertilizerType := range response.FertilizerTypes {
		protoTypes[i] = s.createProtoFertilizerTypeFromResponse(&fertilizerType)
	}

	return &fertilizerTypeP.ListFertilizerTypesResponse{
		FertilizerTypes: protoTypes,
		Total:           int32(response.Total),
		Page:            int32(response.Page),
		PageSize:        int32(response.PageSize),
		TotalPages:      int32(response.TotalPages),
	}
}

func (s *FertilizerTypeService) createProtoFertilizerTypeFromResponse(fertilizerType *entity.FertilizerTypeResponse) *fertilizerTypeP.FertilizerType {
	response := &fertilizerTypeP.FertilizerType{
		Id:                   fertilizerType.ID,
		Name:                 fertilizerType.Name,
		Type:                 fertilizerType.Type,
		NpkRatio:             fertilizerType.NPKRatio,
		NitrogenPercentage:   fertilizerType.NitrogenPercentage,
		PhosphorusPercentage: fertilizerType.PhosphorusPercentage,
		PotassiumPercentage:  fertilizerType.PotassiumPercentage,
		TraceElements:        fertilizerType.TraceElements,
		ApplicationMethod:    fertilizerType.ApplicationMethod,
		DosagePerPlant:       fertilizerType.DosagePerPlant,
		DosagePerM2:          fertilizerType.DosagePerM2,
		Unit:                 fertilizerType.Unit,
		Manufacturer:         fertilizerType.Manufacturer,
		BatchNumber:          fertilizerType.BatchNumber,
		CostPerUnit:          fertilizerType.CostPerUnit,
		Description:          fertilizerType.Description,
		SafetyNotes:          fertilizerType.SafetyNotes,
		Status:               fertilizerType.Status,
		CreatedBy:            fertilizerType.CreatedBy,
	}

	if fertilizerType.ExpiryDate != nil {
		response.ExpiryDate = timestamppb.New(*fertilizerType.ExpiryDate)
	}

	return response
}
