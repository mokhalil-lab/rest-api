package middleware

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized , gin.H{"message": "invalid access"})
		return
	}

	userid , err := utils.Verifytoken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized , gin.H{"message": "invalid token"})
		return
	}

	context.Set("userid" , userid)
	context.Next()

}