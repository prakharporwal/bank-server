package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/services"
)

// type Auth interface {
// 	Login()
// 	SignUp()
// 	RefreshToken()
// }

type LoginInput struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginInput LoginInput
	err := ctx.ShouldBindJSON(&loginInput)

	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	// convert provided password and username to hash with salt
	// check for the hash strng match or not
	// respond with a paseto / jwt for further login

	var auth services.Auth

	loginResponse := auth.LoginService(loginInput.UserId, loginInput.Password)

	ctx.JSON(http.StatusOK, loginResponse)
}
