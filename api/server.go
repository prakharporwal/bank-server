package api

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/auth"
	"github.com/prakharporwal/bank-server/db"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	router := gin.Default()
	server := &Server{store: store}

	account := AccountController{db: store}

	router.Use(static.Serve("/", static.LocalFile("./ui", true)))

	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, "Hello There")
	})
	router.GET("/account", account.GetAccount)
	router.GET("/account/list/:page", account.ListAccount)
	router.GET("/:account_id/statement/:page", server.GetAccountStatement)
	router.POST("/account", account.CreateAccount)

	transfer := TransactionController{db: store}
	router.POST("/transfer", transfer.Transfer)

	// auth
	auth := auth.AuthController{DB: store}
	router.POST("/login", auth.Login)
	router.POST("/signup", auth.SignUp)

	router.GET("/health", HealthCheck)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Stop() {
	fmt.Print("Stopping Server!")
}
