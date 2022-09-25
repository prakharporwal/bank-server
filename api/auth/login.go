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

const TokenAgeInMinutes = 5

type loginRequest struct {
	UserId   string `json:"user_id" binding:"required,email"` // userId can be username or userEmail
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken  string    `json:"access_token"` // userId can be username or userEmail
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func Login(ctx *gin.Context) {
	var request loginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error Parsing! Invalid format", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!"})
		return
	}
	klog.Debug(request.UserId, " is attempting login!")

	// convert provided password and username to hash with salt
	// check for the hash string match or not
	user, err := store.GetInstance().GetUserDetails(ctx, request.UserId)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		klog.Error("Incorrect password", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{apierror.MESSAGE: "Incorrect Password"})
		return
	}

	// respond with a paseto / jwt for further login

	response := generateLoginSession(user.UserEmail, ctx.Request.UserAgent(), ctx.ClientIP())

	klog.Debug(request.UserId, " is logged in successfully!")

	// not working IDK why
	//ctx.SetCookie("access_token", token, TokenAgeInMinutes*60, "/", "localhost:3000", false, false)

	ctx.JSON(http.StatusOK, response)
}

var currentUser string

func SetCurrentUser(username string) {
	currentUser = username
}

func GetCurrentUser() string {
	return currentUser
}

func generateLoginSession(useremail string, useragent string, clientip string) *loginResponse {
	//var auth services.Auth
	tokenMaker, err := NewPasetoMaker()
	if err != nil {
		klog.Error("Paseto Maker Failed ! ", err)
		return nil
	}
	token, err := tokenMaker.CreateToken(useremail, TokenAgeInMinutes*time.Minute)
	if err != nil {
		klog.Error("Token Creation Failed ! ", err)
		return nil
	}

	session, err := CreateSession(useremail, useragent, clientip)

	return &loginResponse{
		AccessToken:  token,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
	}
}
