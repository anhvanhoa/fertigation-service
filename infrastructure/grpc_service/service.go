package grpc_service

import (
	"fertigation-Service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	proto_fertilizer_type "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	proto_irrigation_log "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	proto_irrigation_schedule "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	fertilizer_schedule proto_fertilizer_schedule.FertilizerScheduleServiceServer,
	fertilizer_type proto_fertilizer_type.FertilizerTypeServiceServer,
	irrigation_schedule proto_irrigation_schedule.IrrigationScheduleServiceServer,
	irrigation_log proto_irrigation_log.IrrigationLogServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_fertilizer_schedule.RegisterFertilizerScheduleServiceServer(server, fertilizer_schedule)
			proto_fertilizer_type.RegisterFertilizerTypeServiceServer(server, fertilizer_type)
			proto_irrigation_schedule.RegisterIrrigationScheduleServiceServer(server, irrigation_schedule)
			proto_irrigation_log.RegisterIrrigationLogServiceServer(server, irrigation_log)
		},
	)
}
