-- 7. BẢNG LỊCH TƯỚI NƯỚC
CREATE TABLE irrigation_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    growing_zone_id UUID,
    planting_cycle_id UUID,
    schedule_name VARCHAR(255) NOT NULL,
    irrigation_type VARCHAR(100),
    start_time TIME,
    duration_minutes INTEGER,
    frequency VARCHAR(50),
    days_of_week JSONB,
    water_amount_liters NUMERIC(10,2),
    fertilizer_mix BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    last_executed TIMESTAMP,
    next_execution TIMESTAMP,
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (growing_zone_id) REFERENCES growing_zones(id) ON DELETE CASCADE,
    FOREIGN KEY (planting_cycle_id) REFERENCES planting_cycles(id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_irrigation_schedules_zone ON irrigation_schedules(growing_zone_id);
CREATE INDEX idx_irrigation_schedules_cycle ON irrigation_schedules(planting_cycle_id);
CREATE INDEX idx_irrigation_schedules_next_exec ON irrigation_schedules(next_execution);
CREATE INDEX idx_irrigation_schedules_active ON irrigation_schedules(is_active);

COMMENT ON COLUMN irrigation_schedules.irrigation_type IS 'manual, automatic, drip, spray, flood';
COMMENT ON COLUMN irrigation_schedules.frequency IS 'daily, weekly, bi_weekly, custom';
COMMENT ON COLUMN irrigation_schedules.days_of_week IS 'Array: ["monday", "tuesday", ...]';

-- 8. BẢNG NHẬT KÝ TƯỚI NƯỚC
CREATE TABLE irrigation_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    irrigation_schedule_id UUID,
    device_id UUID,
    started_at TIMESTAMP,
    ended_at TIMESTAMP,
    planned_duration_minutes INTEGER,
    actual_duration_minutes INTEGER,
    water_used_liters NUMERIC(10,2),
    water_pressure NUMERIC(5,2),
    status VARCHAR(50),
    failure_reason VARCHAR(500),
    notes TEXT,
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (irrigation_schedule_id) REFERENCES irrigation_schedules(id),
    FOREIGN KEY (device_id) REFERENCES iot_devices(id),
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_irrigation_logs_schedule ON irrigation_logs(irrigation_schedule_id);
CREATE INDEX idx_irrigation_logs_device ON irrigation_logs(device_id);
CREATE INDEX idx_irrigation_logs_status_date ON irrigation_logs(status, started_at);

COMMENT ON COLUMN irrigation_logs.water_pressure IS 'Áp suất nước (bar)';
COMMENT ON COLUMN irrigation_logs.status IS 'completed, failed, interrupted, manual_override';

-- 9. BẢNG LOẠI PHÂN BÓN
CREATE TABLE fertilizer_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100),
    npk_ratio VARCHAR(20),
    nitrogen_percentage NUMERIC(5,2),
    phosphorus_percentage NUMERIC(5,2),
    potassium_percentage NUMERIC(5,2),
    trace_elements JSONB,
    application_method VARCHAR(100),
    dosage_per_plant NUMERIC(10,4),
    dosage_per_m2 NUMERIC(10,4),
    unit VARCHAR(20),
    manufacturer VARCHAR(255),
    batch_number VARCHAR(100),
    expiry_date DATE,
    cost_per_unit NUMERIC(10,2),
    description TEXT,
    safety_notes TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_fertilizer_types_type ON fertilizer_types(type);
CREATE INDEX idx_fertilizer_types_method ON fertilizer_types(application_method);
CREATE INDEX idx_fertilizer_types_status ON fertilizer_types(status);

COMMENT ON COLUMN fertilizer_types.type IS 'organic, chemical, liquid, granular, powder';
COMMENT ON COLUMN fertilizer_types.npk_ratio IS 'e.g., "10-10-10", "20-20-20"';
COMMENT ON COLUMN fertilizer_types.trace_elements IS 'Các nguyên tố vi lượng';
COMMENT ON COLUMN fertilizer_types.application_method IS 'foliar, soil, hydroponic, fertigation';

-- 10. BẢNG LỊCH BÓN PHÂN
CREATE TABLE fertilizer_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    planting_cycle_id UUID NOT NULL,
    fertilizer_type_id UUID NOT NULL,
    application_date DATE,
    dosage NUMERIC(10,4),
    unit VARCHAR(20),
    application_method VARCHAR(100),
    growth_stage VARCHAR(100),
    weather_conditions VARCHAR(200),
    soil_conditions VARCHAR(200),
    is_completed BOOLEAN DEFAULT FALSE,
    completed_date DATE,
    actual_dosage NUMERIC(10,4),
    effectiveness_rating INTEGER,
    notes TEXT,
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (planting_cycle_id) REFERENCES planting_cycles(id) ON DELETE CASCADE,
    FOREIGN KEY (fertilizer_type_id) REFERENCES fertilizer_types(id),
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_fertilizer_schedules_cycle ON fertilizer_schedules(planting_cycle_id);
CREATE INDEX idx_fertilizer_schedules_type ON fertilizer_schedules(fertilizer_type_id);
CREATE INDEX idx_fertilizer_schedules_date ON fertilizer_schedules(application_date);
CREATE INDEX idx_fertilizer_schedules_completed ON fertilizer_schedules(is_completed);

COMMENT ON COLUMN fertilizer_schedules.growth_stage IS 'seedling, vegetative, flowering, fruiting, pre_harvest';
COMMENT ON COLUMN fertilizer_schedules.effectiveness_rating IS '1-5 rating';
