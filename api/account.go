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
	}

	tx := server.store.BeginTx(ctx, &sql.TxOptions{})

	recordStatement := "INSERT INTO transactions(from_account_id,to_account_id,amount) VALUES($1,$2,$3)"
	tx.ExecContext(ctx, recordStatement, transaction.FromAccountID, transaction.ToAccountID, transaction.Amount)

	senderRecordStatement := "INSERT INTO account_transactions_entries(account_id, amount) VALUES($1,$2)"
	tx.ExecContext(ctx, senderRecordStatement, transaction.FromAccountID, -transaction.Amount)

	receiverRecordStatement := "INSERT INTO account_transactions_entries(account_id, amount) VALUES($1,$2)"
	tx.ExecContext(ctx, receiverRecordStatement, transaction.ToAccountID, transaction.Amount)

	/******* WRAP IN A TRANSACTION ********/
	deductStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
	tx.ExecContext(ctx, deductStatement, -transaction.Amount, transaction.FromAccountID)

	incrementStatement := "UPDATE accounts SET balance=($1) WHERE id=($2)"
	tx.ExecContext(ctx, incrementStatement, transaction.Amount, transaction.ToAccountID)
	/******* WRAP IN A TRANSACTION ********/

	if err = tx.Commit(); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
