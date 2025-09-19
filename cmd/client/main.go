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

	proto_plant_variety "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	proto_planting_cycle "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Enter offset (default 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type CropServiceClient struct {
	plantVarietyClient  proto_plant_variety.PlantVarietyServiceClient
	plantingCycleClient proto_planting_cycle.PlantingCycleServiceClient
	conn                *grpc.ClientConn
}

func NewCropServiceClient(address string) (*CropServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &CropServiceClient{
		plantVarietyClient:  proto_plant_variety.NewPlantVarietyServiceClient(conn),
		plantingCycleClient: proto_planting_cycle.NewPlantingCycleServiceClient(conn),
		conn:                conn,
	}, nil
}

func (c *CropServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Plant Variety Service Tests ==================

func (c *CropServiceClient) TestCreatePlantVariety() {
	fmt.Println("\n=== Test Create Plant Variety ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plant variety name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter scientific name: ")
	scientificName, _ := reader.ReadString('\n')
	scientificName = cleanInput(scientificName)

	fmt.Print("Enter category: ")
	category, _ := reader.ReadString('\n')
	category = cleanInput(category)

	fmt.Print("Enter growing season: ")
	growingSeason, _ := reader.ReadString('\n')
	growingSeason = cleanInput(growingSeason)

	fmt.Print("Enter growth duration days: ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := int32(90)
	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil {
			duration = int32(d)
		}
	}

	fmt.Print("Enter optimal temp min: ")
	tempMinStr, _ := reader.ReadString('\n')
	tempMinStr = cleanInput(tempMinStr)
	tempMin := float64(20.0)
	if tempMinStr != "" {
		if t, err := strconv.ParseFloat(tempMinStr, 64); err == nil {
			tempMin = t
		}
	}

	fmt.Print("Enter optimal temp max: ")
	tempMaxStr, _ := reader.ReadString('\n')
	tempMaxStr = cleanInput(tempMaxStr)
	tempMax := float64(30.0)
	if tempMaxStr != "" {
		if t, err := strconv.ParseFloat(tempMaxStr, 64); err == nil {
			tempMax = t
		}
	}

	fmt.Print("Enter water requirement: ")
	waterReq, _ := reader.ReadString('\n')
	waterReq = cleanInput(waterReq)

	fmt.Print("Enter light requirement: ")
	lightReq, _ := reader.ReadString('\n')
	lightReq = cleanInput(lightReq)

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.CreatePlantVariety(ctx, &proto_plant_variety.CreatePlantVarietyRequest{
		Name:               name,
		ScientificName:     scientificName,
		Category:           category,
		GrowingSeason:      growingSeason,
		GrowthDurationDays: duration,
		OptimalTempMin:     tempMin,
		OptimalTempMax:     tempMax,
		OptimalHumidityMin: 50.0,
		OptimalHumidityMax: 80.0,
		PhMin:              6.0,
		PhMax:              7.0,
		WaterRequirement:   waterReq,
		LightRequirement:   lightReq,
		Description:        description,
		CreatedBy:          createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreatePlantVariety: %v\n", err)
		return
	}

	fmt.Printf("Create Plant Variety result:\n")
	if resp.PlantVariety != nil {
		fmt.Printf("ID: %s\n", resp.PlantVariety.Id)
		fmt.Printf("Name: %s\n", resp.PlantVariety.Name)
		fmt.Printf("Scientific Name: %s\n", resp.PlantVariety.ScientificName)
		fmt.Printf("Category: %s\n", resp.PlantVariety.Category)
		fmt.Printf("Growing Season: %s\n", resp.PlantVariety.GrowingSeason)
	}
}

func (c *CropServiceClient) TestGetPlantVariety() {
	fmt.Println("\n=== Test Get Plant Variety ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plant variety ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVariety(ctx, &proto_plant_variety.GetPlantVarietyRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVariety: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Variety result:\n")
	if resp.PlantVariety != nil {
		fmt.Printf("ID: %s\n", resp.PlantVariety.Id)
		fmt.Printf("Name: %s\n", resp.PlantVariety.Name)
		fmt.Printf("Scientific Name: %s\n", resp.PlantVariety.ScientificName)
		fmt.Printf("Category: %s\n", resp.PlantVariety.Category)
		fmt.Printf("Growing Season: %s\n", resp.PlantVariety.GrowingSeason)
		fmt.Printf("Growth Duration: %d days\n", resp.PlantVariety.GrowthDurationDays)
		fmt.Printf("Temperature Range: %.1f°C - %.1f°C\n", resp.PlantVariety.OptimalTempMin, resp.PlantVariety.OptimalTempMax)
		fmt.Printf("Water Requirement: %s\n", resp.PlantVariety.WaterRequirement)
		fmt.Printf("Light Requirement: %s\n", resp.PlantVariety.LightRequirement)
	}
}

func (c *CropServiceClient) TestListPlantVarieties() {
	fmt.Println("\n=== Test List Plant Varieties ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.ListPlantVarieties(ctx, &proto_plant_variety.ListPlantVarietiesRequest{
		Limit:         limit,
		Offset:        offset,
		SortBy:        "created_at",
		SortDirection: "desc",
	})
	if err != nil {
		fmt.Printf("Error calling ListPlantVarieties: %v\n", err)
		return
	}

	fmt.Printf("List Plant Varieties result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Category: %s, Season: %s\n",
			i+1, variety.Id, variety.Name, variety.Category, variety.GrowingSeason)
	}
}

func (c *CropServiceClient) TestSearchPlantVarieties() {
	fmt.Println("\n=== Test Search Plant Varieties ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter search name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.SearchPlantVarieties(ctx, &proto_plant_variety.SearchPlantVarietiesRequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		fmt.Printf("Error calling SearchPlantVarieties: %v\n", err)
		return
	}

	fmt.Printf("Search Plant Varieties result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Scientific Name: %s\n",
			i+1, variety.Id, variety.Name, variety.ScientificName)
	}
}

