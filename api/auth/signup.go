package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/prakharporwal/bank-server/api/apierror"
	"github.com/prakharporwal/bank-server/models/sqlc"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// create a user in DB
// set user to be inactive
// send user activation link email
// delete user if not activated on next 7 days

type signUpRequest struct {
	Username  string `json:"username" binding:"required,alphanum"`
	UserEmail string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
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

	args := db.CreateUserParams{
		UserEmail:    request.UserEmail,
		Username:     request.Username,
		PasswordHash: string(passwordHash),
	}

	user, err := store.GetInstance().CreateUser(ctx, args)
	if err != nil {
		klog.Error("db query for user creation failed", err)
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				ctx.JSON(http.StatusConflict, gin.H{"message": "User with this email or username already exists!"})
			} else {
				ctx.JSON(http.StatusInternalServerError, apierror.UnexpectedError)
			}
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}
