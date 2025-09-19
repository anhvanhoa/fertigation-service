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

-- Add trigger for updated_at
CREATE TRIGGER trigger_fertilizer_types_updated_at
    BEFORE UPDATE ON fertilizer_types
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
