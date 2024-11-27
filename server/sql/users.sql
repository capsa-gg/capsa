-- name: AddUser :one
-- Inserts new user into database without a password hash
INSERT INTO users (email, first_name, last_name, user_role)
VALUES ($1, $2, $3, $4)
RETURNING user_uuid;

-- name: AddUserWithPassHash :one
-- Inserts new user into database with a password hash
INSERT INTO users (email, first_name, last_name, password_hash, password_uuid, user_role)
VALUES ($1, $2, $3, $4, uuid_generate_v4(), $5)
RETURNING user_uuid;

-- name: DeactivateUser :exec
-- Marks a user as deactivated
UPDATE users
SET deactivated_on = now(), password_uuid = NULL, password_hash = NULL
WHERE id = $1;

-- name: RemoveUserDeactivation :exec
-- Removes the user deactivation
UPDATE users
SET deactivated_on = NULL
WHERE id = $1;

-- name: UpdateUserPassword :exec
-- Update password_hash for user based on id.
-- Also resets the password_uuid to invalidate existing JWTs
UPDATE users
SET password_hash = $1, password_uuid = uuid_generate_v4()
WHERE id = $2;

-- name: ListUsers :many
-- Lists all users present in the database
SELECT
    u.user_uuid AS user_uuid,
    u.email AS email,
    U.first_name as first_name,
    U.last_name as last_name,
    u.user_role AS role,
    (u.password_uuid IS NOT NULL)::bool AS has_password_set,
    u.deactivated_on AS deactivated_ts,
    u.created_at AS created_at
FROM users u
WHERE ( u.user_uuid = sqlc.narg(filter_by_user_uuid) OR sqlc.narg(filter_by_user_uuid) IS NULL )
ORDER BY u.created_at DESC; -- Optionally filter by Log UUID;

-- name: UpdateUser :one
-- Update a user with optional parameters
UPDATE users
SET first_name = @first_name,
    last_name = @last_name,
    user_role = @user_role
WHERE id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUuid :one
SELECT * FROM users
WHERE user_uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
