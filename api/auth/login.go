package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	db2 "github.com/prakharporwal/bank-server/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthController struct {
	DB *db2.Store
}

type loginRequest struct {
	UserId   string `json:"user_id" binding:"required"` // userId can be username or userEmail
	Password string `json:"password" binding:"required"`
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var request loginRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	// convert provided password and username to hash with salt
	// check for the hash strng match or not
	user, err := controller.DB.GetUserDetails(ctx, request.UserId)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		klog.Error("Incorrect password", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{apierror.MESSAGE: "Incorrect Password"})
		return
	}

	// respond with a paseto / jwt for further login
	//var auth services.Auth
	token := "token"

	ctx.JSON(http.StatusOK, token)
}
