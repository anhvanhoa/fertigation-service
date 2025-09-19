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
