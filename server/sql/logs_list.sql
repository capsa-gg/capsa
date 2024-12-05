-- name: ListAvailableLogs :many
-- Fetches all log chunks and aggregates an overview.
-- LogUUID is an optional field used as a filter, which if set will return only a single result.
WITH cat_counts AS (
    SELECT log, (jsb).key AS category, sum((jsb).value::int) AS count
    FROM (
        SELECT log, jsb
        FROM logs_chunks, jsonb_each_text(category_counts) AS jsb
     ) AS combined_entries
    GROUP BY category, log
    ORDER BY category
),
sev_counts AS (
    SELECT log, (jsb).key AS severity, sum((jsb).value::int) AS count
    FROM (
        SELECT log, jsb
        FROM logs_chunks, jsonb_each_text(severity_counts) AS jsb
    ) AS combined_entries
    GROUP BY severity, log
    ORDER BY severity
),
chunk_data AS (
    SELECT
        log as log,
        sum(line_count) AS line_count,
        count(*) AS chunk_count,
        MIN(chunk_start) AS earliest_start,
        MAX(chunk_end) AS latest_end
    FROM logs_chunks
    GROUP BY log
)
SELECT
    l.log_uuid AS log_uuid,
    l.platform AS platform,
    l.log_type AS log_type,
    t.name AS title,
    e.name AS environment,
    cd.line_count AS line_count,
    cd.chunk_count AS chunk_count,
    cd.earliest_start AS earliest,
    cd.latest_end AS last,
    jsonb_object_agg(cc.category, cc.count) AS categories_count,
    jsonb_object_agg(sc.severity, sc.count) AS severities_count
FROM logs l
JOIN cat_counts cc ON cc.log = l.id
JOIN sev_counts sc ON sc.log = l.id
JOIN chunk_data cd ON cd.log = l.id
JOIN environments e on l.environment = e.id
JOIN titles t on e.title = t.id
WHERE ( l.log_uuid =  sqlc.narg(filter_by_log_uuid)          OR sqlc.narg(filter_by_log_uuid)    IS NULL )  -- Optionally filter by Log UUID
AND   ( e.key =       sqlc.narg(filter_by_environment)::uuid OR sqlc.narg(filter_by_environment) IS NULL )  -- Optionally filter by Environment
AND   ( l.platform =  sqlc.narg(filter_by_platform)::varchar OR sqlc.narg(filter_by_platform)    IS NULL )  -- Optionally filter by Platform
AND   ( l.log_type =  sqlc.narg(filter_by_logtype)           OR sqlc.narg(filter_by_logtype)     IS NULL )  -- Optionally filter by LogType
GROUP BY l.id, t.name, e.name, cd.line_count, cd.chunk_count, cd.earliest_start, cd.latest_end
ORDER BY earliest DESC
LIMIT @fetchlimit::int;
