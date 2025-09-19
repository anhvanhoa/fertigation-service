package fertilizer_type_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *FertilizerTypeService) CreateFertilizerType(ctx context.Context, req *fertilizerTypeP.CreateFertilizerTypeRequest) (*fertilizerTypeP.FertilizerTypeResponse, error) {
	fertilizerTypeReq, err := s.createEntityFertilizerTypeReq(req)
	if err != nil {
		return nil, err
	}
	fertilizerType, err := s.createFertilizerTypeUsecase.Execute(ctx, fertilizerTypeReq)
	if err != nil {
		return nil, err
	}
	return s.createProtoFertilizerType(fertilizerType), nil
}

func (s *FertilizerTypeService) createEntityFertilizerTypeReq(req *fertilizerTypeP.CreateFertilizerTypeRequest) (*entity.CreateFertilizerTypeRequest, error) {
	fertilizerType := &entity.CreateFertilizerTypeRequest{
		Name:                 req.Name,
		Type:                 req.Type,
		NPKRatio:             req.NpkRatio,
		NitrogenPercentage:   req.NitrogenPercentage,
		PhosphorusPercentage: req.PhosphorusPercentage,
		PotassiumPercentage:  req.PotassiumPercentage,
		TraceElements:        req.TraceElements,
		ApplicationMethod:    req.ApplicationMethod,
		DosagePerPlant:       req.DosagePerPlant,
		DosagePerM2:          req.DosagePerM2,
		Unit:                 req.Unit,
		Manufacturer:         req.Manufacturer,
		BatchNumber:          req.BatchNumber,
		CostPerUnit:          req.CostPerUnit,
		Description:          req.Description,
		SafetyNotes:          req.SafetyNotes,
		Status:               req.Status,
		CreatedBy:            req.CreatedBy,
	}

	if req.ExpiryDate != nil {
		expiryDate, err := time.Parse(time.RFC3339, req.ExpiryDate.String())
		if err != nil {
			return nil, err
		}
		fertilizerType.ExpiryDate = &expiryDate
	}

	return fertilizerType, nil
}

func (s *FertilizerTypeService) createProtoFertilizerType(fertilizerType *entity.FertilizerType) *fertilizerTypeP.FertilizerTypeResponse {
	response := &fertilizerTypeP.FertilizerTypeResponse{
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
