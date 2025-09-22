syntax = "proto3";

package fertilizer_schedule.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1;proto_fertilizer_schedule";

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

// Fertilizer Schedule Service
service FertilizerScheduleService {
  // Basic CRUD operations
  rpc CreateFertilizerSchedule(CreateFertilizerScheduleRequest)
      returns (FertilizerScheduleResponse);
  rpc GetFertilizerSchedule(GetFertilizerScheduleRequest)
      returns (FertilizerScheduleResponse);
  rpc UpdateFertilizerSchedule(UpdateFertilizerScheduleRequest)
      returns (FertilizerScheduleResponse);
  rpc DeleteFertilizerSchedule(DeleteFertilizerScheduleRequest)
      returns (DeleteFertilizerScheduleResponse);
  rpc ListFertilizerSchedules(FilterFertilizerSchedulesRequest)
      returns (ListFertilizerSchedulesResponse);

  // Special operations
  rpc GetSchedulesByPlantingCycle(GetSchedulesByPlantingCycleRequest)
      returns (ListFertilizerSchedulesResponse);
  rpc GetCompletedSchedules(Pagination)
      returns (ListFertilizerSchedulesResponse);
  rpc GetPendingSchedules(Pagination)
      returns (ListFertilizerSchedulesResponse);
  rpc GetUpcomingSchedules(GetUpcomingSchedulesRequest)
      returns (ListFertilizerSchedulesResponse);
}

message Pagination {
  int32 page = 1 [ (buf.validate.field).int32 = {gte : 1} ];
  int32 page_size = 2 [ (buf.validate.field).int32 = {gte : 1, lte : 100} ];
  string sort_by = 3 [ (buf.validate.field).string = {max_len : 50} ];
  string sort_order = 4 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
  string search = 5 [ (buf.validate.field).string = {max_len : 255} ];
}

// Request/Response messages
message CreateFertilizerScheduleRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.uuid = true ];
  string fertilizer_type_id = 2 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp application_date = 3;
  double dosage = 4 [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  string unit = 5 [ (buf.validate.field).string = {max_len : 20} ];
  string application_method = 6 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  string growth_stage = 7 [
    (buf.validate.field).string.in = "seedling",
    (buf.validate.field).string.in = "vegetative",
    (buf.validate.field).string.in = "flowering",
    (buf.validate.field).string.in = "fruiting",
    (buf.validate.field).string.in = "pre_harvest"
  ];
  string weather_conditions = 8
      [ (buf.validate.field).string = {max_len : 200} ];
  string soil_conditions = 9 [ (buf.validate.field).string = {max_len : 200} ];
  bool is_completed = 10;
  google.protobuf.Timestamp completed_date = 11;
  double actual_dosage = 12
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  int32 effectiveness_rating = 13
      [ (buf.validate.field).int32 = {gte : 1, lte : 5} ];
  string notes = 14 [ (buf.validate.field).string = {max_len : 1000} ];
  string created_by = 15 [ (buf.validate.field).string.uuid = true ];
}

message GetFertilizerScheduleRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message UpdateFertilizerScheduleRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
  string planting_cycle_id = 2 [ (buf.validate.field).string.uuid = true ];
  string fertilizer_type_id = 3 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp application_date = 4;
  double dosage = 5 [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  string unit = 6 [ (buf.validate.field).string = {max_len : 20} ];
  string application_method = 7 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  string growth_stage = 8 [
    (buf.validate.field).string.in = "seedling",
    (buf.validate.field).string.in = "vegetative",
    (buf.validate.field).string.in = "flowering",
    (buf.validate.field).string.in = "fruiting",
    (buf.validate.field).string.in = "pre_harvest"
  ];
  string weather_conditions = 9
      [ (buf.validate.field).string = {max_len : 200} ];
  string soil_conditions = 10 [ (buf.validate.field).string = {max_len : 200} ];
  bool is_completed = 11;
  google.protobuf.Timestamp completed_date = 12;
  double actual_dosage = 13
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  int32 effectiveness_rating = 14
      [ (buf.validate.field).int32 = {gte : 1, lte : 5} ];
  string notes = 15 [ (buf.validate.field).string = {max_len : 1000} ];
}

message DeleteFertilizerScheduleRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message DeleteFertilizerScheduleResponse {
  string message = 1;
}

message FilterFertilizerSchedulesRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.uuid = true ];
  string fertilizer_type_id = 2 [ (buf.validate.field).string.uuid = true ];
  string application_method = 3 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  string growth_stage = 4 [
    (buf.validate.field).string.in = "seedling",
    (buf.validate.field).string.in = "vegetative",
    (buf.validate.field).string.in = "flowering",
    (buf.validate.field).string.in = "fruiting",
    (buf.validate.field).string.in = "pre_harvest"
  ];
  bool is_completed = 5;
  string created_by = 6 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp application_date_from = 7;
  google.protobuf.Timestamp application_date_to = 8;
  google.protobuf.Timestamp completed_date_from = 9;
  google.protobuf.Timestamp completed_date_to = 10;
  google.protobuf.Timestamp created_at_from = 11;
  google.protobuf.Timestamp created_at_to = 12;
  int32 page = 13 [ (buf.validate.field).int32 = {gte : 1} ];
  int32 limit = 14 [ (buf.validate.field).int32 = {gte : 1, lte : 100} ];
  string sort_by = 15 [ (buf.validate.field).string = {max_len : 50} ];
  string sort_order = 16 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
}

message GetSchedulesByPlantingCycleRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.uuid = true ];
  Pagination pagination = 2;
}

message GetUpcomingSchedulesRequest {
  int32 days = 1 [ (buf.validate.field).int32 = {gte : 1, lte : 365} ];
  Pagination pagination = 2;
}

message FertilizerScheduleResponse {
  string id = 1;
  string planting_cycle_id = 2;
  string fertilizer_type_id = 3;
  google.protobuf.Timestamp application_date = 4;
  double dosage = 5;
  string unit = 6;
  string application_method = 7;
  string growth_stage = 8;
  string weather_conditions = 9;
  string soil_conditions = 10;
  bool is_completed = 11;
  google.protobuf.Timestamp completed_date = 12;
  double actual_dosage = 13;
  int32 effectiveness_rating = 14;
  string notes = 15;
  string created_by = 16;
  google.protobuf.Timestamp created_at = 17;
  google.protobuf.Timestamp updated_at = 18;
}

message ListFertilizerSchedulesResponse {
  repeated FertilizerScheduleResponse fertilizer_schedules = 1;
  int64 total = 2;
  int64 page = 3;
  int64 page_size = 4;
  int64 total_pages = 5;
}


syntax = "proto3";

package fertilizer_type.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1;proto_fertilizer_type";

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

// Fertilizer Type Service
service FertilizerTypeService {
  // Basic CRUD operations
  rpc CreateFertilizerType(CreateFertilizerTypeRequest) returns (FertilizerTypeResponse);
  rpc GetFertilizerType(GetFertilizerTypeRequest) returns (FertilizerTypeResponse);
  rpc UpdateFertilizerType(UpdateFertilizerTypeRequest) returns (FertilizerTypeResponse);
  rpc DeleteFertilizerType(DeleteFertilizerTypeRequest) returns (DeleteFertilizerTypeResponse);
  rpc ListFertilizerTypes(ListFertilizerTypesRequest) returns (ListFertilizerTypesResponse);

  // Special operations
  rpc GetFertilizerTypesByType(GetFertilizerTypesByTypeRequest) returns (ListFertilizerTypesResponse);
  rpc GetExpiredFertilizers(Pagination) returns (ListFertilizerTypesResponse);
  rpc GetExpiringSoon(GetExpiringSoonRequest) returns (ListFertilizerTypesResponse);
}

message Pagination {
  int32 page = 1 [ (buf.validate.field).int32 = {gte : 1} ];
  int32 page_size = 2 [ (buf.validate.field).int32 = {gte : 1, lte : 100} ];
  string sort_by = 3 [ (buf.validate.field).string = {max_len : 50} ];
  string sort_order = 4 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
  string search = 5 [ (buf.validate.field).string = {max_len : 255} ];
}

