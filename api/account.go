package api

import (
	"context"
	"database/sql"
	"github.com/prakharporwal/bank-server/api/apierror"
	db "github.com/prakharporwal/bank-server/db/sqlc"
	"github.com/prakharporwal/bank-server/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	OwnerEmail string `json:"owner_email" binding:"required"`
	Currency   string `json:"currency" binding:"required,oneof=INR USD CAD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var request createAccountRequest

	err := ctx.ShouldBindJSON(&request)
	request.OwnerEmail = strings.ToLower(request.OwnerEmail)
	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, apierror.InvalidRequestBody)
		return
	}

	if !utils.IsValidEmail(request.OwnerEmail) {
		klog.Error("Invalid Email Address", request.OwnerEmail)
		ctx.JSON(http.StatusInternalServerError, gin.H{MESSAGE: "Invalid Email Address! Enter a valid email address!"})
		return
	}

	args := db.CreateAccountParams{
		OwnerEmail: request.OwnerEmail,
		Balance:    0,
		Currency:   request.Currency,
	}
	//statement := `INSERT INTO accounts (owner_email,currency,balance) VALUES ($1,$2,$3)`
	//server.store.Query(statement, account.OwnerEmail, account.Currency, 0) // initial balance will be 0

	server.store.CreateAccount(context.Background(), args)

	log.Println(request.OwnerEmail)
	log.Println(request.Currency)

	ctx.JSON(http.StatusOK, gin.H{MESSAGE: "Account Created for :" + request.OwnerEmail})
}

func (server *Server) GetAccount(ctx *gin.Context) {
	ownerEmail := ctx.Query("owner_email")
	accountId := ctx.Query("account_id")

	klog.Debug("owner email", ownerEmail)
	klog.Debug("account id", accountId)

	//statement := `SELECT * FROM accounts WHERE owner_email=($1)`
	//result := server.store.Query(statement, ownerEmail)
	//defer result.Close()
	var account db.Account
	var err error
	if accountId == "" && ownerEmail == "" {
		log.Println("Error Scanning Results !", err)
		ctx.JSON(http.StatusExpectationFailed, gin.H{MESSAGE: "pass request QUERY PARAM owner_email or account_id"})
		return
	} else if accountId != "" {
		id, _ := strconv.Atoi(accountId)
		account, err = server.store.GetAccountById(context.Background(), int64(id))
	} else {
		account, err = server.store.GetAccountByOwnerEmail(context.Background(), ownerEmail)
	}
	//var accounts model.Account
	//for result.Next() {
	//	var account model.Account
	//	if err := result.Scan(&account.Id, &account.OwnerEmail, &account.Balance, &account.Currency, &account.CreatedAt, &account.UpdatedAt); err != nil {
	//		// if err := result.Scan(account); err != nil {
	//		log.Println("Error Scanning All Results!", err)
	//		ctx.JSON(http.StatusAccepted, accounts)
	//		return
	//	}
	//	accounts = append(accounts, account)
	//}

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Account Not Found for Email : ", ownerEmail)
			ctx.JSON(http.StatusNotFound, apierror.NotFound)
			return
		}
		log.Println("Error Scanning Results !", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}

	log.Println("Accounts :", account)
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) ListAccount(ctx *gin.Context) {

	pageNum, _ := strconv.Atoi(ctx.Param("page"))

	//statement := `SELECT * FROM accounts ORDER BY id ASC OFFSET ($1) LIMIT 2`
	//result := server.store.Query(statement, (pageNum-1)*2)
	//defer result.Close()
	//
	//var accounts []model.Account
	//for result.Next() {
	//	var account model.Account
	//	err := result.Scan(&account.Id, &account.OwnerEmail, &account.Balance, &account.Currency, &account.CreatedAt, &account.UpdatedAt)
	//	if err != nil {
	//		log.Println("Error Scanning All Results!", err)
	//		ctx.JSON(http.StatusAccepted, accounts)
	//		return
	//	}
	//	accounts = append(accounts, account)
	//}

	//if err := result.Err(); err != nil {
	//	log.Println("Error Scanning Results!", err)
	//	ctx.JSON(http.StatusPartialContent, accounts)
	//}

	//log.Println("Accounts :", accounts)
	args := db.ListAccountsParams{
		Offset: int32((pageNum - 1) * pageSize),
		Limit:  pageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, args)
	if err != nil {
		klog.Error("listing accounts failed ", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