func (c *CropServiceClient) TestGetActivePlantVarieties() {
	fmt.Println("\n=== Test Get Active Plant Varieties ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetActivePlantVarieties(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Error calling GetActivePlantVarieties: %v\n", err)
		return
	}

	fmt.Printf("Get Active Plant Varieties result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Active Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Category: %s, Status: %s\n",
			i+1, variety.Id, variety.Name, variety.Category, variety.Status)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByCategory() {
	fmt.Println("\n=== Test Get Plant Varieties By Category ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter category: ")
	category, _ := reader.ReadString('\n')
	category = cleanInput(category)

	offset, limit := inputPaging(reader)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByCategory(ctx, &proto_plant_variety.GetPlantVarietiesByCategoryRequest{
		Category: category,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByCategory: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Category result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Category: %s\n",
			i+1, variety.Id, variety.Name, variety.Category)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesBySeason() {
	fmt.Println("\n=== Test Get Plant Varieties By Season ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter season: ")
	season, _ := reader.ReadString('\n')
	season = cleanInput(season)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesBySeason(ctx, &proto_plant_variety.GetPlantVarietiesBySeasonRequest{
		Season: season,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesBySeason: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Season result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Season: %s\n",
			i+1, variety.Id, variety.Name, variety.GrowingSeason)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByStatus() {
	fmt.Println("\n=== Test Get Plant Varieties By Status ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByStatus(ctx, &proto_plant_variety.GetPlantVarietiesByStatusRequest{
		Status: status,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByStatus: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Status result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Status: %s\n",
			i+1, variety.Id, variety.Name, variety.Status)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByTemperatureRange() {
	fmt.Println("\n=== Test Get Plant Varieties By Temperature Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter min temperature: ")
	minTempStr, _ := reader.ReadString('\n')
	minTempStr = cleanInput(minTempStr)
	minTemp := float64(20.0)
	if minTempStr != "" {
		if t, err := strconv.ParseFloat(minTempStr, 64); err == nil {
			minTemp = t
		}
	}

	fmt.Print("Enter max temperature: ")
	maxTempStr, _ := reader.ReadString('\n')
	maxTempStr = cleanInput(maxTempStr)
	maxTemp := float64(30.0)
	if maxTempStr != "" {
		if t, err := strconv.ParseFloat(maxTempStr, 64); err == nil {
			maxTemp = t
		}
	}

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByTemperatureRange(ctx, &proto_plant_variety.GetPlantVarietiesByTemperatureRangeRequest{
		MinTemp: minTemp,
		MaxTemp: maxTemp,
		Offset:  offset,
		Limit:   limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByTemperatureRange: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Temperature Range result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Temp Range: %.1f-%.1f°C\n",
			i+1, variety.Id, variety.Name, variety.OptimalTempMin, variety.OptimalTempMax)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByHumidityRange() {
	fmt.Println("\n=== Test Get Plant Varieties By Humidity Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter min humidity: ")
	minHumidityStr, _ := reader.ReadString('\n')
	minHumidityStr = cleanInput(minHumidityStr)
	minHumidity := float64(50.0)
	if minHumidityStr != "" {
		if h, err := strconv.ParseFloat(minHumidityStr, 64); err == nil {
			minHumidity = h
		}
	}

	fmt.Print("Enter max humidity: ")
	maxHumidityStr, _ := reader.ReadString('\n')
	maxHumidityStr = cleanInput(maxHumidityStr)
	maxHumidity := float64(80.0)
	if maxHumidityStr != "" {
		if h, err := strconv.ParseFloat(maxHumidityStr, 64); err == nil {
			maxHumidity = h
		}
	}

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByHumidityRange(ctx, &proto_plant_variety.GetPlantVarietiesByHumidityRangeRequest{
		MinHumidity: minHumidity,
		MaxHumidity: maxHumidity,
		Offset:      offset,
		Limit:       limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByHumidityRange: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Humidity Range result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Humidity Range: %.1f-%.1f%%\n",
			i+1, variety.Id, variety.Name, variety.OptimalHumidityMin, variety.OptimalHumidityMax)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByWaterRequirement() {
	fmt.Println("\n=== Test Get Plant Varieties By Water Requirement ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter water requirement: ")
	waterReq, _ := reader.ReadString('\n')
	waterReq = cleanInput(waterReq)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByWaterRequirement(ctx, &proto_plant_variety.GetPlantVarietiesByWaterRequirementRequest{
		WaterRequirement: waterReq,
		Offset:           offset,
		Limit:            limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByWaterRequirement: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Water Requirement result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Water Requirement: %s\n",
			i+1, variety.Id, variety.Name, variety.WaterRequirement)
	}
}

func (c *CropServiceClient) TestGetPlantVarietiesByLightRequirement() {
	fmt.Println("\n=== Test Get Plant Varieties By Light Requirement ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter light requirement: ")
	lightReq, _ := reader.ReadString('\n')
	lightReq = cleanInput(lightReq)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantVarietyClient.GetPlantVarietiesByLightRequirement(ctx, &proto_plant_variety.GetPlantVarietiesByLightRequirementRequest{
		LightRequirement: lightReq,
		Offset:           offset,
		Limit:            limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantVarietiesByLightRequirement: %v\n", err)
		return
	}

	fmt.Printf("Get Plant Varieties By Light Requirement result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Plant Varieties:\n")
	for i, variety := range resp.PlantVarieties {
		fmt.Printf("  [%d] ID: %s, Name: %s, Light Requirement: %s\n",
			i+1, variety.Id, variety.Name, variety.LightRequirement)
	}
}

// ================== Planting Cycle Service Tests ==================

func (c *CropServiceClient) TestCreatePlantingCycle() {
	fmt.Println("\n=== Test Create Planting Cycle ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter cycle name: ")
	cycleName, _ := reader.ReadString('\n')
	cycleName = cleanInput(cycleName)

	fmt.Print("Enter growing zone ID: ")
	growingZoneId, _ := reader.ReadString('\n')
	growingZoneId = cleanInput(growingZoneId)

	fmt.Print("Enter plant variety ID: ")
	plantVarietyId, _ := reader.ReadString('\n')
	plantVarietyId = cleanInput(plantVarietyId)

	fmt.Print("Enter seed date (YYYY-MM-DD): ")
	seedDateStr, _ := reader.ReadString('\n')
	seedDateStr = cleanInput(seedDateStr)
	var seedDate *timestamppb.Timestamp
	if seedDateStr != "" {
		if t, err := time.Parse("2006-01-02", seedDateStr); err == nil {
			seedDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter expected harvest date (YYYY-MM-DD): ")
	harvestDateStr, _ := reader.ReadString('\n')
	harvestDateStr = cleanInput(harvestDateStr)
	var harvestDate *timestamppb.Timestamp
	if harvestDateStr != "" {
		if t, err := time.Parse("2006-01-02", harvestDateStr); err == nil {
			harvestDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter plant quantity: ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = cleanInput(quantityStr)
	quantity := int32(100)
	if quantityStr != "" {
		if q, err := strconv.Atoi(quantityStr); err == nil {
			quantity = int32(q)
		}
	}

	fmt.Print("Enter seed batch: ")
	seedBatch, _ := reader.ReadString('\n')
	seedBatch = cleanInput(seedBatch)

	fmt.Print("Enter status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "planning"
	}

	fmt.Print("Enter notes: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.CreatePlantingCycle(ctx, &proto_planting_cycle.CreatePlantingCycleRequest{
		CycleName:           cycleName,
		GrowingZoneId:       growingZoneId,
		PlantVarietyId:      plantVarietyId,
		SeedDate:            seedDate,
		ExpectedHarvestDate: harvestDate,
		PlantQuantity:       quantity,
		SeedBatch:           seedBatch,
		Status:              status,
		Notes:               notes,
		CreatedBy:           createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreatePlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Create Planting Cycle result:\n")
	if resp.PlantingCycle != nil {
		fmt.Printf("ID: %s\n", resp.PlantingCycle.Id)
		fmt.Printf("Cycle Name: %s\n", resp.PlantingCycle.CycleName)
		fmt.Printf("Growing Zone ID: %s\n", resp.PlantingCycle.GrowingZoneId)
		fmt.Printf("Plant Variety ID: %s\n", resp.PlantingCycle.PlantVarietyId)
		fmt.Printf("Plant Quantity: %d\n", resp.PlantingCycle.PlantQuantity)
		fmt.Printf("Status: %s\n", resp.PlantingCycle.Status)
	}
}

func (c *CropServiceClient) TestGetPlantingCycle() {
	fmt.Println("\n=== Test Get Planting Cycle ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCycle(ctx, &proto_planting_cycle.GetPlantingCycleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycle result:\n")
	if resp.PlantingCycle != nil {
		fmt.Printf("ID: %s\n", resp.PlantingCycle.Id)
		fmt.Printf("Cycle Name: %s\n", resp.PlantingCycle.CycleName)
		fmt.Printf("Growing Zone ID: %s\n", resp.PlantingCycle.GrowingZoneId)
		fmt.Printf("Plant Variety ID: %s\n", resp.PlantingCycle.PlantVarietyId)
		fmt.Printf("Plant Quantity: %d\n", resp.PlantingCycle.PlantQuantity)
		fmt.Printf("Status: %s\n", resp.PlantingCycle.Status)
		if resp.PlantingCycle.SeedDate != nil {
			fmt.Printf("Seed Date: %s\n", resp.PlantingCycle.SeedDate.AsTime().Format("2006-01-02"))
		}
		if resp.PlantingCycle.ExpectedHarvestDate != nil {
			fmt.Printf("Expected Harvest Date: %s\n", resp.PlantingCycle.ExpectedHarvestDate.AsTime().Format("2006-01-02"))
		}
	}
}

func (c *CropServiceClient) TestListPlantingCycles() {
	fmt.Println("\n=== Test List Planting Cycles ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.ListPlantingCycles(ctx, &proto_planting_cycle.ListPlantingCyclesRequest{
		Offset:        offset,
		Limit:         limit,
		SortBy:        "created_at",
		SortDirection: "desc",
	})

	if err != nil {
		fmt.Printf("Error calling ListPlantingCycles: %v\n", err)
		return
	}

	fmt.Printf("List Planting Cycles result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Status: %s, Plant Quantity: %d\n",
			i+1, cycle.Id, cycle.CycleName, cycle.Status, cycle.PlantQuantity)
	}
}

func (c *CropServiceClient) TestGetActivePlantingCycles() {
	fmt.Println("\n=== Test Get Active Planting Cycles ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetActivePlantingCycles(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Error calling GetActivePlantingCycles: %v\n", err)
		return
	}

	fmt.Printf("Get Active Planting Cycles result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Active Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Status: %s\n",
			i+1, cycle.Id, cycle.CycleName, cycle.Status)
	}
}

func (c *CropServiceClient) TestGetUpcomingHarvests() {
	fmt.Println("\n=== Test Get Upcoming Harvests ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter days ahead (default 7): ")
	daysStr, _ := reader.ReadString('\n')
	daysStr = cleanInput(daysStr)
	days := int32(7)
	if daysStr != "" {
		if d, err := strconv.Atoi(daysStr); err == nil {
			days = int32(d)
		}
	}

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetUpcomingHarvests(ctx, &proto_planting_cycle.GetUpcomingHarvestsRequest{
		Days:   days,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetUpcomingHarvests: %v\n", err)
		return
	}

	fmt.Printf("Get Upcoming Harvests result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Upcoming Harvests:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Expected Harvest: %s\n",
			i+1, cycle.Id, cycle.CycleName,
			cycle.ExpectedHarvestDate.AsTime().Format("2006-01-02"))
	}
}

func (c *CropServiceClient) TestGetOverdueHarvests() {
	fmt.Println("\n=== Test Get Overdue Harvests ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetOverdueHarvests(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Error calling GetOverdueHarvests: %v\n", err)
		return
	}

	fmt.Printf("Get Overdue Harvests result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Overdue Harvests:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Expected Harvest: %s\n",
			i+1, cycle.Id, cycle.CycleName,
			cycle.ExpectedHarvestDate.AsTime().Format("2006-01-02"))
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesByZone() {
	fmt.Println("\n=== Test Get Planting Cycles By Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter growing zone ID: ")
	zoneId, _ := reader.ReadString('\n')
	zoneId = cleanInput(zoneId)

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesByZone(ctx, &proto_planting_cycle.GetPlantingCyclesByZoneRequest{
		GrowingZoneId: zoneId,
		Limit:         limit,
		Offset:        offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesByZone: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Zone result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Zone ID: %s\n",
			i+1, cycle.Id, cycle.CycleName, cycle.GrowingZoneId)
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesByVariety() {
	fmt.Println("\n=== Test Get Planting Cycles By Variety ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plant variety ID: ")
	varietyId, _ := reader.ReadString('\n')
	varietyId = cleanInput(varietyId)

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesByVariety(ctx, &proto_planting_cycle.GetPlantingCyclesByVarietyRequest{
		PlantVarietyId: varietyId,
		Limit:          limit,
		Offset:         offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesByVariety: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Variety result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Variety ID: %s\n",
			i+1, cycle.Id, cycle.CycleName, cycle.PlantVarietyId)
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesByStatus() {
	fmt.Println("\n=== Test Get Planting Cycles By Status ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesByStatus(ctx, &proto_planting_cycle.GetPlantingCyclesByStatusRequest{
		Status: status,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesByStatus: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Status result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Status: %s\n",
			i+1, cycle.Id, cycle.CycleName, cycle.Status)
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesByDateRange() {
	fmt.Println("\n=== Test Get Planting Cycles By Date Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start date (YYYY-MM-DD): ")
	startDateStr, _ := reader.ReadString('\n')
	startDateStr = cleanInput(startDateStr)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using 30 days ago.\n")
		startDate = time.Now().AddDate(0, 0, -30)
	}

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	endDateStr, _ := reader.ReadString('\n')
	endDateStr = cleanInput(endDateStr)
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		endDate = time.Now()
	}

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesByDateRange(ctx, &proto_planting_cycle.GetPlantingCyclesByDateRangeRequest{
		StartDate: timestamppb.New(startDate),
		EndDate:   timestamppb.New(endDate),
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesByDateRange: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Date Range result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Seed Date: %s\n",
			i+1, cycle.Id, cycle.CycleName,
			cycle.SeedDate.AsTime().Format("2006-01-02"))
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesByHarvestDateRange() {
	fmt.Println("\n=== Test Get Planting Cycles By Harvest Date Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start harvest date (YYYY-MM-DD): ")
	startDateStr, _ := reader.ReadString('\n')
	startDateStr = cleanInput(startDateStr)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using 30 days ago.\n")
		startDate = time.Now().AddDate(0, 0, -30)
	}

	fmt.Print("Enter end harvest date (YYYY-MM-DD): ")
	endDateStr, _ := reader.ReadString('\n')
	endDateStr = cleanInput(endDateStr)
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		endDate = time.Now()
	}

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesByHarvestDateRange(ctx, &proto_planting_cycle.GetPlantingCyclesByHarvestDateRangeRequest{
		StartDate: timestamppb.New(startDate),
		EndDate:   timestamppb.New(endDate),
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesByHarvestDateRange: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Harvest Date Range result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Expected Harvest: %s\n",
			i+1, cycle.Id, cycle.CycleName,
			cycle.ExpectedHarvestDate.AsTime().Format("2006-01-02"))
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesBySeedDateRange() {
	fmt.Println("\n=== Test Get Planting Cycles By Seed Date Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start seed date (YYYY-MM-DD): ")
	startDateStr, _ := reader.ReadString('\n')
	startDateStr = cleanInput(startDateStr)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using 30 days ago.\n")
		startDate = time.Now().AddDate(0, 0, -30)
	}

	fmt.Print("Enter end seed date (YYYY-MM-DD): ")
	endDateStr, _ := reader.ReadString('\n')
	endDateStr = cleanInput(endDateStr)
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		endDate = time.Now()
	}

	limit, offset := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesBySeedDateRange(ctx, &proto_planting_cycle.GetPlantingCyclesBySeedDateRangeRequest{
		StartDate: timestamppb.New(startDate),
		EndDate:   timestamppb.New(endDate),
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesBySeedDateRange: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles By Seed Date Range result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCycles {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Seed Date: %s\n",
			i+1, cycle.Id, cycle.CycleName,
			cycle.SeedDate.AsTime().Format("2006-01-02"))
	}
}

func (c *CropServiceClient) TestGetPlantingCycleWithDetails() {
	fmt.Println("\n=== Test Get Planting Cycle With Details ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCycleWithDetails(ctx, &proto_planting_cycle.GetPlantingCycleWithDetailsRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCycleWithDetails: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycle With Details result:\n")
	if resp.PlantingCycleWithDetails != nil {
		fmt.Printf("ID: %s\n", resp.PlantingCycleWithDetails.PlantingCycle.Id)
		fmt.Printf("Cycle Name: %s\n", resp.PlantingCycleWithDetails.PlantingCycle.CycleName)
		fmt.Printf("Growing Zone ID: %s\n", resp.PlantingCycleWithDetails.PlantingCycle.GrowingZoneId)
		fmt.Printf("Plant Variety ID: %s\n", resp.PlantingCycleWithDetails.PlantingCycle.PlantVarietyId)
		fmt.Printf("Status: %s\n", resp.PlantingCycleWithDetails.PlantingCycle.Status)
		if resp.PlantingCycleWithDetails.PlantVariety != nil {
			fmt.Printf("Plant Variety: %s (%s)\n",
				resp.PlantingCycleWithDetails.PlantVariety.Name,
				resp.PlantingCycleWithDetails.PlantVariety.ScientificName)
		}
	}
}

func (c *CropServiceClient) TestGetPlantingCyclesWithDetails() {
	fmt.Println("\n=== Test Get Planting Cycles With Details ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.GetPlantingCyclesWithDetails(ctx, &proto_planting_cycle.GetPlantingCyclesWithDetailsRequest{
		Limit:  limit,
		Offset: (page - 1) * limit,
	})
	if err != nil {
		fmt.Printf("Error calling GetPlantingCyclesWithDetails: %v\n", err)
		return
	}

	fmt.Printf("Get Planting Cycles With Details result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Planting Cycles:\n")
	for i, cycle := range resp.PlantingCyclesWithDetails {
		fmt.Printf("  [%d] ID: %s, Cycle Name: %s, Status: %s\n",
			i+1, cycle.PlantingCycle.Id, cycle.PlantingCycle.CycleName, cycle.PlantingCycle.Status)
		if cycle.PlantVariety != nil {
			fmt.Printf("      Plant Variety: %s\n", cycle.PlantVariety.Name)
		}
	}
}

func (c *CropServiceClient) TestUpdatePlantingCycleStatus() {
	fmt.Println("\n=== Test Update Planting Cycle Status ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.UpdatePlantingCycleStatus(ctx, &proto_planting_cycle.UpdatePlantingCycleStatusRequest{
		Id:     id,
		Status: status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdatePlantingCycleStatus: %v\n", err)
		return
	}

	fmt.Printf("Update Planting Cycle Status result:\n")
	if resp.PlantingCycle != nil {
		fmt.Printf("ID: %s\n", resp.PlantingCycle.Id)
		fmt.Printf("Cycle Name: %s\n", resp.PlantingCycle.CycleName)
		fmt.Printf("New Status: %s\n", resp.PlantingCycle.Status)
	}
}

func (c *CropServiceClient) TestUpdatePlantingCycleHarvestDate() {
	fmt.Println("\n=== Test Update Planting Cycle Harvest Date ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter harvest date (YYYY-MM-DD): ")
	harvestDateStr, _ := reader.ReadString('\n')
	harvestDateStr = cleanInput(harvestDateStr)
	harvestDate, err := time.Parse("2006-01-02", harvestDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		harvestDate = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.plantingCycleClient.UpdatePlantingCycleHarvestDate(ctx, &proto_planting_cycle.UpdatePlantingCycleHarvestDateRequest{
		Id:          id,
		HarvestDate: timestamppb.New(harvestDate),
	})
	if err != nil {
		fmt.Printf("Error calling UpdatePlantingCycleHarvestDate: %v\n", err)
		return
	}

	fmt.Printf("Update Planting Cycle Harvest Date result:\n")
	if resp.PlantingCycle != nil {
		fmt.Printf("ID: %s\n", resp.PlantingCycle.Id)
		fmt.Printf("Cycle Name: %s\n", resp.PlantingCycle.CycleName)
		if resp.PlantingCycle.ActualHarvestDate != nil {
			fmt.Printf("Actual Harvest Date: %s\n",
				resp.PlantingCycle.ActualHarvestDate.AsTime().Format("2006-01-02"))
		}
	}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Crop Service Test Client ===")
	fmt.Println("1. Plant Variety Service")
	fmt.Println("2. Planting Cycle Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printPlantVarietyMenu() {
	fmt.Println("\n=== Plant Variety Service ===")
	fmt.Println("1. Create Plant Variety")
	fmt.Println("2. Get Plant Variety")
	fmt.Println("3. List Plant Varieties")
	fmt.Println("4. Search Plant Varieties")
	fmt.Println("5. Get Active Plant Varieties")
	fmt.Println("6. Get By Category")
	fmt.Println("7. Get By Season")
	fmt.Println("8. Get By Status")
	fmt.Println("9. Get By Temperature Range")
	fmt.Println("10. Get By Humidity Range")
	fmt.Println("11. Get By Water Requirement")
	fmt.Println("12. Get By Light Requirement")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printPlantingCycleMenu() {
	fmt.Println("\n=== Planting Cycle Service ===")
	fmt.Println("1. Create Planting Cycle")
	fmt.Println("2. Get Planting Cycle")
	fmt.Println("3. List Planting Cycles")
	fmt.Println("4. Get Active Planting Cycles")
	fmt.Println("5. Get Upcoming Harvests")
	fmt.Println("6. Get Overdue Harvests")
	fmt.Println("7. Get By Zone")
	fmt.Println("8. Get By Variety")
	fmt.Println("9. Get By Status")
	fmt.Println("10. Get By Date Range")
	fmt.Println("11. Get By Harvest Date Range")
	fmt.Println("12. Get By Seed Date Range")
	fmt.Println("13. Get Cycle With Details")
	fmt.Println("14. Get Cycles With Details")
	fmt.Println("15. Update Status")
	fmt.Println("16. Update Harvest Date")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewCropServiceClient(address)
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
			// Plant Variety Service
			for {
				printPlantVarietyMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreatePlantVariety()
				case "2":
					client.TestGetPlantVariety()
				case "3":
					client.TestListPlantVarieties()
				case "4":
					client.TestSearchPlantVarieties()
				case "5":
					client.TestGetActivePlantVarieties()
				case "6":
					client.TestGetPlantVarietiesByCategory()
				case "7":
					client.TestGetPlantVarietiesBySeason()
				case "8":
					client.TestGetPlantVarietiesByStatus()
				case "9":
					client.TestGetPlantVarietiesByTemperatureRange()
				case "10":
					client.TestGetPlantVarietiesByHumidityRange()
				case "11":
					client.TestGetPlantVarietiesByWaterRequirement()
				case "12":
					client.TestGetPlantVarietiesByLightRequirement()
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
			// Planting Cycle Service
			for {
				printPlantingCycleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreatePlantingCycle()
				case "2":
					client.TestGetPlantingCycle()
				case "3":
					client.TestListPlantingCycles()
				case "4":
					client.TestGetActivePlantingCycles()
				case "5":
					client.TestGetUpcomingHarvests()
				case "6":
					client.TestGetOverdueHarvests()
				case "7":
					client.TestGetPlantingCyclesByZone()
				case "8":
					client.TestGetPlantingCyclesByVariety()
				case "9":
					client.TestGetPlantingCyclesByStatus()
				case "10":
					client.TestGetPlantingCyclesByDateRange()
				case "11":
					client.TestGetPlantingCyclesByHarvestDateRange()
				case "12":
					client.TestGetPlantingCyclesBySeedDateRange()
				case "13":
					client.TestGetPlantingCycleWithDetails()
				case "14":
					client.TestGetPlantingCyclesWithDetails()
				case "15":
					client.TestUpdatePlantingCycleStatus()
				case "16":
					client.TestUpdatePlantingCycleHarvestDate()
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
