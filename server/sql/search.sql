-- name: SearchResources :many
-- Searches in some database tables to find matching resources based on a "contains string" pattern (LIKE '%<arg>%')
WITH resources AS (
    SELECT
        'Logs' AS table_name,
        lower(l.log_uuid::text) AS identifier,
        concat(l.log_type, ' log on ', platform, ' (', t.name, ', ', e.name, ')') AS details
    FROM logs l
    LEFT JOIN environments e on e.id = l.environment
    LEFT JOIN titles t on t.id = e.title

    UNION ALL

    SELECT
        'Environment',
        lower(e.key::text),
        concat(e.name, ' environment for ', t.name)
    FROM environments e
    LEFT JOIN titles t on t.id = e.title
)
SELECT table_name, identifier, details::text
FROM resources
WHERE identifier LIKE '%' || lower(@search) || '%'
LIMIT sqlc.arg('limit');
