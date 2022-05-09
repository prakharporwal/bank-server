package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/back-server/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func Stop() {

}