// Request/Response messages
message CreateFertilizerTypeRequest {
  string name = 1
      [ (buf.validate.field).string = {min_len : 1, max_len : 255} ];
  string type = 2 [
    (buf.validate.field).string.in = "organic",
    (buf.validate.field).string.in = "chemical",
    (buf.validate.field).string.in = "liquid",
    (buf.validate.field).string.in = "granular",
    (buf.validate.field).string.in = "powder"
  ];
  string npk_ratio = 3 [ (buf.validate.field).string = {max_len : 20} ];
  double nitrogen_percentage = 4
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  double phosphorus_percentage = 5
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  double potassium_percentage = 6
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  string trace_elements = 7;
  string application_method = 8 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  double dosage_per_plant = 9
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  double dosage_per_m2 = 10
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  string unit = 11 [ (buf.validate.field).string = {max_len : 20} ];
  string manufacturer = 12 [ (buf.validate.field).string = {max_len : 255} ];
  string batch_number = 13 [ (buf.validate.field).string = {max_len : 100} ];
  google.protobuf.Timestamp expiry_date = 14;
  double cost_per_unit = 15 [ (buf.validate.field).double = {gte : 0} ];
  string description = 16 [ (buf.validate.field).string = {max_len : 1000} ];
  string safety_notes = 17 [ (buf.validate.field).string = {max_len : 1000} ];
  string status = 18 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "expired"
  ];
  string created_by = 19 [ (buf.validate.field).string.uuid = true ];
}

message GetFertilizerTypeRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message UpdateFertilizerTypeRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
  string name = 2
      [ (buf.validate.field).string = {min_len : 1, max_len : 255} ];
  string type = 3 [
    (buf.validate.field).string.in = "organic",
    (buf.validate.field).string.in = "chemical",
    (buf.validate.field).string.in = "liquid",
    (buf.validate.field).string.in = "granular",
    (buf.validate.field).string.in = "powder"
  ];
  string npk_ratio = 4 [ (buf.validate.field).string = {max_len : 20} ];
  double nitrogen_percentage = 5
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  double phosphorus_percentage = 6
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  double potassium_percentage = 7
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  string trace_elements = 8;
  string application_method = 9 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  double dosage_per_plant = 10
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  double dosage_per_m2 = 11
      [ (buf.validate.field).double = {gte : 0, lte : 1000} ];
  string unit = 12 [ (buf.validate.field).string = {max_len : 20} ];
  string manufacturer = 13 [ (buf.validate.field).string = {max_len : 255} ];
  string batch_number = 14 [ (buf.validate.field).string = {max_len : 100} ];
  google.protobuf.Timestamp expiry_date = 15;
  double cost_per_unit = 16 [ (buf.validate.field).double = {gte : 0} ];
  string description = 17 [ (buf.validate.field).string = {max_len : 1000} ];
  string safety_notes = 18 [ (buf.validate.field).string = {max_len : 1000} ];
  string status = 19 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "expired"
  ];
}

message DeleteFertilizerTypeRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}
message DeleteFertilizerTypeResponse {
  string message = 1;
}

message ListFertilizerTypesRequest {
  string name = 1 [ (buf.validate.field).string = {max_len : 255} ];
  string type = 2 [
    (buf.validate.field).string.in = "organic",
    (buf.validate.field).string.in = "chemical",
    (buf.validate.field).string.in = "liquid",
    (buf.validate.field).string.in = "granular",
    (buf.validate.field).string.in = "powder"
  ];
  string application_method = 3 [
    (buf.validate.field).string.in = "foliar",
    (buf.validate.field).string.in = "soil",
    (buf.validate.field).string.in = "hydroponic",
    (buf.validate.field).string.in = "fertigation"
  ];
  string status = 4 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "inactive",
    (buf.validate.field).string.in = "expired"
  ];
  string manufacturer = 5 [ (buf.validate.field).string = {max_len : 255} ];
  string created_by = 6 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp expiry_date_from = 7;
  google.protobuf.Timestamp expiry_date_to = 8;
  google.protobuf.Timestamp created_at_from = 9;
  google.protobuf.Timestamp created_at_to = 10;
  Pagination pagination = 11;
}

