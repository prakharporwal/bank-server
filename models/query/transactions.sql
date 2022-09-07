-- name: GetTransaction :one
SELECT * from transactions where uid=($1);

-- name: GetTransactionsList :many
SELECT * from transactions;