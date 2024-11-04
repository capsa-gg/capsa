-- name: AddTitle :exec
-- Inserts new title into database
INSERT INTO titles (name)
VALUES ($1);

-- name: GetTitleById :one
SELECT * FROM titles
WHERE id = $1;

-- name: GetTitleByName :one
SELECT * FROM titles
WHERE name = $1;

-- name: AddEnvironment :exec
-- Inserts new environment for a title into database
INSERT INTO environments (title, name)
VALUES ($1, $2);

-- name: GetEnvironmentById :one
SELECT * FROM environments
WHERE id = $1;

-- name: GetEnvironmentByTitleName :one
SELECT * FROM environments
WHERE title = $1
AND name = $2;

-- name: GetEnvironmentsForTitle :many
SELECT * FROM environments
WHERE title = $1;