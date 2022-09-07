package account

//import (
//	"context"
//	db "github.com/prakharporwal/bank-server/db/sqlc"
//	"github.com/stretchr/testify/require"
//	"testing"
//
//	"github.com/go-playground/assert/v2"
//)
//
//func TestAccount_createAccount(t *testing.T) {
//	tests := []struct {
//		name    string
//		account createAccountInput
//	}{
//		{
//			name: "Returns Expected User",
//			account: createAccountInput{
//				OwnerEmail: "fake-owner",
//				Currency:   "INR",
//			},
//		},
//	}
//
//	for _, test := range tests {
//
//		t.Run(test.name, func(t *testing.T) {
//			// TODO: Fix this test
//			assert.Equal(t, test.account.OwnerEmail, "fake-owner")
//		})
//	}
//}
//
//func createRandomAccount() db.Account {
//	account := db.Account{}
//	return account
//}
//
//func TestGetAccount(t *testing.T) {
//	account1 := createRandomAccount()
//	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
//
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//}
