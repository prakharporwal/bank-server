package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountInput struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=INR USD CAD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {

	var account createAccountInput

	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	//TODO: connect to db

	statement := `INSERT INTO accounts (owner,currency,balance) VALUES ($1,$2,$3)`
	server.store.Query(statement, account.Owner, account.Currency, 0) // initial balance will be 0

	log.Println(account.Owner)
	log.Println(account.Currency)

	ctx.JSON(http.StatusOK, gin.H{"message": "Account " + account.Owner})
}
