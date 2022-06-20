package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/auth"
	"github.com/prakharporwal/bank-server/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.GET("/account", server.GetAccount)
	router.GET("/account/list/:page", server.ListAccount)
	router.GET("/:account_id/statement/:page", server.GetAccountStatement)
	router.POST("/account", server.CreateAccount)

	router.POST("/login", auth.Login)
	router.POST("/transfer", server.TransferMoney)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func Stop() {
	fmt.Print("Stopping Server!")
}
