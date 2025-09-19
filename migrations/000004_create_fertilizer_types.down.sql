-- Drop trigger first
DROP TRIGGER IF EXISTS trigger_fertilizer_types_updated_at ON fertilizer_types;

-- Drop indexes
DROP INDEX IF EXISTS idx_fertilizer_types_status;
DROP INDEX IF EXISTS idx_fertilizer_types_method;
DROP INDEX IF EXISTS idx_fertilizer_types_type;

-- Drop table
DROP TABLE IF EXISTS fertilizer_types;
