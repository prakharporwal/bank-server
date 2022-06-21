-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
OFFSET $1
LIMIT $2;

-- name: CreateAccount :one
INSERT INTO accounts (owner_email, currency)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = $2
WHERE id = $1;