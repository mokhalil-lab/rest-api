package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getevent(context *gin.Context){
	eventid , err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse event id"})
		return
	}
	event , err := models.Geteventbyid(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't fetch event"})
		return
	}
	context.JSON(http.StatusOK , event)
}

func getevents(context *gin.Context){
	events , err := models.Getallevents()
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"error": "couldn't get events"})
		return
	}

	context.JSON(http.StatusOK , events)
}

func createevent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't process your request"})
		return
	}
	userid := context.GetInt64("userid")
	event.Userid = int(userid)
	
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"error": "couldn't save event"})
		return
	}
	context.JSON(http.StatusCreated , gin.H{"message" : "event created" , "event": event})


}

func updateevent(context *gin.Context){
	eventid , err := strconv.ParseInt(context.Param("id") , 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse event id"})
		return
	}

	userid := context.GetInt64("userid")
	event , err := models.Geteventbyid(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't get the event from the database"})
		return 
	}

	if int64(event.Userid) != userid {
		context.JSON(http.StatusUnauthorized , gin.H{"message": "you can't do that"})
		return
	}

	var updatedevent models.Event
	err = context.ShouldBindJSON(&updatedevent)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse request data"})
		return
	}

	updatedevent.ID = eventid
	err = updatedevent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't update the event"})
	}
	context.JSON(http.StatusOK , gin.H{"message": "updated succesfully"})
}

func deleteevent(context *gin.Context){
	eventid , err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message": "couldn't parse event id"})
		return
	}
	userid := context.GetInt64("userid")
	event , err := models.Geteventbyid(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message": "couldn't fetch event from database"})
		return
	}

	if int64(event.Userid) != userid {
		context.JSON(http.StatusUnauthorized , gin.H{"message": "you can't do that"})
		return
	}

	event.Delete()
	context.JSON(http.StatusOK , gin.H{"message": "event deleted succesfully"})
}
