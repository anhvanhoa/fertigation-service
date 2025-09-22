package main

import (
	"context"
	"fertigation-Service/bootstrap"
	"fertigation-Service/infrastructure/grpc_service"
	fertilizer_schedule_service "fertigation-Service/infrastructure/grpc_service/fertilizer_schedule"
	fertilizer_type_service "fertigation-Service/infrastructure/grpc_service/fertilizer_type"
	irrigation_log_service "fertigation-Service/infrastructure/grpc_service/irrigation_log"
	irrigation_schedule_service "fertigation-Service/infrastructure/grpc_service/irrigation_schedule"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	fertilizerScheduleService := fertilizer_schedule_service.NewFertilizerScheduleService(app.Repos.FertilizerScheduleRepository)
	fertilizerTypeService := fertilizer_type_service.NewFertilizerTypeService(app.Repos.FertilizerTypeRepository)
	irrigationScheduleService := irrigation_schedule_service.NewIrrigationScheduleService(app.Repos.IrrigationScheduleRepository)
	irrigationLogService := irrigation_log_service.NewIrrigationLogService(app.Repos.IrrigationLogRepository)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		fertilizerScheduleService,
		fertilizerTypeService,
		irrigationScheduleService,
		irrigationLogService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
