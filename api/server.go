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
	router.POST("/account", server.CreateAccount)
	router.POST("/login", auth.Login)

	router.POST("/transaction", server.UpdateBalance)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func Stop() {
	fmt.Print("Stoping Server!")
}
