-- Drop indexes
DROP INDEX IF EXISTS idx_irrigation_logs_status_date;
DROP INDEX IF EXISTS idx_irrigation_logs_device;
DROP INDEX IF EXISTS idx_irrigation_logs_schedule;

-- Drop table
DROP TABLE IF EXISTS irrigation_logs;
