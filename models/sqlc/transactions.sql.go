// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transactions.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getTransaction = `-- name: GetTransaction :one
SELECT uid, transaction_id, from_account_id, to_account_id, amount, currency, created_at from transactions where uid=($1)
`

func (q *Queries) GetTransaction(ctx context.Context, uid uuid.UUID) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, uid)
	var i Transaction
	err := row.Scan(
		&i.Uid,
		&i.TransactionID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactionsList = `-- name: GetTransactionsList :many
SELECT uid, transaction_id, from_account_id, to_account_id, amount, currency, created_at from transactions
`

func (q *Queries) GetTransactionsList(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactionsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.Uid,
			&i.TransactionID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.Currency,
			&i.CreatedAt,
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