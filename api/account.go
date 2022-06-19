package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/model"
)

type createAccountInput struct {
	OwnerEmail string `json:"owner_email" binding:"required"`
	Currency   string `json:"currency" binding:"required,oneof=INR USD CAD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var account createAccountInput

	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	statement := `INSERT INTO accounts (owner_email,currency,balance) VALUES ($1,$2,$3)`
	server.store.Query(statement, account.OwnerEmail, account.Currency, 0) // initial balance will be 0

	log.Println(account.OwnerEmail)
	log.Println(account.Currency)

	ctx.JSON(http.StatusOK, gin.H{"message": "Account " + account.OwnerEmail})
}

func (server *Server) GetAccount(ctx *gin.Context) {
	ownerEmail := ctx.Query("owner_email")

	log.Println(ownerEmail)
	statement := `SELECT * FROM accounts WHERE owner_email=($1)`

	result := server.store.Query(statement, ownerEmail)

	defer result.Close()

	var accounts []model.Account
	for result.Next() {
		var account model.Account
		if err := result.Scan(&account.Id, &account.OwnerEmail, &account.Balance, &account.Currency, &account.CreatedAt); err != nil {
			// if err := result.Scan(account); err != nil {
			log.Println("Error Scanning All Results!", err)
			ctx.JSON(http.StatusAccepted, accounts)
			return
		}
		accounts = append(accounts, account)
	}

	if err := result.Err(); err != nil {
		log.Println("Error Scanning Results!", err)
		ctx.JSON(http.StatusPartialContent, accounts)
	}

	log.Println("Accounts :", accounts)
	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) ListAccount(ctx *gin.Context) {

	pageNum, _ := strconv.Atoi(ctx.Param("page"))

	statement := `SELECT * FROM accounts ORDER BY id ASC OFFSET ($1) LIMIT 2`
	result := server.store.Query(statement, (pageNum-1)*2)
	defer result.Close()

	var accounts []model.Account
	for result.Next() {
		var account model.Account
		if err := result.Scan(&account.Id, &account.OwnerEmail, &account.Balance, &account.Currency, &account.CreatedAt); err != nil {
			log.Println("Error Scanning All Results!", err)
			ctx.JSON(http.StatusAccepted, accounts)
			return
		}
		accounts = append(accounts, account)
	}

	if err := result.Err(); err != nil {
		log.Println("Error Scanning Results!", err)
		ctx.JSON(http.StatusPartialContent, accounts)
	}

	log.Println("Accounts :", accounts)
	ctx.JSON(http.StatusOK, accounts)
}

type updateBalanceInput struct {
	FromAccountID int `json:"from_account_id" binding:"required"`
	ToAccountID   int `json:"to_account_id" binding:"required"`
	Amount        int `json:"amount" binding:"required"`
}

// TODO : Requires Authentication ADD its
func (server *Server) UpdateBalance(ctx *gin.Context) {

	var transaction updateBalanceInput
	err := ctx.ShouldBindJSON(&transaction)

	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	tx := server.store.BeginTx(ctx, &sql.TxOptions{})

	recordStatement := "INSERT INTO transactions(from_account_id,to_account_id,amount) VALUES($1,$2,$3)"
	_, err = tx.ExecContext(ctx, recordStatement, transaction.FromAccountID, transaction.ToAccountID, transaction.Amount)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		log.Println("Failed executing recording transactions !")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	senderRecordStatement := "INSERT INTO account_transactions_entries(account_id, amount) VALUES($1,$2)"
	_, err = tx.ExecContext(ctx, senderRecordStatement, transaction.FromAccountID, -transaction.Amount)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Println("Failed executing record statement query for sender!")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	receiverRecordStatement := "INSERT INTO account_transactions_entries(account_id, amount) VALUES($1,$2)"
	_, err = tx.ExecContext(ctx, receiverRecordStatement, transaction.ToAccountID, transaction.Amount)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Println("Failed executing record statement query for receiver!")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	// get current balance
	senderBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
	rows, err := tx.QueryContext(ctx, senderBalanceQuery, transaction.FromAccountID)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Println("Error failed to query sender balance !")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}
	defer rows.Close()

	var senderBalance int
	for rows.Next() {
		if err := rows.Scan(&senderBalance); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
	}

	// if balance less than zero don't allow transaction fail!
	if senderBalance < transaction.Amount {
		log.Println("Low Balance transaction declined!")
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Low Balance. Transaction Failed!"})
		return
	}

	deductStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
	_, err = tx.ExecContext(ctx, deductStatement, senderBalance-transaction.Amount, transaction.FromAccountID)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Println("Failed deduction query!")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	// get receiver balance and update
	receiverBalanceQuery := "SELECT balance FROM accounts WHERE id=($1)"
	rows, err = tx.QueryContext(ctx, receiverBalanceQuery, transaction.ToAccountID)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Println("Error failed to query receiver balance !")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}
	defer rows.Close()

	var receiverBalance int
	for rows.Next() {
		if err := rows.Scan(&receiverBalance); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
	}

	incrementStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
	_, err = tx.ExecContext(ctx, incrementStatement, receiverBalance+transaction.Amount, transaction.ToAccountID)
	if err != nil {
		log.Println("Failed Increment query!")
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	if err = tx.Commit(); err != nil {
		log.Println("transaction commit failed!", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction Successful!"})
}
