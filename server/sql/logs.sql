-- name: AddNewLogSession :one
-- Inserts new log session into the database
INSERT INTO logs (environment, log_type, platform)
VALUES ($1, $2, $3)
RETURNING log_uuid;

-- name: GetLogByUuid :one
SELECT * FROM logs
WHERE log_uuid = $1;

-- name: AddLogMetadata :exec
-- Inserts metadata for a given log
INSERT INTO logs_metadata(log, metadata)
VALUES ($1, $2);

-- name: AddLogLink :exec
-- Inserts a link between a source and another log session, with a description
INSERT INTO logs_links(source, link, description)
VALUES ($1, $2, $3)
ON CONFLICT (source, link)
DO UPDATE SET description = $3;

-- name: GetMetadataForLog :many
-- Fetches the stored additional metadata for the log
SELECT saved_on, metadata
FROM logs_metadata
WHERE log = $1
ORDER BY saved_on;

-- name: GetLinkedLogsForLog :many
-- Fetches the linked logs for a log
SELECT
    l.log_uuid AS linked_log,
    links.description AS description
FROM logs_links links
JOIN logs l ON links.link = l.id
WHERE source = $1
ORDER BY links.created_on;

-- name: AddLogChunk :exec
-- Adds data for an uploaded log chunk
INSERT INTO logs_chunks(log, blob_path, chunk_start, chunk_end, line_count, category_counts, severity_counts)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateLogTimestamps :exec
-- Updates the timestamps for a log based on the chunk metadata
UPDATE logs
SET log_start =
    CASE
        WHEN log_start IS NULL THEN @log_start::timestamp -- In case the timestamp is not set
        WHEN log_start < '2024-01-01' THEN @log_start::timestamp -- In case the timestamp is an empty time.Time{}
        ELSE log_start
    END,
    log_end = COALESCE(GREATEST(@log_end::timestamp, log_end), log_end)
WHERE log_uuid = $1;

-- name: GetLogChunksForLog :many
-- Gets the log chunk information for a given log
SELECT created_on, blob_path
FROM logs_chunks
WHERE log = $1
ORDER BY created_on;
