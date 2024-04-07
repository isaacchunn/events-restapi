package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacchunn/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//Setup handler for incoming get request
	//GET POST PUT PATCH DELETE

	//Group Routes that need authentication
	authenticated := server.Group("/")
	//Use authentication middleware before handling sensitive operations
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	//GET REQUESTS
	server.GET("/status", getStatus)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //events/1 /events/5

	//POST REQUESTS
	server.POST("/signup", signup)
	server.POST("/login", login)
}
