package apierror

import "github.com/gin-gonic/gin"

const MESSAGE = "message"

var (
	UnexpectedError    = gin.H{MESSAGE: "Something went wrong!"}
	InvalidRequestBody = gin.H{MESSAGE: "Invalida Request Body. Check the request JSON again!"}
	NotFound           = gin.H{MESSAGE: "Record Not Found!"}
	Forbidden          = gin.H{MESSAGE: "You can't access this resource"}
)
