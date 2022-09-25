package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/bank-server/api/apierror"
	"github.com/prakharporwal/bank-server/api/auth"
	"github.com/prakharporwal/bank-server/services/klog"
	"net/http"
)

const symmetricKey = "TjWnZr4u7x!A%D*G-KaPdSgUkXp2s5v8"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// implement auth check from token
		accessToken := ctx.GetHeader("access_token")

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Need Access Token Header"})
			return
		}

		tokenMaker, err := auth.NewPasetoMaker()
		if err != nil {
			klog.Error("Token Creation Failed ! ", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, apierror.UnexpectedError)
			return
		}

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			klog.Error("Token Verification Failed! ", err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, apierror.Forbidden)
			return
		}

		klog.Info("current user ", payload.Username)
		auth.SetCurrentUser(payload.Username)

		ctx.Next()
	}
}
