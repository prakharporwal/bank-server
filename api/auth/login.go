package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

const symmetricKey = "wicqr8u362uduijq"

type loginRequest struct {
	UserId   string `json:"user_id" binding:"required"` // userId can be username or userEmail
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var request loginRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}

	// convert provided password and username to hash with salt
	// check for the hash strng match or not
	user, err := store.GetInstance().GetUserDetails(ctx, request.UserId)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		klog.Error("Incorrect password", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{apierror.MESSAGE: "Incorrect Password"})
		return
	}

	// respond with a paseto / jwt for further login
	//var auth services.Auth
	tokenMaker, _ := NewPasetoMaker(symmetricKey)
	token, _ := tokenMaker.CreateToken(user.UserEmail, 5*time.Minute)

	ctx.JSON(http.StatusOK, token)
}
