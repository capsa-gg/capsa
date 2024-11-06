-- name: AddNewLogSession :one
-- Inserts new log session into the database
INSERT INTO logs (environment, log_type)
VALUES ($1, $2)
RETURNING log_uuid;

-- name: GetLogByUuid :one
SELECT * FROM logs
WHERE log_uuid = $1;
