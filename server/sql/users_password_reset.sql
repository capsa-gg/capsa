-- name: InitializeUserPasswordReset :exec
-- Adds password_forgot entry for user.
-- In case there is a password reset (database unique conflict)
-- the existing entry will be reset with new values
INSERT INTO users_password_reset (user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE SET
    reset_token = DEFAULT,
    valid_until = DEFAULT;

-- name: GetPasswordResetByUserId :one
SELECT * FROM users_password_reset
WHERE user_id = $1;

-- name: GetPasswordResetByResetToken :one
SELECT * FROM users_password_reset
WHERE reset_token = $1;

-- name: DeletePasswordResetForUser :exec
-- Marks a password_forgot entry used based on user_id
DELETE FROM users_password_reset
WHERE user_id = $1;
