package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	proto_fertilizer_type "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	proto_irrigation_log "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	proto_irrigation_schedule "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter page size (default 10): ")
	pageSizeStr, _ := reader.ReadString('\n')
	pageSizeStr = cleanInput(pageSizeStr)
	pageSize := int32(10)
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pageSize = int32(ps)
		}
	}

	return page, pageSize
}

type FertigationServiceClient struct {
	fertilizerScheduleClient proto_fertilizer_schedule.FertilizerScheduleServiceClient
	fertilizerTypeClient     proto_fertilizer_type.FertilizerTypeServiceClient
	irrigationLogClient      proto_irrigation_log.IrrigationLogServiceClient
	irrigationScheduleClient proto_irrigation_schedule.IrrigationScheduleServiceClient
	conn                     *grpc.ClientConn
}

func NewFertigationServiceClient(address string) (*FertigationServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &FertigationServiceClient{
		fertilizerScheduleClient: proto_fertilizer_schedule.NewFertilizerScheduleServiceClient(conn),
		fertilizerTypeClient:     proto_fertilizer_type.NewFertilizerTypeServiceClient(conn),
		irrigationLogClient:      proto_irrigation_log.NewIrrigationLogServiceClient(conn),
		irrigationScheduleClient: proto_irrigation_schedule.NewIrrigationScheduleServiceClient(conn),
		conn:                     conn,
	}, nil
}

