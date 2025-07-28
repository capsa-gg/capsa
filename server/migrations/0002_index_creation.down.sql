-- Not adding BEGIN; COMMIT; as CONCURRENTLY cannot be used in transactions.

-- Drop filtering indexes
DROP INDEX CONCURRENTLY IF EXISTS idx_logs_log_uuid;
DROP INDEX CONCURRENTLY IF EXISTS idx_logs_platform_logtype;
DROP INDEX CONCURRENTLY IF EXISTS idx_environments_key;

-- Drop join indexes
DROP INDEX CONCURRENTLY IF EXISTS idx_logs_environment;
DROP INDEX CONCURRENTLY IF EXISTS idx_logs_chunks_log;
DROP INDEX CONCURRENTLY IF EXISTS idx_logs_links_source;
DROP INDEX CONCURRENTLY IF EXISTS idx_environments_title;
