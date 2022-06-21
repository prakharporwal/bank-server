package db

import (
	"context"
	"github.com/prakharporwal/bank-server/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		OwnerEmail: utils.RandomEmail(),
		Balance:    utils.RandomNumber64(),
		Currency:   utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.OwnerEmail, account.OwnerEmail)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestQueries_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.OwnerEmail)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.OwnerEmail, account2.OwnerEmail)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: utils.RandomNumber64(),
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, account.OwnerEmail, updatedAccount.OwnerEmail)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.NotEqual(t, account.UpdatedAt, updatedAccount.UpdatedAt)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)
}
