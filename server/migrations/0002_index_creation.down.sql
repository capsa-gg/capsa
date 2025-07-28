BEGIN;

-- Drop filtering indexes
DROP INDEX IF EXISTS idx_logs_log_uuid;
DROP INDEX IF EXISTS idx_logs_platform_logtype;
DROP INDEX IF EXISTS idx_environments_key;

-- Drop join indexes
DROP INDEX IF EXISTS idx_logs_environment;
DROP INDEX IF EXISTS idx_logs_chunks_log;
DROP INDEX IF EXISTS idx_logs_links_source;
DROP INDEX IF EXISTS idx_environments_title;

COMMIT;