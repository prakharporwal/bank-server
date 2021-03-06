// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( user_id, username, user_email, password_hash, is_verified)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, username, user_email, is_verified
`

type CreateUserParams struct {
	UserID       int64  `json:"user_id"`
	Username     string `json:"username"`
	UserEmail    string `json:"user_email"`
	PasswordHash string `json:"password_hash"`
	IsVerified   bool   `json:"is_verified"`
}

type CreateUserRow struct {
	UserID     int64  `json:"user_id"`
	Username   string `json:"username"`
	UserEmail  string `json:"user_email"`
	IsVerified bool   `json:"is_verified"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserID,
		arg.Username,
		arg.UserEmail,
		arg.PasswordHash,
		arg.IsVerified,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.UserEmail,
		&i.IsVerified,
	)
	return i, err
}

const getUserDetails = `-- name: GetUserDetails :one
SELECT user_email, username, password_hash
FROM users WHERE user_email=($1) or username=($1)
`

type GetUserDetailsRow struct {
	UserEmail    string `json:"user_email"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) GetUserDetails(ctx context.Context, userEmail string) (GetUserDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDetails, userEmail)
	var i GetUserDetailsRow
	err := row.Scan(&i.UserEmail, &i.Username, &i.PasswordHash)
	return i, err
}
