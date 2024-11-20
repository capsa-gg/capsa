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