func (c *FertigationServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Fertilizer Schedule Service Tests ==================

func (c *FertigationServiceClient) TestCreateFertilizerSchedule() {
	fmt.Println("\n=== Test Create Fertilizer Schedule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Enter fertilizer type ID: ")
	fertilizerTypeId, _ := reader.ReadString('\n')
	fertilizerTypeId = cleanInput(fertilizerTypeId)

	fmt.Print("Enter application date (YYYY-MM-DD): ")
	appDateStr, _ := reader.ReadString('\n')
	appDateStr = cleanInput(appDateStr)
	var appDate *timestamppb.Timestamp
	if appDateStr != "" {
		if t, err := time.Parse("2006-01-02", appDateStr); err == nil {
			appDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter dosage: ")
	dosageStr, _ := reader.ReadString('\n')
	dosageStr = cleanInput(dosageStr)
	dosage := 0.0
	if dosageStr != "" {
		if d, err := strconv.ParseFloat(dosageStr, 64); err == nil {
			dosage = d
		}
	}

	fmt.Print("Enter unit: ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Enter application method (foliar/soil/hydroponic/fertigation): ")
	appMethod, _ := reader.ReadString('\n')
	appMethod = cleanInput(appMethod)

	fmt.Print("Enter growth stage (seedling/vegetative/flowering/fruiting/pre_harvest): ")
	growthStage, _ := reader.ReadString('\n')
	growthStage = cleanInput(growthStage)

	fmt.Print("Enter weather conditions: ")
	weather, _ := reader.ReadString('\n')
	weather = cleanInput(weather)

	fmt.Print("Enter soil conditions: ")
	soil, _ := reader.ReadString('\n')
	soil = cleanInput(soil)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerScheduleClient.CreateFertilizerSchedule(ctx, &proto_fertilizer_schedule.CreateFertilizerScheduleRequest{
		PlantingCycleId:   plantingCycleId,
		FertilizerTypeId:  fertilizerTypeId,
		ApplicationDate:   appDate,
		Dosage:            dosage,
		Unit:              unit,
		ApplicationMethod: appMethod,
		GrowthStage:       growthStage,
		WeatherConditions: weather,
		SoilConditions:    soil,
		IsCompleted:       false,
		CreatedBy:         createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateFertilizerSchedule: %v\n", err)
		return
	}

	fmt.Printf("Create Fertilizer Schedule result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PlantingCycleId)
		fmt.Printf("Fertilizer Type ID: %s\n", resp.FertilizerTypeId)
		fmt.Printf("Dosage: %.2f %s\n", resp.Dosage, resp.Unit)
		fmt.Printf("Application Method: %s\n", resp.ApplicationMethod)
		fmt.Printf("Growth Stage: %s\n", resp.GrowthStage)
	}
}

func (c *FertigationServiceClient) TestGetFertilizerSchedule() {
	fmt.Println("\n=== Test Get Fertilizer Schedule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter fertilizer schedule ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerScheduleClient.GetFertilizerSchedule(ctx, &proto_fertilizer_schedule.GetFertilizerScheduleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetFertilizerSchedule: %v\n", err)
		return
	}

	fmt.Printf("Get Fertilizer Schedule result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PlantingCycleId)
		fmt.Printf("Fertilizer Type ID: %s\n", resp.FertilizerTypeId)
		fmt.Printf("Dosage: %.2f %s\n", resp.Dosage, resp.Unit)
		fmt.Printf("Application Method: %s\n", resp.ApplicationMethod)
		fmt.Printf("Growth Stage: %s\n", resp.GrowthStage)
		fmt.Printf("Is Completed: %t\n", resp.IsCompleted)
	}
}

func (c *FertigationServiceClient) TestListFertilizerSchedules() {
	fmt.Println("\n=== Test List Fertilizer Schedules ===")

	reader := bufio.NewReader(os.Stdin)

	page, pageSize := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerScheduleClient.ListFertilizerSchedules(ctx, &proto_fertilizer_schedule.FilterFertilizerSchedulesRequest{
		Page:      page,
		Limit:     pageSize,
		SortBy:    "created_at",
		SortOrder: "desc",
	})
	if err != nil {
		fmt.Printf("Error calling ListFertilizerSchedules: %v\n", err)
		return
	}

	fmt.Printf("List Fertilizer Schedules result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Fertilizer Schedules:\n")
	for i, schedule := range resp.FertilizerSchedules {
		fmt.Printf("  [%d] ID: %s, Planting Cycle: %s, Dosage: %.2f %s\n",
			i+1, schedule.Id, schedule.PlantingCycleId, schedule.Dosage, schedule.Unit)
	}
}

func (c *FertigationServiceClient) TestGetSchedulesByPlantingCycle() {
	fmt.Println("\n=== Test Get Schedules By Planting Cycle ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	page, pageSize := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerScheduleClient.GetSchedulesByPlantingCycle(ctx, &proto_fertilizer_schedule.GetSchedulesByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_fertilizer_schedule.Pagination{
			Page:     page,
			PageSize: pageSize,
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetSchedulesByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Get Schedules By Planting Cycle result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Fertilizer Schedules:\n")
	for i, schedule := range resp.FertilizerSchedules {
		fmt.Printf("  [%d] ID: %s, Dosage: %.2f %s, Method: %s\n",
			i+1, schedule.Id, schedule.Dosage, schedule.Unit, schedule.ApplicationMethod)
	}
}

// ================== Fertilizer Type Service Tests ==================

func (c *FertigationServiceClient) TestCreateFertilizerType() {
	fmt.Println("\n=== Test Create Fertilizer Type ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter fertilizer name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter type (organic/chemical/liquid/granular/powder): ")
	fertilizerType, _ := reader.ReadString('\n')
	fertilizerType = cleanInput(fertilizerType)

	fmt.Print("Enter NPK ratio: ")
	npkRatio, _ := reader.ReadString('\n')
	npkRatio = cleanInput(npkRatio)

	fmt.Print("Enter nitrogen percentage: ")
	nitrogenStr, _ := reader.ReadString('\n')
	nitrogenStr = cleanInput(nitrogenStr)
	nitrogen := 0.0
	if nitrogenStr != "" {
		if n, err := strconv.ParseFloat(nitrogenStr, 64); err == nil {
			nitrogen = n
		}
	}

	fmt.Print("Enter phosphorus percentage: ")
	phosphorusStr, _ := reader.ReadString('\n')
	phosphorusStr = cleanInput(phosphorusStr)
	phosphorus := 0.0
	if phosphorusStr != "" {
		if p, err := strconv.ParseFloat(phosphorusStr, 64); err == nil {
			phosphorus = p
		}
	}

	fmt.Print("Enter potassium percentage: ")
	potassiumStr, _ := reader.ReadString('\n')
	potassiumStr = cleanInput(potassiumStr)
	potassium := 0.0
	if potassiumStr != "" {
		if k, err := strconv.ParseFloat(potassiumStr, 64); err == nil {
			potassium = k
		}
	}

	fmt.Print("Enter application method (foliar/soil/hydroponic/fertigation): ")
	appMethod, _ := reader.ReadString('\n')
	appMethod = cleanInput(appMethod)

	fmt.Print("Enter dosage per plant: ")
	dosagePlantStr, _ := reader.ReadString('\n')
	dosagePlantStr = cleanInput(dosagePlantStr)
	dosagePlant := 0.0
	if dosagePlantStr != "" {
		if d, err := strconv.ParseFloat(dosagePlantStr, 64); err == nil {
			dosagePlant = d
		}
	}

	fmt.Print("Enter unit: ")
	unit, _ := reader.ReadString('\n')
	unit = cleanInput(unit)

	fmt.Print("Enter manufacturer: ")
	manufacturer, _ := reader.ReadString('\n')
	manufacturer = cleanInput(manufacturer)

	fmt.Print("Enter expiry date (YYYY-MM-DD): ")
	expiryStr, _ := reader.ReadString('\n')
	expiryStr = cleanInput(expiryStr)
	var expiryDate *timestamppb.Timestamp
	if expiryStr != "" {
		if t, err := time.Parse("2006-01-02", expiryStr); err == nil {
			expiryDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerTypeClient.CreateFertilizerType(ctx, &proto_fertilizer_type.CreateFertilizerTypeRequest{
		Name:                 name,
		Type:                 fertilizerType,
		NpkRatio:             npkRatio,
		NitrogenPercentage:   nitrogen,
		PhosphorusPercentage: phosphorus,
		PotassiumPercentage:  potassium,
		ApplicationMethod:    appMethod,
		DosagePerPlant:       dosagePlant,
		Unit:                 unit,
		Manufacturer:         manufacturer,
		ExpiryDate:           expiryDate,
		Status:               "active",
		CreatedBy:            createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateFertilizerType: %v\n", err)
		return
	}

	fmt.Printf("Create Fertilizer Type result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Name: %s\n", resp.Name)
		fmt.Printf("Type: %s\n", resp.Type)
		fmt.Printf("NPK Ratio: %s\n", resp.NpkRatio)
		fmt.Printf("N: %.1f%%, P: %.1f%%, K: %.1f%%\n", resp.NitrogenPercentage, resp.PhosphorusPercentage, resp.PotassiumPercentage)
		fmt.Printf("Manufacturer: %s\n", resp.Manufacturer)
	}
}

func (c *FertigationServiceClient) TestGetFertilizerType() {
	fmt.Println("\n=== Test Get Fertilizer Type ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter fertilizer type ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerTypeClient.GetFertilizerType(ctx, &proto_fertilizer_type.GetFertilizerTypeRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetFertilizerType: %v\n", err)
		return
	}

	fmt.Printf("Get Fertilizer Type result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Name: %s\n", resp.Name)
		fmt.Printf("Type: %s\n", resp.Type)
		fmt.Printf("NPK Ratio: %s\n", resp.NpkRatio)
		fmt.Printf("N: %.1f%%, P: %.1f%%, K: %.1f%%\n", resp.NitrogenPercentage, resp.PhosphorusPercentage, resp.PotassiumPercentage)
		fmt.Printf("Manufacturer: %s\n", resp.Manufacturer)
		fmt.Printf("Status: %s\n", resp.Status)
	}
}

func (c *FertigationServiceClient) TestListFertilizerTypes() {
	fmt.Println("\n=== Test List Fertilizer Types ===")

	reader := bufio.NewReader(os.Stdin)

	page, pageSize := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.fertilizerTypeClient.ListFertilizerTypes(ctx, &proto_fertilizer_type.ListFertilizerTypesRequest{
		Pagination: &proto_fertilizer_type.Pagination{
			Page:     page,
			PageSize: pageSize,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListFertilizerTypes: %v\n", err)
		return
	}

	fmt.Printf("List Fertilizer Types result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Fertilizer Types:\n")
	for i, fertilizerType := range resp.FertilizerTypes {
		fmt.Printf("  [%d] ID: %s, Name: %s, Type: %s, NPK: %s\n",
			i+1, fertilizerType.Id, fertilizerType.Name, fertilizerType.Type, fertilizerType.NpkRatio)
	}
}

// ================== Irrigation Log Service Tests ==================

func (c *FertigationServiceClient) TestCreateIrrigationLog() {
	fmt.Println("\n=== Test Create Irrigation Log ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter irrigation schedule ID: ")
	scheduleId, _ := reader.ReadString('\n')
	scheduleId = cleanInput(scheduleId)

	fmt.Print("Enter device ID: ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Enter started at (YYYY-MM-DD HH:MM:SS): ")
	startedStr, _ := reader.ReadString('\n')
	startedStr = cleanInput(startedStr)
	var startedAt *timestamppb.Timestamp
	if startedStr != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", startedStr); err == nil {
			startedAt = timestamppb.New(t)
		}
	}

	fmt.Print("Enter ended at (YYYY-MM-DD HH:MM:SS): ")
	endedStr, _ := reader.ReadString('\n')
	endedStr = cleanInput(endedStr)
	var endedAt *timestamppb.Timestamp
	if endedStr != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", endedStr); err == nil {
			endedAt = timestamppb.New(t)
		}
	}

	fmt.Print("Enter planned duration (minutes): ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := int32(0)
	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil {
			duration = int32(d)
		}
	}

	fmt.Print("Enter water used (liters): ")
	waterStr, _ := reader.ReadString('\n')
	waterStr = cleanInput(waterStr)
	water := 0.0
	if waterStr != "" {
		if w, err := strconv.ParseFloat(waterStr, 64); err == nil {
			water = w
		}
	}

	fmt.Print("Enter status (pending/running/completed/failed/cancelled): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.irrigationLogClient.CreateIrrigationLog(ctx, &proto_irrigation_log.CreateIrrigationLogRequest{
		IrrigationScheduleId:   scheduleId,
		DeviceId:               deviceId,
		StartedAt:              startedAt,
		EndedAt:                endedAt,
		PlannedDurationMinutes: duration,
		WaterUsedLiters:        water,
		Status:                 status,
		CreatedBy:              createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateIrrigationLog: %v\n", err)
		return
	}

	fmt.Printf("Create Irrigation Log result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Schedule ID: %s\n", resp.IrrigationScheduleId)
		fmt.Printf("Device ID: %s\n", resp.DeviceId)
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Water Used: %.2f liters\n", resp.WaterUsedLiters)
	}
}

func (c *FertigationServiceClient) TestGetIrrigationLog() {
	fmt.Println("\n=== Test Get Irrigation Log ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter irrigation log ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.irrigationLogClient.GetIrrigationLog(ctx, &proto_irrigation_log.GetIrrigationLogRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetIrrigationLog: %v\n", err)
		return
	}

	fmt.Printf("Get Irrigation Log result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Schedule ID: %s\n", resp.IrrigationScheduleId)
		fmt.Printf("Device ID: %s\n", resp.DeviceId)
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Water Used: %.2f liters\n", resp.WaterUsedLiters)
		fmt.Printf("Planned Duration: %d minutes\n", resp.PlannedDurationMinutes)
	}
}

