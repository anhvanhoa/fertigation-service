-- Drop trigger first
DROP TRIGGER IF EXISTS trigger_fertilizer_schedules_updated_at ON fertilizer_schedules;

-- Drop indexes
DROP INDEX IF EXISTS idx_fertilizer_schedules_completed;
DROP INDEX IF EXISTS idx_fertilizer_schedules_date;
DROP INDEX IF EXISTS idx_fertilizer_schedules_type;
DROP INDEX IF EXISTS idx_fertilizer_schedules_cycle;

-- Drop table
DROP TABLE IF EXISTS fertilizer_schedules;
