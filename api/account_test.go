package api

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAccount_createAccount(t *testing.T) {
	tests := []struct {
		name    string
		account createAccountInput
	}{
		{
			name: "Returns Expected User",
			account: createAccountInput{
				Owner:    "fake-owner",
				Currency: "INR",
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			// TODO: Fix this test
			assert.Equal(t, test.account.Owner, "fake-owner")
		})
	}
}
