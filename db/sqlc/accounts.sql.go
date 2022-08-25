// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (owner_email, balance, currency)
VALUES ($1, $2, $3)
RETURNING id, owner_email, balance, currency, created_at, updated_at
`

type CreateAccountParams struct {
	OwnerEmail string `json:"owner_email"`
	Balance    int64  `json:"balance"`
	Currency   string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.OwnerEmail, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.OwnerEmail,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createAccountStatementEntry = `-- name: CreateAccountStatementEntry :one
INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount, type)
VALUES($1,$2,$3,$4,$5)
RETURNING id, transaction_id, account_id, other_account, amount, type, created_at
`

type CreateAccountStatementEntryParams struct {
	TransactionID int64  `json:"transaction_id"`
	AccountID     int64  `json:"account_id"`
	OtherAccount  int64  `json:"other_account"`
	Amount        int64  `json:"amount"`
	Type          string `json:"type"`
}

func (q *Queries) CreateAccountStatementEntry(ctx context.Context, arg CreateAccountStatementEntryParams) (AccountTransactionsEntry, error) {
	row := q.db.QueryRowContext(ctx, createAccountStatementEntry,
		arg.TransactionID,
		arg.AccountID,
		arg.OtherAccount,
		arg.Amount,
		arg.Type,
	)
	var i AccountTransactionsEntry
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.AccountID,
		&i.OtherAccount,
		&i.Amount,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}

const createTransferRecord = `-- name: CreateTransferRecord :one
INSERT INTO transactions(transaction_id, from_account_id,to_account_id,amount)
VALUES($1,$2,$3,$4)
RETURNING id, transaction_id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferRecordParams struct {
	TransactionID int64 `json:"transaction_id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransferRecord(ctx context.Context, arg CreateTransferRecordParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransferRecord,
		arg.TransactionID,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccountById = `-- name: GetAccountById :one
SELECT id, owner_email, balance, currency, created_at, updated_at FROM accounts
WHERE id = $1
`

func (q *Queries) GetAccountById(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.OwnerEmail,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByOwnerEmail = `-- name: GetAccountByOwnerEmail :one
SELECT id, owner_email, balance, currency, created_at, updated_at FROM accounts
WHERE owner_email = $1
`

func (q *Queries) GetAccountByOwnerEmail(ctx context.Context, ownerEmail string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByOwnerEmail, ownerEmail)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.OwnerEmail,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBalanceByAccountId = `-- name: GetBalanceByAccountId :one
SELECT balance FROM accounts
WHERE id=($1)
`

func (q *Queries) GetBalanceByAccountId(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBalanceByAccountId, id)
	var balance int64
	err := row.Scan(&balance)
	return balance, err
}

const getBalanceByOwnerEmail = `-- name: GetBalanceByOwnerEmail :one
SELECT balance FROM accounts
WHERE owner_email=($1)
`

func (q *Queries) GetBalanceByOwnerEmail(ctx context.Context, ownerEmail string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBalanceByOwnerEmail, ownerEmail)
	var balance int64
	err := row.Scan(&balance)
	return balance, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner_email, balance, currency, created_at, updated_at FROM accounts
ORDER BY id
OFFSET $1
LIMIT $2
`

type ListAccountsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.OwnerEmail,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountBalanceById = `-- name: UpdateAccountBalanceById :one
UPDATE accounts SET balance = $2
WHERE id = $1
RETURNING id, owner_email, balance, currency, created_at, updated_at
`

type UpdateAccountBalanceByIdParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (q *Queries) UpdateAccountBalanceById(ctx context.Context, arg UpdateAccountBalanceByIdParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBalanceById, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.OwnerEmail,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}