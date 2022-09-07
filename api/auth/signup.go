package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	"github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
	"github.com/prakharporwal/bank-server/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// create a user in DB
// set user to be inactive
// send user activation link email
// delete user if not activated on next 7 days

type signUpRequest struct {
	Username  string `json:"username" binding:"required"`
	UserEmail string `json:"user_email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func SignUp(ctx *gin.Context) {
	var request signUpRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		klog.Error("parsing json data failed!")
		ctx.JSON(http.StatusInternalServerError, apierror.InvalidRequestBody)
		return
	}

	// TODO: validate username format not same as email
	// TODO: add email validation
	// TODO: validate password strength

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	if err != nil {
		klog.Error("Failed encrypting password", err)
		return
	}

	userId := utils.GenerateTimeStampMilli()
	args := db.CreateUserParams{
		UserID:       userId,
		UserEmail:    request.UserEmail,
		Username:     request.Username,
		PasswordHash: string(passwordHash),
	}

	user, err := store.GetInstance().CreateUser(ctx, args)
	if err != nil {
		klog.Error("db query for user creation failed", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func refreshToken() {

}
