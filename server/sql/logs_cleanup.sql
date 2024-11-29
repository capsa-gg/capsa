-- name: GetLogsBeforeTimestamp :many
-- Gets the logs that are before a given timestamp
SELECT id FROM logs
WHERE created_on < @before_ts::timestamp;

-- name: DeleteLogAndLinkedResources :exec
-- Deletes a log, including chunks, metadata and links via CASCADE.
DELETE FROM logs l WHERE id = @id::bigint;
