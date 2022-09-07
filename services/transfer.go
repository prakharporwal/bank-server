package services

import (
	"context"
	models "github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/utils"
)

type TransactionTxResult struct {
	Transfer    models.Transaction              `json:"transfer"`
	FromAccount models.Account                  `json:"from_account"`
	ToAccount   models.Account                  `json:"to_account"`
	FromEntry   models.AccountTransactionsEntry `json:"from_entry"`
	ToEntry     models.AccountTransactionsEntry `json:"to_entry"`
}

// Executes a DB transaction for Transfer of money
// Should avoid deadlock and be atomic action
func TransferService(ctx context.Context, transaction models.CreateTransferRecordParams) (TransactionTxResult, error) {
	//recordStatement := "INSERT INTO transactions(transaction_id, from_account_id,to_account_id,amount) VALUES($1,$2,$3,$4)"
	txnId := utils.GenerateTimeStampMicro()

	transaction.TransactionID = txnId

	var result TransactionTxResult

	err := store.GetInstance().ExecTx(ctx, func(q *models.Queries) error {
		var err error

		result.Transfer, err = q.CreateTransferRecord(ctx, models.CreateTransferRecordParams{
			TransactionID: transaction.TransactionID,
			FromAccountID: transaction.FromAccountID,
			ToAccountID:   transaction.ToAccountID,
			Amount:        transaction.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateAccountStatementEntry(ctx, models.CreateAccountStatementEntryParams{
			TransactionID: transaction.TransactionID,
			AccountID:     transaction.FromAccountID,
			Amount:        -transaction.Amount,
			Type:          "DEBIT",
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateAccountStatementEntry(ctx, models.CreateAccountStatementEntryParams{
			TransactionID: transaction.TransactionID,
			AccountID:     transaction.ToAccountID,
			Amount:        transaction.Amount,
			Type:          "CREDIT",
		})
		if err != nil {
			return err
		}

		// move money into account2
		account1, err := q.GetAccountById(ctx, transaction.FromAccountID)
		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccountBalanceById(ctx, models.UpdateAccountBalanceByIdParams{
			ID:      transaction.FromAccountID,
			Balance: account1.Balance - transaction.Amount,
		})
		if err != nil {
			return err
		}

		// move money into account2
		account2, err := q.GetAccountById(ctx, transaction.ToAccountID)
		if err != nil {
			return err
		}

		result.ToAccount, err = q.UpdateAccountBalanceById(ctx, models.UpdateAccountBalanceByIdParams{
			ID:      transaction.ToAccountID,
			Balance: account2.Balance + transaction.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return TransactionTxResult{}, err
	}

	return result, nil
}
