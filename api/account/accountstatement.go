package account

import (
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const pageSize = 10

type responseBody struct {
	TransactionId     int       `json:"transaction_id"`
	AccountId         int       `json:"account_id"`
	OtherAccount      int       `json:"other_account"`
	OtherAccountOwner string    `json:"other_account_owner"`
	Amount            int       `json:"amount"`
	Type              string    `json:"type"`
	Timestamp         time.Time `json:"timestamp"`
}

func GetAccountStatement(ctx *gin.Context) {
	accountId, _ := strconv.Atoi(ctx.Param("account_id"))
	pageNum, _ := strconv.Atoi(ctx.Param("page"))

	statementQuery := `SELECT T.transaction_id, T.account_id, T.other_account, A.owner_email, T.amount, T.type, T.created_at ` +
		`FROM account_transactions_entries as T LEFT JOIN accounts as A ` +
		`ON T.other_account=A.id WHERE T.account_id=($1) ` +
		`ORDER BY T.created_at OFFSET ($2) LIMIT ($3);`
	rows := store.GetInstance().Query(statementQuery, accountId, (pageNum-1)*pageSize, pageSize)
	defer rows.Close()

	var responseList []responseBody

	for rows.Next() {
		var response responseBody
		err := rows.Scan(&response.TransactionId, &response.AccountId, &response.AccountId, &response.OtherAccountOwner, &response.Amount, &response.Type, &response.Timestamp)
		if err != nil {
			klog.Error("failed! scanning transaction : Account Id = ", accountId, err)
			continue
		}
		responseList = append(responseList, response)
	}

	ctx.JSON(http.StatusOK, responseList)
}
