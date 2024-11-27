-- name: AddUser :exec
-- Inserts new user into database without a password hash
INSERT INTO users (email, first_name, last_name, user_role)
VALUES ($1, $2, $3, $4);

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

SELECT
    u.user_uuid AS user_uuid,
    u.email AS email,
    U.first_name as first_name,
    U.last_name as last_name,
    u.user_role AS role,
    u.deactivated_on AS deactivated_ts,
    u.created_at AS created_at
FROM users u;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUuid :one
SELECT * FROM users
WHERE user_uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
