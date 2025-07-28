-- name: ListAvailableLogs :many
-- Fetches all log chunks and aggregates an overview.
-- LogUUID is an optional field used as a filter, which if set will return only a single result.
WITH logs_to_fetch AS (
    SELECT l.id, l.log_uuid, l.platform, l.log_type, l.environment
    FROM logs l
    JOIN environments e ON l.environment = e.id
    JOIN titles t ON e.title = t.id
    WHERE ( l.log_uuid =  sqlc.narg(filter_by_log_uuid)          OR sqlc.narg(filter_by_log_uuid)    IS NULL )
      AND ( e.key =       sqlc.narg(filter_by_environment)::uuid OR sqlc.narg(filter_by_environment) IS NULL )
      AND ( l.platform =  sqlc.narg(filter_by_platform)::varchar OR sqlc.narg(filter_by_platform)    IS NULL )
      AND ( l.log_type =  sqlc.narg(filter_by_logtype)           OR sqlc.narg(filter_by_logtype)     IS NULL )
    ORDER BY l.created_on DESC
    LIMIT @fetchlimit::int
),
cat_counts AS (
    SELECT log, (jsb).key AS category, sum((jsb).value::int) AS count
    FROM (
        SELECT log, jsb
        FROM logs_chunks, jsonb_each_text(category_counts) AS jsb
        WHERE log IN (SELECT id FROM logs_to_fetch)
     ) AS combined_entries
    GROUP BY category, log
),
sev_counts AS (
    SELECT log, (jsb).key AS severity, sum((jsb).value::int) AS count
    FROM (
        SELECT log, jsb
        FROM logs_chunks, jsonb_each_text(severity_counts) AS jsb
        WHERE log IN (SELECT id FROM logs_to_fetch)
    ) AS combined_entries
    GROUP BY severity, log
),
chunk_data AS (
    SELECT
        log AS log,
        sum(line_count) AS line_count,
        count(*) AS chunk_count,
        MIN(chunk_start) AS earliest_start,
        MAX(chunk_end) AS latest_end
    FROM logs_chunks
    WHERE log IN (SELECT id FROM logs_to_fetch)
    GROUP BY log
),
links AS (
    SELECT
        source AS source,
        count(*) AS sum
    FROM logs_links
    WHERE source IN (SELECT id FROM logs_to_fetch)
    GROUP BY source
)
SELECT
    ltf.log_uuid AS log_uuid,
    ltf.platform AS platform,
    ltf.log_type AS log_type,
    t.name AS title,
    e.name AS environment,
    cd.line_count AS line_count,
    cd.chunk_count AS chunk_count,
    COALESCE(ll.sum, 0) AS link_count,
    cd.earliest_start AS earliest,
    cd.latest_end AS last,
    jsonb_object_agg(cc.category, cc.count) AS categories_count,
    jsonb_object_agg(sc.severity, sc.count) AS severities_count
FROM logs_to_fetch ltf
JOIN logs l ON l.id = ltf.id
JOIN environments e ON l.environment = e.id
JOIN titles t ON e.title = t.id
JOIN chunk_data cd ON cd.log = ltf.id
JOIN cat_counts cc ON cc.log = ltf.id
JOIN sev_counts sc ON sc.log = ltf.id
LEFT JOIN links ll ON ltf.id = ll.source
GROUP BY ltf.id, ltf.log_uuid, ltf.platform, ltf.log_type, ltf.environment, t.name, e.name, cd.line_count, cd.chunk_count, cd.earliest_start, cd.latest_end, ll.sum
ORDER BY cd.earliest_start DESC;
