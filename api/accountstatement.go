package api

import (
	"github.com/prakharporwal/bank-server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/model"
)

var klog services.Logger

const pageSize = 10

type responseBody struct {
	TxnId         int    `json:"txn_id"`
	SenderId      int    `json:"sender_id"`
	RecieverId    int    `json:"reciever_id"`
	RecieverEmail string `json:"reciever_email"`
	Amount        int    `json:"amount"`
}

func (server *Server) GetStatement(ctx *gin.Context) {
	accountId, _ := strconv.Atoi(ctx.Param("account_id"))
	pageNum, _ := strconv.Atoi(ctx.Param("page"))

	transferQuery := `SELECT * FROM account_transactions_entries WHERE account_id=($1) ASC OFFSET ($2) LIMIT ($3)`
	rows := server.store.Query(transferQuery, accountId, (pageNum-1)*pageSize, pageSize)
	defer rows.Close()

	var responseList []responseBody

	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.Id, &transaction.FromAccountID, &transaction.ToAccountID, &transaction.Amount)
		if err != nil {
			klog.Error("failed! scanning transaction : Account Id = ", accountId)
			continue
		}

		recieverEmailQuery :=
			`SELECT T.id, T.from_account_id, T.to_account_id, A.owner_email, T.amount` +
				`FROM transactions as T LEFT JOIN accounts as A` +
				` ON T.from_account_id=A.id WHERE T.from_account_id=($1) or to_account_id=$(1);`
		rows = server.store.Query(recieverEmailQuery, accountId)
		for rows.Next() {
			var response responseBody
			err := rows.Scan(response.TxnId, response.SenderId, response.RecieverId, response.RecieverEmail, response.Amount)
			if err != nil {
				klog.Error("Error in scanning some transactions failed at transaction, id :", transaction.Id)
				ctx.JSON(http.StatusAccepted, responseList)
				return
			}
			responseList = append(responseList, response)
		}
	}

	ctx.JSON(http.StatusOK, responseList)
}
