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

-- Add trigger for updated_at
CREATE TRIGGER trigger_fertilizer_schedules_updated_at
    BEFORE UPDATE ON fertilizer_schedules
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
