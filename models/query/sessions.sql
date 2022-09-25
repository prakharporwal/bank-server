-- name: CreateSession :one
INSERT INTO sessions(
    email ,
    user_agent,
    client_ip,
    refresh_token,
    expires_at
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions where session_id=($1);

-- name: BlockSessionById :one
UPDATE sessions SET is_blocked=true
WHERE session_id=($1)
RETURNING *;


-- name: BlockSessionFamily :one
UPDATE sessions SET is_blocked=true
WHERE email=($1)
RETURNING *;