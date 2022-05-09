package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createAccount(ctx *gin.Context) {
	user_id := ctx.Param("id")

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Account " + user_id})
}
