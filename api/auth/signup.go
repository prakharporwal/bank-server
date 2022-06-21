package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	db "github.com/prakharporwal/bank-server/db/sqlc"
	"github.com/prakharporwal/bank-server/services"
	"github.com/prakharporwal/bank-server/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// create a user in DB
// set user to be inactive
// send user activation link email
// delete user if not activated on next 7 days
var klog services.Logger

type signUpRequest struct {
	Username  string `json:"username" binding:"required"`
	UserEmail string `json:"user_email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (controller *AuthController) SignUp(ctx *gin.Context) {
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

	user, err := controller.DB.CreateUser(ctx, args)
	if err != nil {
		klog.Error("db query for user creation failed", err)
		ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func refreshToken() {

}
