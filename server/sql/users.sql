-- name: AddUser :exec
-- Inserts new user into database
INSERT INTO users (email, first_name, last_name)
VALUES ($1, $2, $3);

-- name: UpdateUserPassword :exec
-- Update password_hash for user based on id.
-- Also resets the password_uuid to invalidate existing JWTs
UPDATE users
SET password_hash = $1, password_uuid = uuid_generate_v4()
WHERE id = $2;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUuid :one
SELECT * FROM users
WHERE user_uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
