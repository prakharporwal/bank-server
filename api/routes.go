package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/account"
	"github.com/prakharporwal/bank-server/api/auth"
	"github.com/prakharporwal/bank-server/api/middleware"
	//"github.com/prakharporwal/bank-server/auth"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, "Hello There")
	})

	router.GET("/health", HealthCheck)
	router.Use(middleware.CORSMiddleware())

	protected := router.Use(middleware.AuthMiddleware())

	router.GET("/account", account.GetAccount)
	router.GET("/v1/transaction/list/:page_num", account.ListTransactions)
	router.GET("/account/list/:page", account.ListAccount)
	router.POST("/v1/login", auth.Login)

	protected.GET("/:account_id/statement/:page", account.GetAccountStatement)
	protected.POST("/account", account.CreateAccount)
	//router.GET("/:account_id/statement/:page", account.Get)
	//router.POST("/login", auth.Login)

	protected.POST("/deposit", account.Deposit)
	protected.POST("/transfer", account.TransferTx)

}
