package models

import "time"

type Account struct {
	Id         int       `json:"id"`
	OwnerEmail string    `json:"owner_email"`
	Currency   string    `json:"currency"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Transaction struct {
	Id            int `json:"id"`
	FromAccountID int `json:"from_account_id"`
	ToAccountID   int `json:"to_account_id"`
	Amount        int `json:"amount"`
}
