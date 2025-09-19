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

-- Add trigger for updated_at
CREATE TRIGGER trigger_irrigation_schedules_updated_at
    BEFORE UPDATE ON irrigation_schedules
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
