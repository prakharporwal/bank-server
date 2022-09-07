package db

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/prakharporwal/bank-server/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		OwnerEmail: utils.RandomEmail(),
		Balance:    utils.RandomNumber64(),
		Currency:   utils.RandomCurrency(),
	}
	ctrl := gomock.NewController(t)
	mockStore := NewMockStore(ctrl)

	var account Account
	mockStore.EXPECT().CreateAccount(context.Background(), args).Times(1).Return(Account{
		ID:         utils.GenerateTimeStampMicro(),
		OwnerEmail: args.OwnerEmail,
		Balance:    args.Balance,
		Currency:   args.Currency,
		CreatedAt:  time.Now().Local(),
		UpdatedAt:  time.Now().Local(),
	}, nil)

	account, err := mockStore.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.OwnerEmail, account.OwnerEmail)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestQueries_GetAccountByOwnerEmail(t *testing.T) {
	account1 := utils.CreateRandomAccount(t)
	account2, err := testQueries.GetAccountByOwnerEmail(context.Background(), account1.OwnerEmail)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.OwnerEmail, account2.OwnerEmail)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestQueries_UpdateAccountBalanceById(t *testing.T) {
	account := utils.CreateRandomAccount(t)

	args := UpdateAccountBalanceByIdParams{
		ID:      account.ID,
		Balance: utils.RandomNumber64(),
	}
	updatedAccount, err := testQueries.UpdateAccountBalanceById(context.Background(), args)

	require.NoError(t, err)
	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, account.OwnerEmail, updatedAccount.OwnerEmail)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.NotEqual(t, account.UpdatedAt, updatedAccount.UpdatedAt)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestQueries_MockUpdateAccountBalanceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStore(ctrl)
	account := utils.CreateRandomAccount(t)

	args := UpdateAccountBalanceByIdParams{
		ID:      account.ID,
		Balance: utils.RandomNumber64(),
	}

	// mocking setup
	//updatedAccount, err := testQueries.UpdateAccountBalanceById(context.Background(), args)

	mockStore.EXPECT().UpdateAccountBalanceById(context.Background(), args).Times(1).Return(Account{
		ID:         account.ID,
		OwnerEmail: account.OwnerEmail,
		Balance:    args.Balance,
		Currency:   account.Currency,
		CreatedAt:  account.CreatedAt,
		UpdatedAt:  time.Now().Local(),
	}, nil)
	updatedAccount, err := mockStore.UpdateAccountBalanceById(context.Background(), args)

	require.NoError(t, err)

	require.Equal(t, account.ID, updatedAccount.ID, "Account ID is not expected to change!")
	require.Equal(t, args.Balance, updatedAccount.Balance, "Balance did not change")
	require.Equal(t, account.OwnerEmail, updatedAccount.OwnerEmail)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.NotEqual(t, account.UpdatedAt, updatedAccount.UpdatedAt)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)
}
