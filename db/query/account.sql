-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountByOwnerEmail :one
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

-- name: UpdateAccountBalanceById :one
UPDATE accounts SET balance = $2
WHERE id = $1
RETURNING *;

-- name: GetBalanceByAccountId :one
SELECT balance FROM accounts
WHERE id=($1);

-- name: GetBalanceByOwnerEmail :one
SELECT balance FROM accounts
WHERE owner_email=($1);

-- name: CreateTransferRecord :one
INSERT INTO transactions(transaction_id, from_account_id,to_account_id,amount)
VALUES($1,$2,$3,$4)
RETURNING *;

-- name: CreateAccountStatementEntry :one
INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount, type)
VALUES($1,$2,$3,$4,$5)
RETURNING *;
