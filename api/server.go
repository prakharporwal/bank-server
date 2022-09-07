package api

import (
	"fmt"
	"github.com/prakharporwal/bank-server/models/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *store.SQLStore
	router *gin.Engine
}

func NewServer(store *store.SQLStore) *Server {
	router := gin.Default()
	server := &Server{store: store}

	InitRoutes(router)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Stop() {
	fmt.Print("Stopping Server!")
}
