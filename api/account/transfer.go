package account

import (
	"github.com/gin-gonic/gin"
	models "github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services"
	"github.com/prakharporwal/bank-server/services/klog"
	"log"
	"net/http"
)

type transferRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required"`
	ToAccountId   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}

const (
	DEBIT   = "DEBIT"
	CREDIT  = "CREDIT"
	MESSAGE = "message"
)

const (
	IndianRupees = "INR"
)

// TODO : Requires Authentication ADD its
//func TransferTx(ctx *gin.Context) {
//
//	var transaction transferRequest
//	err := ctx.ShouldBindJSON(&transaction)
//
//	if err != nil {
//		log.Println("Error Parsing! Invalid format", err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//		return
//	}
//	// Validate Data
//	if transaction.Amount <= 0 {
//		klog.Error("Illegal! Amount should be greater than zero")
//		ctx.JSON(http.StatusBadRequest, gin.H{MESSAGE: "Amount should be positive!"})
//	}
//
//	//recordStatement := "INSERT INTO transactions(transaction_id, from_account_id,to_account_id,amount) VALUES($1,$2,$3,$4)"
//	txnId := utils.GenerateTimeStampMicro()
//
//	arg := models.CreateTransferRecordParams{
//		TransactionID: txnId,
//		FromAccountID: transaction.FromAccountId,
//		ToAccountID:   transaction.ToAccountId,
//		Amount:        transaction.Amount,
//	}
//	_, err = store.GetInstance().CreateTransferRecord(ctx, arg)
//	if err != nil {
//		// Incase we find any error in the query execution, rollback the transaction
//		klog.Error("Failed executing recording transactions !", err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//		return
//	}
//
//	//// SEND A REQUEST TO SQS FOR
//	//go server.Deposit()
//	withdraw := WithdrawRequestParams{
//		TransactionId:     txnId,
//		AccountId:         transaction.FromAccountId,
//		Amount:            transaction.Amount,
//		WithdrawAccountId: transaction.ToAccountId,
//		Currency:          transaction.Currency,
//	}
//	controller.Withdraw(withdraw)
//
//	deposit := DepositRequestParams{
//		TransactionId:      txnId,
//		AccountId:          transaction.ToAccountId,
//		Amount:             transaction.Amount,
//		DepositToAccountId: transaction.FromAccountId,
//		Currency:           transaction.Currency,
//	}
//	controller.Deposit(deposit)
//
//	//tx := server.store.BeginTx(ctx, &sql.TxOptions{})
//	//senderRecordStatement := "INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount,type) VALUES($1,$2,$3,$4,$5)"
//	//_, err = tx.ExecContext(ctx, senderRecordStatement, txnId, transaction.FromAccountId, transaction.ToAccountId, transaction.Amount, DEBIT)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	klog.Error("Failed executing record statement query for sender!", err)
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//
//	//receiverRecordStatement := "INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount, type) VALUES($1,$2,$3,$4,$5)"
//	//_, err = tx.ExecContext(ctx, receiverRecordStatement, txnId, transaction.ToAccountId, transaction.FromAccountId, transaction.Amount, CREDIT)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	klog.Error("Failed executing record statement query for receiver!", err)
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//	//	return
//	//}
//
//	//// get current balance
//	//senderBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
//	//rows, err := tx.QueryContext(ctx, senderBalanceQuery, transaction.FromAccountID)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	klog.Error("Error failed to query sender balance !", err)
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//	//defer rows.Close()
//	//
//	//var senderBalance int
//	//for rows.Next() {
//	//	if err := rows.Scan(&senderBalance); err != nil {
//	//		// Check for a scan error.
//	//		// Query rows will be closed with defer.
//	//		log.Fatal(err)
//	//	}
//	//}
//	//
//	//// if balance less than zero don't allow transaction fail!
//	//if senderBalance < transaction.Amount {
//	//	log.Println("Low Balance transaction declined!")
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Low Balance. Transaction Failed!"})
//	//	return
//	//}
//	//
//	//deductStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
//	//_, err = tx.ExecContext(ctx, deductStatement, senderBalance-transaction.Amount, transaction.FromAccountId)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	klog.Error("Failed deduction query!", err)
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//	//
//	//// get receiver balance and update
//	//receiverBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
//	//rows, err = tx.QueryContext(ctx, receiverBalanceQuery, transaction.ToAccountId)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	log.Println("Error failed to query receiver balance !")
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//	//defer rows.Close()
//	//
//	//var receiverBalance int64
//	//for rows.Next() {
//	//	if err := rows.Scan(&receiverBalance); err != nil {
//	//		// Check for a scan error.
//	//		// Query rows will be closed with defer.
//	//		log.Fatal(err)
//	//	}
//	//}
//	//
//	//incrementStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
//	//_, err = tx.ExecContext(ctx, incrementStatement, receiverBalance+transaction.Amount, transaction.ToAccountId)
//	//if err != nil {
//	//	log.Println("Failed Increment query!")
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//
//	//err = tx.Commit()
//	//if err != nil {
//	//	log.Println("transaction commit failed!", err)
//	//	ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//	//	return
//	//}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction Initiated!"})
//}

func TransferTx(ctx *gin.Context) {
	var transaction transferRequest
	err := ctx.ShouldBindJSON(&transaction)

	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}
	// Validate Data
	if transaction.Amount <= 0 {
		klog.Error("Illegal! Amount should be greater than zero")
		ctx.JSON(http.StatusBadRequest, gin.H{MESSAGE: "Amount should be positive!"})
		return
	}
	account, _ := store.GetInstance().GetAccountById(ctx, transaction.FromAccountId)

	if account.Balance <= 0 {
		klog.Error("Balance Low for user!", transaction.FromAccountId)
		ctx.JSON(http.StatusBadRequest, gin.H{MESSAGE: "Balance Low!"})
		return
	}
	args := models.CreateTransferRecordParams{
		FromAccountID: transaction.FromAccountId,
		ToAccountID:   transaction.ToAccountId,
		Amount:        transaction.Amount,
		Currency:      IndianRupees,
	}

	result, err := services.TransferService(ctx, args)

	if err != nil {
		klog.Error("Failed Transferring Amount ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{MESSAGE: "Failed"})
		return
	}
	klog.Debug(result)

	ctx.JSON(http.StatusOK, gin.H{MESSAGE: "success"})
}
