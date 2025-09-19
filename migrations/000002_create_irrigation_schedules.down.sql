-- Drop trigger first
DROP TRIGGER IF EXISTS trigger_irrigation_schedules_updated_at ON irrigation_schedules;

-- Drop indexes
DROP INDEX IF EXISTS idx_irrigation_schedules_active;
DROP INDEX IF EXISTS idx_irrigation_schedules_next_exec;
DROP INDEX IF EXISTS idx_irrigation_schedules_cycle;
DROP INDEX IF EXISTS idx_irrigation_schedules_zone;

-- Drop table
DROP TABLE IF EXISTS irrigation_schedules;
