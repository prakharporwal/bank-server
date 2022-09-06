package api

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/prakharporwal/bank-server/db"
	//"github.com/prakharporwal/bank-server/auth"
)

func InitRouter(router *gin.Engine) {
	account := AccountController{db: db2.GetInstance()}
	router.GET("/account", account.GetAccount)
	router.GET("/account/list/:page", account.ListAccount)
	router.POST("/account", account.CreateAccount)
	transfer := TransactionController{db: db2.GetInstance()}
	//router.GET("/:account_id/statement/:page", account.Get)
	//router.POST("/login", auth.Login)
	router.POST("/transfer", transfer.Transfer)
}