func (c *FertigationServiceClient) TestListIrrigationLogs() {
	fmt.Println("\n=== Test List Irrigation Logs ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter irrigation schedule ID (optional): ")
	scheduleId, _ := reader.ReadString('\n')
	scheduleId = cleanInput(scheduleId)

	page, pageSize := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_irrigation_log.ListIrrigationLogsRequest{
		Page:  int32(page),
		Limit: int32(pageSize),
	}
	if scheduleId != "" {
		req.IrrigationScheduleId = scheduleId
	}

	resp, err := c.irrigationLogClient.ListIrrigationLogs(ctx, req)
	if err != nil {
		fmt.Printf("Error calling ListIrrigationLogs: %v\n", err)
		return
	}

	fmt.Printf("List Irrigation Logs result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Irrigation Logs:\n")
	for i, log := range resp.IrrigationLogs {
		fmt.Printf("  [%d] ID: %s, Schedule: %s, Status: %s, Water: %.2fL\n",
			i+1, log.Id, log.IrrigationScheduleId, log.Status, log.WaterUsedLiters)
	}
}

// ================== Irrigation Schedule Service Tests ==================

func (c *FertigationServiceClient) TestCreateIrrigationSchedule() {
	fmt.Println("\n=== Test Create Irrigation Schedule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter growing zone ID: ")
	growingZoneId, _ := reader.ReadString('\n')
	growingZoneId = cleanInput(growingZoneId)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Enter schedule name: ")
	scheduleName, _ := reader.ReadString('\n')
	scheduleName = cleanInput(scheduleName)

	fmt.Print("Enter irrigation type: ")
	irrigationType, _ := reader.ReadString('\n')
	irrigationType = cleanInput(irrigationType)

	fmt.Print("Enter start time (HH:MM): ")
	startTime, _ := reader.ReadString('\n')
	startTime = cleanInput(startTime)

	fmt.Print("Enter duration (minutes): ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := int32(0)
	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil {
			duration = int32(d)
		}
	}

	fmt.Print("Enter frequency: ")
	frequency, _ := reader.ReadString('\n')
	frequency = cleanInput(frequency)

	fmt.Print("Enter water amount (liters): ")
	waterStr, _ := reader.ReadString('\n')
	waterStr = cleanInput(waterStr)
	water := 0.0
	if waterStr != "" {
		if w, err := strconv.ParseFloat(waterStr, 64); err == nil {
			water = w
		}
	}

	fmt.Print("Enter fertilizer mix (true/false): ")
	fertilizerMixStr, _ := reader.ReadString('\n')
	fertilizerMixStr = cleanInput(fertilizerMixStr)
	fertilizerMix := false
	if fertilizerMixStr == "true" {
		fertilizerMix = true
	}

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.irrigationScheduleClient.CreateIrrigationSchedule(ctx, &proto_irrigation_schedule.CreateIrrigationScheduleRequest{
		GrowingZoneId:     growingZoneId,
		PlantingCycleId:   plantingCycleId,
		ScheduleName:      scheduleName,
		IrrigationType:    irrigationType,
		StartTime:         startTime,
		DurationMinutes:   duration,
		Frequency:         frequency,
		WaterAmountLiters: water,
		FertilizerMix:     fertilizerMix,
		IsActive:          true,
		CreatedBy:         createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateIrrigationSchedule: %v\n", err)
		return
	}

	fmt.Printf("Create Irrigation Schedule result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Schedule Name: %s\n", resp.ScheduleName)
		fmt.Printf("Growing Zone ID: %s\n", resp.GrowingZoneId)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PlantingCycleId)
		fmt.Printf("Start Time: %s\n", resp.StartTime)
		fmt.Printf("Duration: %d minutes\n", resp.DurationMinutes)
		fmt.Printf("Water Amount: %.2f liters\n", resp.WaterAmountLiters)
	}
}

