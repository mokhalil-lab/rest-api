package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"couldn't parse user data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't save user to db"})
		return
	}
	context.JSON(http.StatusOK , gin.H{"message": "user saved succesfully"})
}

func login(context *gin.Context){
	var user models.User 
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse data"})
		return
	}
	err = user.Authenticate()
	if err != nil {
		context.JSON(http.StatusUnauthorized , gin.H{"message": "you can't login"})
		return
	}

	token , err := utils.Generatetoken(user.Email,user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"couldn't generate a jwt token"})
		return
	}

	context.JSON(http.StatusOK , gin.H{"message": "login successful", "token": token})
}