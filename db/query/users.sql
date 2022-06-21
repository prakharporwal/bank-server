-- name: CreateUser :one
INSERT INTO users ( user_id, username, user_email, password_hash, is_verified)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, username, user_email, is_verified;

-- name: GetUserDetails :one
SELECT user_email, username, password_hash
FROM users WHERE user_email=($1) or username=($1);