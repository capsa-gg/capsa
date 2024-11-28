-- name: SearchResources :many
-- Searches in some database tables to find matching resources based on a "contains string" pattern (LIKE '%<arg>%')
WITH resources AS (
    SELECT
        'Environment' AS table_name,
        lower(e.key::text) AS identifier,
        concat(e.name, ' environment for ', t.name) AS description,
        '' AS details
    FROM environments e
    LEFT JOIN titles t ON t.id = e.title

    UNION ALL

    SELECT
        'Logs' ,
        lower(l.log_uuid::text) ,
        concat(l.log_type, ' log on ', platform, ' (', t.name, ', ', e.name, ')'),
        count(lc)::text
    FROM logs l
    LEFT JOIN environments e ON e.id = l.environment
    LEFT JOIN titles t on t.id = e.title
    LEFT JOIN logs_chunks lc ON l.id = lc.log
    GROUP BY l.log_uuid, l.log_type, platform, t.name, e.name
)
SELECT table_name, identifier, description::text, details::text
FROM resources
WHERE identifier LIKE '%' || lower(@search) || '%'
LIMIT sqlc.arg('limit');
