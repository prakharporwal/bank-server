-- name: GetAccount :one
SELECT * FROM accounts
WHERE owner_email = $1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
OFFSET $1
LIMIT $2;

-- name: CreateAccount :one
INSERT INTO accounts (owner_email, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2
WHERE id = $1
RETURNING *;