func (c *FertigationServiceClient) TestGetIrrigationSchedule() {
	fmt.Println("\n=== Test Get Irrigation Schedule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter irrigation schedule ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.irrigationScheduleClient.GetIrrigationSchedule(ctx, &proto_irrigation_schedule.GetIrrigationScheduleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetIrrigationSchedule: %v\n", err)
		return
	}

	fmt.Printf("Get Irrigation Schedule result:\n")
	if resp != nil {
		fmt.Printf("ID: %s\n", resp.Id)
		fmt.Printf("Schedule Name: %s\n", resp.ScheduleName)
		fmt.Printf("Growing Zone ID: %s\n", resp.GrowingZoneId)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PlantingCycleId)
		fmt.Printf("Start Time: %s\n", resp.StartTime)
		fmt.Printf("Duration: %d minutes\n", resp.DurationMinutes)
		fmt.Printf("Water Amount: %.2f liters\n", resp.WaterAmountLiters)
		fmt.Printf("Is Active: %t\n", resp.IsActive)
	}
}

func (c *FertigationServiceClient) TestListIrrigationSchedules() {
	fmt.Println("\n=== Test List Irrigation Schedules ===")

	reader := bufio.NewReader(os.Stdin)

	page, pageSize := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.irrigationScheduleClient.ListIrrigationSchedules(ctx, &proto_irrigation_schedule.ListIrrigationSchedulesRequest{
		Page:  int32(page),
		Limit: int32(pageSize),
	})
	if err != nil {
		fmt.Printf("Error calling ListIrrigationSchedules: %v\n", err)
		return
	}

	fmt.Printf("List Irrigation Schedules result:\n")
	fmt.Printf("Total: %d\n", resp.Pagination.Total)
	fmt.Printf("Irrigation Schedules:\n")
	for i, schedule := range resp.IrrigationSchedules {
		fmt.Printf("  [%d] ID: %s, Name: %s, Zone: %s, Water: %.2fL\n",
			i+1, schedule.Id, schedule.ScheduleName, schedule.GrowingZoneId, schedule.WaterAmountLiters)
	}
}

