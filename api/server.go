package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/back-server/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.POST("/account", server.CreateAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func Stop() {

}
