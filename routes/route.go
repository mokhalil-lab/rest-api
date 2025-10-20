package routes

import "github.com/gin-gonic/gin"

func Registerroute(server *gin.Engine){
	server.GET("/events" , getevents)
	server.GET("/events/:id" , getevent)
	server.POST("/events" , createevent)
	server.PUT("/events/:id" , updateevent)
	server.DELETE("/events/:id" , deleteevent)
}