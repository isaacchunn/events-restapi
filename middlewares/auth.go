package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacchunn/rest-api/utils"
)

func Authenticate(context *gin.Context) {
	//Have to attach valid tokens to go on protected routes
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		//No token part of request
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	//Handle for token
	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	//Set some data on the context
	context.Set("userId", userID)
	//Handle next request in line correctly
	context.Next()
}