message GetFertilizerTypesByTypeRequest {
  string type = 1 [
    (buf.validate.field).string.in = "organic",
    (buf.validate.field).string.in = "chemical",
    (buf.validate.field).string.in = "liquid",
    (buf.validate.field).string.in = "granular",
    (buf.validate.field).string.in = "powder"
  ];
  Pagination pagination = 2;
}

message GetExpiringSoonRequest {
  int32 days = 1 [ (buf.validate.field).int32 = {gte : 1, lte : 365} ];
  Pagination pagination = 2;
}

message FertilizerTypeResponse {
  string id = 1;
  string name = 2;
  string type = 3;
  string npk_ratio = 4;
  double nitrogen_percentage = 5;
  double phosphorus_percentage = 6;
  double potassium_percentage = 7;
  string trace_elements = 8;
  string application_method = 9;
  double dosage_per_plant = 10;
  double dosage_per_m2 = 11;
  string unit = 12;
  string manufacturer = 13;
  string batch_number = 14;
  google.protobuf.Timestamp expiry_date = 15;
  double cost_per_unit = 16;
  string description = 17;
  string safety_notes = 18;
  string status = 19;
  string created_by = 20;
  google.protobuf.Timestamp created_at = 21;
  google.protobuf.Timestamp updated_at = 22;
}

message ListFertilizerTypesResponse {
  repeated FertilizerTypeResponse fertilizer_types = 1;
  int64 total = 2;
  int64 page = 3;
  int64 page_size = 4;
  int64 total_pages = 5;
}


syntax = "proto3";

package irrigation_log;

option go_package = "fertigation-Service/proto/irrigation_log";

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

// Irrigation Log Service
service IrrigationLogService {
  // Basic CRUD operations
  rpc CreateIrrigationLog(CreateIrrigationLogRequest)
      returns (IrrigationLogResponse);
  rpc GetIrrigationLog(GetIrrigationLogRequest) returns (IrrigationLogResponse);
  rpc UpdateIrrigationLog(UpdateIrrigationLogRequest)
      returns (IrrigationLogResponse);
  rpc DeleteIrrigationLog(DeleteIrrigationLogRequest)
      returns (DeleteIrrigationLogResponse);
  rpc ListIrrigationLogs(ListIrrigationLogsRequest)
      returns (ListIrrigationLogsResponse);
}

// Request/Response messages
message CreateIrrigationLogRequest {
  string irrigation_schedule_id = 1 [ (buf.validate.field).string.uuid = true ];
  string device_id = 2 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp started_at = 3;
  google.protobuf.Timestamp ended_at = 4;
  int32 planned_duration_minutes = 5
      [ (buf.validate.field).int32 = {gte : 0, lte : 1440} ];
  int32 actual_duration_minutes = 6
      [ (buf.validate.field).int32 = {gte : 0, lte : 1440} ];
  double water_used_liters = 7
      [ (buf.validate.field).double = {gte : 0, lte : 10000} ];
  double water_pressure = 8
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  string status = 9 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "running",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "failed",
    (buf.validate.field).string.in = "cancelled"
  ];
  string failure_reason = 10 [ (buf.validate.field).string = {max_len : 500} ];
  string notes = 11 [ (buf.validate.field).string = {max_len : 1000} ];
  string created_by = 12 [ (buf.validate.field).string.uuid = true ];
}

message GetIrrigationLogRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message UpdateIrrigationLogRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
  string irrigation_schedule_id = 2 [ (buf.validate.field).string.uuid = true ];
  string device_id = 3 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp started_at = 4;
  google.protobuf.Timestamp ended_at = 5;
  int32 planned_duration_minutes = 6
      [ (buf.validate.field).int32 = {gte : 0, lte : 1440} ];
  int32 actual_duration_minutes = 7
      [ (buf.validate.field).int32 = {gte : 0, lte : 1440} ];
  double water_used_liters = 8
      [ (buf.validate.field).double = {gte : 0, lte : 10000} ];
  double water_pressure = 9
      [ (buf.validate.field).double = {gte : 0, lte : 100} ];
  string status = 10 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "running",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "failed",
    (buf.validate.field).string.in = "cancelled"
  ];
  string failure_reason = 11 [ (buf.validate.field).string = {max_len : 500} ];
  string notes = 12 [ (buf.validate.field).string = {max_len : 1000} ];
}

message DeleteIrrigationLogRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message DeleteIrrigationLogResponse {
  string message = 1;
}

message ListIrrigationLogsRequest {
  string irrigation_schedule_id = 1 [ (buf.validate.field).string.uuid = true ];
  string device_id = 2 [ (buf.validate.field).string.uuid = true ];
  string status = 3 [
    (buf.validate.field).string.in = "pending",
    (buf.validate.field).string.in = "running",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "failed",
    (buf.validate.field).string.in = "cancelled"
  ];
  string created_by = 4 [ (buf.validate.field).string.uuid = true ];
  google.protobuf.Timestamp started_at_from = 5;
  google.protobuf.Timestamp started_at_to = 6;
  google.protobuf.Timestamp ended_at_from = 7;
  google.protobuf.Timestamp ended_at_to = 8;
  google.protobuf.Timestamp created_at_from = 9;
  google.protobuf.Timestamp created_at_to = 10;
  int32 page = 11 [ (buf.validate.field).int32 = {gte : 1} ];
  int32 limit = 12 [ (buf.validate.field).int32 = {gte : 1, lte : 100} ];
  string sort_by = 13 [ (buf.validate.field).string = {max_len : 50} ];
  string sort_order = 14 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
}

message IrrigationLogResponse {
  string id = 1;
  string irrigation_schedule_id = 2;
  string device_id = 3;
  google.protobuf.Timestamp started_at = 4;
  google.protobuf.Timestamp ended_at = 5;
  int32 planned_duration_minutes = 6;
  int32 actual_duration_minutes = 7;
  double water_used_liters = 8;
  double water_pressure = 9;
  string status = 10;
  string failure_reason = 11;
  string notes = 12;
  string created_by = 13;
  google.protobuf.Timestamp created_at = 14;
}

message ListIrrigationLogsResponse {
  repeated IrrigationLogResponse irrigation_logs = 1;
  int32 total = 2;
  int32 page = 3;
  int32 limit = 4;
  int32 total_pages = 5;
}


syntax = "proto3";

package irrigation_schedule.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1;proto_irrigation_schedule";

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

// Irrigation Schedule Service
service IrrigationScheduleService {
  // Basic CRUD operations
  rpc CreateIrrigationSchedule(CreateIrrigationScheduleRequest) returns (IrrigationScheduleResponse);
  rpc GetIrrigationSchedule(GetIrrigationScheduleRequest) returns (IrrigationScheduleResponse);
  rpc UpdateIrrigationSchedule(UpdateIrrigationScheduleRequest) returns (IrrigationScheduleResponse);
  rpc DeleteIrrigationSchedule(DeleteIrrigationScheduleRequest) returns (common.CommonResponse);
  rpc ListIrrigationSchedules(ListIrrigationSchedulesRequest) returns (ListIrrigationSchedulesResponse);
  
  // Special operations
  rpc GetSchedulesByGrowingZone(GetSchedulesByGrowingZoneRequest) returns (ListIrrigationSchedulesResponse);
  rpc GetActiveSchedules(common.PaginationRequest) returns (ListIrrigationSchedulesResponse);
  rpc GetSchedulesForExecution(GetSchedulesForExecutionRequest) returns (ListIrrigationSchedulesResponse);
}

// Request/Response messages
message CreateIrrigationScheduleRequest {
  string growing_zone_id = 1 [(buf.validate.field).string.uuid = true];
  string planting_cycle_id = 2 [(buf.validate.field).string.uuid = true];
  string schedule_name = 3 [(buf.validate.field).string = {
    min_len: 1,
    max_len: 255
  }];
  string irrigation_type = 4 [(buf.validate.field).string = {
    max_len: 100
  }];
  string start_time = 5 [(buf.validate.field).string.pattern = "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"];
  int32 duration_minutes = 6 [(buf.validate.field).int32 = {
    gte: 1,
    lte: 1440
  }];
  string frequency = 7 [(buf.validate.field).string = {
    max_len: 50
  }];
  string days_of_week = 8;
  double water_amount_liters = 9 [(buf.validate.field).double = {
    gte: 0,
    lte: 10000
  }];
  bool fertilizer_mix = 10;
  bool is_active = 11;
  string created_by = 12 [(buf.validate.field).string.uuid = true];
}

message GetIrrigationScheduleRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message UpdateIrrigationScheduleRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
  string growing_zone_id = 2 [(buf.validate.field).string.uuid = true];
  string planting_cycle_id = 3 [(buf.validate.field).string.uuid = true];
  string schedule_name = 4 [(buf.validate.field).string = {
    min_len: 1,
    max_len: 255
  }];
  string irrigation_type = 5 [(buf.validate.field).string = {
    max_len: 100
  }];
  string start_time = 6 [(buf.validate.field).string.pattern = "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$"];
  int32 duration_minutes = 7 [(buf.validate.field).int32 = {
    gte: 1,
    lte: 1440
  }];
  string frequency = 8 [(buf.validate.field).string = {
    max_len: 50
  }];
  string days_of_week = 9;
  double water_amount_liters = 10 [(buf.validate.field).double = {
    gte: 0,
    lte: 10000
  }];
  bool fertilizer_mix = 11;
  bool is_active = 12;
  google.protobuf.Timestamp last_executed = 13;
  google.protobuf.Timestamp next_execution = 14;
}

message DeleteIrrigationScheduleRequest {
  string id = 1 [(buf.validate.field).string.uuid = true];
}

message ListIrrigationSchedulesRequest {
  string growing_zone_id = 1 [(buf.validate.field).string.uuid = true];
  string planting_cycle_id = 2 [(buf.validate.field).string.uuid = true];
  string schedule_name = 3 [(buf.validate.field).string = {
    max_len: 255
  }];
  string irrigation_type = 4 [(buf.validate.field).string = {
    max_len: 100
  }];
  string frequency = 5 [(buf.validate.field).string = {
    max_len: 50
  }];
  bool is_active = 6;
  bool fertilizer_mix = 7;
  string created_by = 8 [(buf.validate.field).string.uuid = true];
  google.protobuf.Timestamp created_at_from = 9;
  google.protobuf.Timestamp created_at_to = 10;
  google.protobuf.Timestamp next_execution_from = 11;
  google.protobuf.Timestamp next_execution_to = 12;
  int32 page = 13 [(buf.validate.field).int32 = {
    gte: 1
  }];
  int32 limit = 14 [(buf.validate.field).int32 = {
    gte: 1,
    lte: 100
  }];
  string sort_by = 15 [(buf.validate.field).string = {
    max_len: 50
  }];
  string sort_order = 16 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
}

message GetSchedulesByGrowingZoneRequest {
  string growing_zone_id = 1 [(buf.validate.field).string.uuid = true];
  common.PaginationRequest pagination = 2;
}

message GetSchedulesForExecutionRequest {
  google.protobuf.Timestamp from_time = 1;
  google.protobuf.Timestamp to_time = 2;
  common.PaginationRequest pagination = 3;
}

message IrrigationScheduleResponse {
  string id = 1;
  string growing_zone_id = 2;
  string planting_cycle_id = 3;
  string schedule_name = 4;
  string irrigation_type = 5;
  string start_time = 6;
  int32 duration_minutes = 7;
  string frequency = 8;
  string days_of_week = 9;
  double water_amount_liters = 10;
  bool fertilizer_mix = 11;
  bool is_active = 12;
  google.protobuf.Timestamp last_executed = 13;
  google.protobuf.Timestamp next_execution = 14;
  string created_by = 15;
  google.protobuf.Timestamp created_at = 16;
  google.protobuf.Timestamp updated_at = 17;
}

message ListIrrigationSchedulesResponse {
  repeated IrrigationScheduleResponse irrigation_schedules = 1;
  common.PaginationResponse pagination = 2;
}