func (c *FertigationServiceClient) TestGetActiveSchedules() {
	fmt.Println("\n=== Test Get Active Schedules ===")
	fmt.Println("Note: This method requires the common package for PaginationRequest")
	fmt.Println("Skipping for now - would need proper import of common/v1/common.proto")

	// TODO: Implement when common package is available
	// This would require importing: "github.com/anhvanhoa/sf-proto/gen/common/v1"
	// and using: &proto_common.PaginationRequest{Page: 1, PageSize: 10}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Fertigation Service Test Client ===")
	fmt.Println("1. Fertilizer Schedule Service")
	fmt.Println("2. Fertilizer Type Service")
	fmt.Println("3. Irrigation Log Service")
	fmt.Println("4. Irrigation Schedule Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printFertilizerScheduleMenu() {
	fmt.Println("\n=== Fertilizer Schedule Service ===")
	fmt.Println("1. Create Fertilizer Schedule")
	fmt.Println("2. Get Fertilizer Schedule")
	fmt.Println("3. List Fertilizer Schedules")
	fmt.Println("4. Get Schedules By Planting Cycle")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printFertilizerTypeMenu() {
	fmt.Println("\n=== Fertilizer Type Service ===")
	fmt.Println("1. Create Fertilizer Type")
	fmt.Println("2. Get Fertilizer Type")
	fmt.Println("3. List Fertilizer Types")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printIrrigationLogMenu() {
	fmt.Println("\n=== Irrigation Log Service ===")
	fmt.Println("1. Create Irrigation Log")
	fmt.Println("2. Get Irrigation Log")
	fmt.Println("3. List Irrigation Logs")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printIrrigationScheduleMenu() {
	fmt.Println("\n=== Irrigation Schedule Service ===")
	fmt.Println("1. Create Irrigation Schedule")
	fmt.Println("2. Get Irrigation Schedule")
	fmt.Println("3. List Irrigation Schedules")
	fmt.Println("4. Get Active Schedules")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewFertigationServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Fertilizer Schedule Service
			for {
				printFertilizerScheduleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateFertilizerSchedule()
				case "2":
					client.TestGetFertilizerSchedule()
				case "3":
					client.TestListFertilizerSchedules()
				case "4":
					client.TestGetSchedulesByPlantingCycle()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Fertilizer Type Service
			for {
				printFertilizerTypeMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateFertilizerType()
				case "2":
					client.TestGetFertilizerType()
				case "3":
					client.TestListFertilizerTypes()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "3":
			// Irrigation Log Service
			for {
				printIrrigationLogMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateIrrigationLog()
				case "2":
					client.TestGetIrrigationLog()
				case "3":
					client.TestListIrrigationLogs()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "4":
			// Irrigation Schedule Service
			for {
				printIrrigationScheduleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateIrrigationSchedule()
				case "2":
					client.TestGetIrrigationSchedule()
				case "3":
					client.TestListIrrigationSchedules()
				case "4":
					client.TestGetActiveSchedules()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
