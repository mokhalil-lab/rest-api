package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func registerevent(context *gin.Context){
	userid := context.GetInt64("userid")
	eventid , err := strconv.ParseInt(context.Param("id") , 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse event id"})
		return
	}
	event , err := models.Geteventbyid(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "event doesn't exist"})
		return
	}

	err = event.Register(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "you can't register"})
		return
	}

	context.JSON(http.StatusCreated , gin.H{"message":"registered succesfully"})
}

func cancelregister(context *gin.Context){
	userid := context.GetInt64("userid")
	eventid , err := strconv.ParseInt(context.Param("id") , 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse event id"})
		return
	}

	var event models.Event
	event.ID = eventid

	err = event.Cancel(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't cancel registration"})
		return
	}

	context.JSON(http.StatusOK , gin.H{"message": "cancelation was succesful"})
}