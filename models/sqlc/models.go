// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         int64     `json:"id"`
	OwnerEmail string    `json:"owner_email"`
	Balance    int64     `json:"balance"`
	Currency   string    `json:"currency"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AccountTransactionsEntry struct {
	Uid           uuid.UUID `json:"uid"`
	TransactionID int64     `json:"transaction_id"`
	AccountID     int64     `json:"account_id"`
	OtherAccount  int64     `json:"other_account"`
	// must be positive
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	SessionID    uuid.UUID `json:"session_id"`
	Email        string    `json:"email"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	IsBlocked    bool      `json:"is_blocked"`
	CreatedAt    time.Time `json:"created_at"`
}

type Transaction struct {
	Uid           uuid.UUID `json:"uid"`
	TransactionID int64     `json:"transaction_id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	// can be negative depending on debit or credit
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	UserID       int64     `json:"user_id"`
	Username     string    `json:"username"`
	UserEmail    string    `json:"user_email"`
	PasswordHash string    `json:"password_hash"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
