package routes

import (
	controller "eliferden.com/restapi/controller"
	"eliferden.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use( middleware.Authenticate)
	authenticated.POST("/events", controller.CreateEvent)
	authenticated.PUT("/events/:id", controller.UpdateEvent)
	authenticated.DELETE("/events/:id", controller.DeleteEvent)
	authenticated.PUT("/events/:id/register", controller.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", controller.CancelRegistration)

	server.GET("/events", controller.GetEvents)
	server.GET("/events/:id", controller.GetEvent)
	
	server.POST("/signup", controller.SignUp)
	server.POST("/login", controller.Login)
}