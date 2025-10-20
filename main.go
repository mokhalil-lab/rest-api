package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	db.Initdb()
	server := gin.Default()
	routes.Registerroute(server)
	server.Run(":8080")
}

