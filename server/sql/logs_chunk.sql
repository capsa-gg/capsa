-- name: AddLogChunk :exec
-- Adds data for an uploaded log chunk
INSERT INTO logs_chunks(log, blob_path, chunk_start, chunk_end, line_count, category_counts, severity_counts)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetLogChunksForLog :many
-- Gets the log chunk information for a given log
SELECT *
FROM logs_chunks
WHERE log = $1
ORDER BY created_on;
