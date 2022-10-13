package api

import (
	"log"
	"net"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/account"
	"github.com/prakharporwal/bank-server/api/auth"
	"github.com/prakharporwal/bank-server/api/middleware"
	//"github.com/prakharporwal/bank-server/auth"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		// Get preferred outbound ip of this machine
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		localAddr := conn.LocalAddr().(*net.UDPAddr)

		context.JSON(200, gin.H{"msg": "Hello There", "ip": localAddr.IP})
	})

	public := router.Group("/public")
	public.Use(cors.Default()) // as this is public we don't need access_token header

	public.GET("/health", HealthCheck)

	public.POST("/v1/login", auth.Login)
	public.POST("/v1/signup", auth.SignUp)

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthMiddleware())

	router.GET("/account", account.GetAccount)
	public.GET("/v1/transaction/list/:page_num", account.ListTransactions)
	router.GET("/account/list/:page", account.ListAccount)

	router.GET("/v1/account/:account_id/statement/:page", account.GetAccountStatement)
	router.POST("/account", account.CreateAccount)
	router.POST("/v1/session/refresh", auth.RefreshSession)
	//router.GET("/:account_id/statement/:page", account.Get)
	//router.POST("/login", auth.Login)
	router.POST("/deposit", account.Deposit)
	router.POST("/transfer", account.TransferTx)

}
