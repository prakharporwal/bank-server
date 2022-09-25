package api

import (
	"github.com/gin-contrib/cors"
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

	public := router.Group("/public")
	public.Use(cors.Default()) // as this is public we don't need access_token header

	public.GET("/health", HealthCheck)

	public.POST("/v1/login", auth.Login)
	public.POST("/v1/signup", auth.SignUp)

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthMiddleware())

	router.GET("/account", account.GetAccount)
	router.GET("/v1/transaction/list/:page_num", account.ListTransactions)
	router.GET("/account/list/:page", account.ListAccount)

	router.GET("/v1/account/:account_id/statement/:page", account.GetAccountStatement)
	router.POST("/account", account.CreateAccount)
	router.POST("/v1/session/refresh", auth.RefreshSession)
	//router.GET("/:account_id/statement/:page", account.Get)
	//router.POST("/login", auth.Login)
	router.POST("/deposit", account.Deposit)
	router.POST("/transfer", account.TransferTx)

}
