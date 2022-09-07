package utils

import (
	"context"
	"github.com/golang/mock/gomock"
	models "github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomAccount(t *testing.T) models.Account {
	args := models.CreateAccountParams{
		OwnerEmail: RandomEmail(),
		Balance:    RandomNumber64(),
		Currency:   RandomCurrency(),
	}
	ctrl := gomock.NewController(t)
	mockStore := models.NewMockStore(ctrl)

	var account models.Account
	mockStore.EXPECT().CreateAccount(context.Background(), args).Times(1).Return(models.Account{
		ID:         GenerateTimeStampMicro(),
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
