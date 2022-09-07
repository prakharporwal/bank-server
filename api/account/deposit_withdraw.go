package account

// POST
// deposit in bank
//func Deposit(reqBody DepositRequestParams) {
//
//	tx := controller.db.BeginTx(context.Background(), &sql.TxOptions{})
//
//	// get receiver balance and update
//	//receiverBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
//	//rows, err := tx.QueryContext(ctx, receiverBalanceQuery, reqBody.AccountId)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	log.Println("Error failed to query receiver balance !")
//	//	tx.Rollback()
//	//	//ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//	//	return
//	//}
//	//defer rows.Close()
//
//	//var receiverBalance int64
//	//for rows.Next() {
//	//	if err := rows.Scan(&receiverBalance); err != nil {
//	//		// Check for a scan error.
//	//		// Query rows will be closed with defer.
//	//		log.Fatal(err)
//	//	}
//	//}
//	receiverBalance, err := controller.db.GetBalanceByAccountId(context.Background(), reqBody.AccountId)
//	if err != nil {
//		klog.Error("failed getting receiver balance")
//		rollBackErr := tx.Rollback()
//		if rollBackErr != nil {
//			log.Fatal("Failed tx rollback! update drivers", rollBackErr)
//			return
//		}
//		return
//	}
//
//	arg := db.UpdateAccountBalanceByIdParams{
//		ID:      reqBody.AccountId,
//		Balance: receiverBalance + reqBody.Amount,
//	}
//	_, err = controller.db.WithTx(tx).UpdateAccountBalanceById(context.Background(), arg)
//	//incrementStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
//	//_, err = tx.ExecContext(ctx, incrementStatement, receiverBalance+reqBody.Amount, reqBody.AccountId)
//	if err != nil {
//		log.Println("Failed Increment query!")
//		// Incase we find any error in the query execution, rollback the transaction
//		rollBackErr := tx.Rollback()
//		if rollBackErr != nil {
//			log.Fatal("Failed tx rollback! update drivers", rollBackErr)
//			return
//		}
//		//ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//		return
//	}
//
//	args := db.CreateAccountStatementEntryParams{
//		TransactionID: reqBody.TransactionId,
//		Amount:        reqBody.Amount,
//		AccountID:     reqBody.AccountId,
//		OtherAccount:  reqBody.DepositToAccountId,
//		Type:          CREDIT,
//	}
//	//senderRecordStatement := "INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount,type) VALUES($1,$2,$3,$4,$5)"
//	//_, err = tx.ExecContext(ctx, senderRecordStatement, reqBody.TransactionId, reqBody.DepositToAccountId, reqBody.AccountId, reqBody.Amount, DEBIT)
//	_, err = controller.db.WithTx(tx).CreateAccountStatementEntry(context.Background(), args)
//	if err != nil {
//		// Incase we find any error in the query execution, rollback the transaction
//		klog.Error("Failed executing record statement query for sender!", err)
//		rollBackErr := tx.Rollback()
//		if rollBackErr != nil {
//			log.Fatal("Failed tx rollback! update drivers ", rollBackErr)
//			return
//		}
//		//ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
//		return
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		log.Println("transaction commit failed!", err)
//		//ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//		return
//	}
//	//ctx.JSON(http.StatusOK, gin.H{MESSAGE: "Transaction Successful!"})
//}
//
//type WithdrawRequestParams struct {
//	TransactionId     int64  `json:"transaction_id" binding:"required"`
//	AccountId         int64  `json:"account_id" binding:"required"`
//	WithdrawAccountId int64  `json:"withdraw_account_id"`
//	Amount            int64  `json:"amount" binding:"required"`
//	Currency          string `json:"currency" binding:"required"`
//}
//
//func (controller *api.TransactionController) Withdraw(reqBody WithdrawRequestParams) {
//	//var reqBody withdrawRequest
//	//err := ctx.ShouldBindJSON(reqBody)
//	//if err != nil {
//	//	ctx.JSON(http.StatusInternalServerError, apierror.InvalidRequestBody)
//	//	return
//	//}
//
//	tx := controller.db.BeginTx(context.Background(), &sql.TxOptions{})
//
//	senderBalance, err := controller.db.WithTx(tx).GetBalanceByAccountId(context.Background(), reqBody.AccountId)
//	if err != nil {
//		klog.Error("getting account balance failed!", err)
//		return
//	}
//	// get current balance
//	//senderBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
//	//rows, err := tx.QueryContext(ctx, senderBalanceQuery, reqBody.AccountId)
//	//if err != nil {
//	//	// Incase we find any error in the query execution, rollback the transaction
//	//	klog.Error("Error failed to query sender balance !", err)
//	//	tx.Rollback()
//	//	ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//	//	return
//	//}
//	//defer rows.Close()
//
//	//var senderBalance int64
//	//for rows.Next() {
//	//	if err := rows.Scan(&senderBalance); err != nil {
//	//		// Check for a scan error.
//	//		// Query rows will be closed with defer.
//	//		log.Fatal(err)
//	//	}
//	//}
//
//	// if balance less than zero don't allow transaction fail!
//	if senderBalance < reqBody.Amount {
//		log.Println("Low Balance Withdraw declined!")
//		//ctx.JSON(http.StatusBadRequest, gin.H{"message": "Low Balance. Transaction Failed!"})
//		return
//	}
//
//	arg := db.UpdateAccountBalanceByIdParams{
//		Balance: senderBalance - reqBody.Amount,
//		ID:      reqBody.AccountId,
//	}
//	// deduct amount
//	_, err = controller.db.WithTx(tx).UpdateAccountBalanceById(context.Background(), arg)
//	if err != nil {
//		klog.Error("updating account balance failed!", err)
//		return
//	}
//	//deductStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
//	//_, err = tx.ExecContext(ctx, deductStatement, senderBalance-reqBody.Amount, reqBody.AccountId)
//	if err != nil {
//		// Incase we find any error in the query execution, rollback the transaction
//		klog.Error("Failed deduction query!", err)
//		rollBackErr := tx.Rollback()
//		if rollBackErr != nil {
//			log.Fatal("Failed tx rollback! update drivers", rollBackErr)
//			return
//		}
//		//ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//		return
//	}
//	args := db.CreateAccountStatementEntryParams{
//		TransactionID: reqBody.TransactionId,
//		Amount:        reqBody.Amount,
//		AccountID:     reqBody.AccountId,
//		OtherAccount:  reqBody.WithdrawAccountId,
//		Type:          DEBIT,
//	}
//	_, err = controller.db.WithTx(tx).CreateAccountStatementEntry(context.Background(), args)
//	//receiverRecordStatement := "INSERT INTO account_transactions_entries(transaction_id,account_id,other_account, amount, type) VALUES($1,$2,$3,$4,$5)"
//	//_, err = tx.ExecContext(ctx, receiverRecordStatement, reqBody.TransactionId, reqBody.AccountId, reqBody.WithdrawAccountId, reqBody.Amount, CREDIT)
//	if err != nil {
//		// Incase we find any error in the query execution, rollback the transaction
//		klog.Error("Failed executing record statement query for receiver!", err)
//		rollBackErr := tx.Rollback()
//		if rollBackErr != nil {
//			log.Fatal("Failed tx rollback! update drivers", rollBackErr)
//			return
//		}
//		//ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//		return
//	}
//	err = tx.Commit()
//	if err != nil {
//		log.Println("withdraw commit failed!", err)
//		//ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
//		return
//	}
//
//	//ctx.JSON(http.StatusOK, gin.H{MESSAGE: "Money Withdrawn Successful!"})
//	return
//}
