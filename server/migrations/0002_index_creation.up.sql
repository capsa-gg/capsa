BEGIN;

-- Filtering indexes
CREATE INDEX IF NOT EXISTS idx_logs_platform_logtype ON logs(platform, log_type);
CREATE UNIQUE INDEX IF NOT EXISTS idx_logs_log_uuid ON logs(log_uuid);
CREATE UNIQUE INDEX IF NOT EXISTS idx_environments_key ON environments(key);

-- Join indexes
CREATE INDEX IF NOT EXISTS idx_logs_environment ON logs(environment);
CREATE INDEX IF NOT EXISTS idx_logs_chunks_log ON logs_chunks(log);
CREATE INDEX IF NOT EXISTS idx_logs_links_source ON logs_links(source);
CREATE INDEX IF NOT EXISTS idx_environments_title ON environments(title);

COMMIT;