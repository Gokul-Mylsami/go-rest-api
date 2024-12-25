package routes

import (
	"gokul-mylsami/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getEvents)
	server.GET("/events/:id",getEvent)

	authenicated := server.Group("/")
	authenicated.Use(middlewares.Authenticate)
	authenicated.POST("/events", createEvent)
	authenicated.PUT("/events/:id",updateEventById)
	authenicated.DELETE("/events/:id",deleteEventById)
	authenicated.POST("/events/:id/register", registerForEvent)
	authenicated.DELETE("/events/:id/register", cancelRegisteration)
	
	server.POST("/signup",signup)
	server.POST("/login",login)
}