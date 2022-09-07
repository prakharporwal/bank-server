package account

import (
	"context"
	"database/sql"
	"github.com/prakharporwal/bank-server/api/apierror"
	models "github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"

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

func CreateAccount(ctx *gin.Context) {
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

	args := models.CreateAccountParams{
		OwnerEmail: request.OwnerEmail,
		Balance:    0,
		Currency:   request.Currency,
	}

	_, err = store.GetInstance().CreateAccount(context.Background(), args)
	if err != nil {
		klog.Error("creating account failed", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}

	klog.Debug(request.OwnerEmail)
	klog.Debug(request.Currency)

	ctx.JSON(http.StatusOK, gin.H{MESSAGE: "Account Created for :" + request.OwnerEmail})
}

func GetAccount(ctx *gin.Context) {
	ownerEmail := ctx.Query("owner_email")
	accountId := ctx.Query("account_id")

	klog.Debug("owner email", ownerEmail)
	klog.Debug("account id", accountId)
	conn := store.GetInstance()

	var account models.Account
	var err error

	if accountId == "" && ownerEmail == "" {
		log.Println("Error Scanning Results !", err)
		ctx.JSON(http.StatusExpectationFailed, gin.H{MESSAGE: "pass request QUERY PARAM owner_email or account_id"})
		return
	} else if accountId != "" {
		id, _ := strconv.Atoi(accountId)
		account, err = conn.GetAccountById(context.Background(), int64(id))
	} else {
		account, err = conn.GetAccountByOwnerEmail(context.Background(), ownerEmail)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Account Not Found for : ", ownerEmail, accountId)
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

func ListAccount(ctx *gin.Context) {
	conn := store.GetInstance()
	pageNum, _ := strconv.Atoi(ctx.Param("page"))

	args := models.ListAccountsParams{
		Offset: int32((pageNum - 1) * pageSize),
		Limit:  pageSize,
	}
	accounts, err := conn.ListAccounts(ctx, args)
	if err != nil {
		klog.Error("listing accounts failed ", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type depositRequestParams struct {
	DepositToAccountId int64  `json:"deposit_to_account_id"`
	Amount             int64  `json:"amount"`
	Currency           string `json:"currency"`
}

func Deposit(ctx *gin.Context) {
	var request depositRequestParams
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apierror.InvalidRequestBody)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{MESSAGE: "done!"})
}

func ListTransactions(ctx *gin.Context) {
	pageNum := ctx.Param("page_num")
	pageNumInt, _ := strconv.Atoi(pageNum)
	klog.Debug(pageNumInt)

	list, err := store.GetInstance().GetTransactionsList(ctx)
	if err != nil {
		klog.Error("error getting transaction list ", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}
	ctx.JSON(http.StatusOK, list)
}
