package routes

import (
	controller "eliferden.com/restapi/controller"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controller.GetEvents)
	server.GET("/events/:id", controller.GetEvent)
	server.POST("/events", controller.CreateEvent)
	server.PUT("/events/:id", controller.UpdateEvent)
	server.DELETE("/events/:id", controller.DeleteEvent)
	server.POST("/signup", controller.SignUp)
}