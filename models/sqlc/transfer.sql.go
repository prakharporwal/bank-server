// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransferEntry = `-- name: CreateTransferEntry :one
INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount, type)
VALUES($1,$2,$3,$4,$5)
RETURNING id, transaction_id, account_id, other_account, amount, type, created_at
`

type CreateTransferEntryParams struct {
	TransactionID int64  `json:"transaction_id"`
	AccountID     int64  `json:"account_id"`
	OtherAccount  int64  `json:"other_account"`
	Amount        int64  `json:"amount"`
	Type          string `json:"type"`
}

func (q *Queries) CreateTransferEntry(ctx context.Context, arg CreateTransferEntryParams) (AccountTransactionsEntry, error) {
	row := q.db.QueryRowContext(ctx, createTransferEntry,
		arg.TransactionID,
		arg.AccountID,
		arg.OtherAccount,
		arg.Amount,
		arg.Type,
	)
	var i AccountTransactionsEntry
	err := row.Scan(
		&i.Uid,
		&i.TransactionID,
		&i.AccountID,
		&i.OtherAccount,
		&i.Amount,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}
