package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func Registerroute(server *gin.Engine){

	protected := server.Group("/")
	protected.Use(middleware.Authenticate)
	protected.POST("/events" , createevent)
	protected.PUT("/events/:id" , updateevent)
	protected.DELETE("/events/:id" , deleteevent)
	protected.POST("/events/:id/register", registerevent)
	protected.DELETE("/events/:id/register", cancelregister)
	server.GET("/events" , getevents)
	server.GET("/events/:id" , getevent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}