package auth

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	models "github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
	"net/http"
	"time"
)

const SessionTokenAgeInMinutes = 60

func CreateSession(email string, userAgent string, clientIP string) (*models.Session, error) {
	tokenMaker, err := NewPasetoMaker()
	if err != nil {
		klog.Error("Paseto Maker Failed ! ", err)
		return nil, errors.New("paseto Maker Failed ")
	}

	token, err := tokenMaker.CreateToken(email, TokenAgeInMinutes*time.Minute)
	if err != nil {
		klog.Error("Token Creation Failed ! ", err)
		return nil, errors.New("paseto Maker Failed ")
	}

	args := models.CreateSessionParams{
		Email:        email,
		ClientIp:     clientIP,
		UserAgent:    userAgent,
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(SessionTokenAgeInMinutes * time.Minute),
	}
	// create session in DB
	session, _ := store.GetInstance().CreateSession(context.Background(), args)

	return &session, nil
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshSession(ctx *gin.Context) {
	var request refreshRequest
	err := ctx.ShouldBindJSON(&request)
	useremail := GetCurrentUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apierror.InvalidRequestBody)
		return
	}

	// generate a new pair of tokens and return
	response := generateLoginSession(useremail, ctx.Request.UserAgent(), ctx.ClientIP())

	ctx.JSON(http.StatusOK, response)
}
