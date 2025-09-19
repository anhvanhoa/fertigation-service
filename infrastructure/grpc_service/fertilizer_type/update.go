package fertilizer_type_service

import (
	"context"
	"fertigation-Service/domain/entity"
	"time"

	fertilizerTypeP "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
)

func (s *FertilizerTypeService) UpdateFertilizerType(ctx context.Context, req *fertilizerTypeP.UpdateFertilizerTypeRequest) (*fertilizerTypeP.UpdateFertilizerTypeResponse, error) {
	fertilizerTypeReq, err := s.createEntityUpdateFertilizerTypeReq(req)
	if err != nil {
		return nil, err
	}
	fertilizerType, err := s.updateFertilizerTypeUsecase.Execute(ctx, fertilizerTypeReq)
	if err != nil {
		return nil, err
	}
	return &fertilizerTypeP.UpdateFertilizerTypeResponse{
		Success:        true,
		Message:        "Fertilizer type updated successfully",
		FertilizerType: s.createProtoFertilizerType(fertilizerType),
	}, nil
}

func (s *FertilizerTypeService) createEntityUpdateFertilizerTypeReq(req *fertilizerTypeP.UpdateFertilizerTypeRequest) (*entity.UpdateFertilizerTypeRequest, error) {
	fertilizerType := &entity.UpdateFertilizerTypeRequest{
		ID:                   req.Id,
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
