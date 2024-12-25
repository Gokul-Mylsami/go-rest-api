package middlewares

import (
	"gokul-mylsami/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message" : "Not Authorized"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message": "Not Authorized"})
		return
	}

	context.Set("userId",userId)
	context.Next()
}