package main

import (
	"eliferden.com/restapi/db"
	"eliferden.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //it configures an http server behind the scenes
	routes.RegisterEventRoutes(server)
	server.Run(":8080")
}

