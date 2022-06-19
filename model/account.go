package model

import "time"

type Account struct {
	Id         int       `json:"id"`
	OwnerEmail string    `json:"owner_email"`
	Currency   string    `json:"currency"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}
