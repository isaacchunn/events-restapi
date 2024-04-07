package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//Setup handler for incoming get request
	//GET POST PUT PATCH DELETE

	//GET REQUESTS
	server.GET("/status", getStatus)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //events/1 /events/5

	//POST REQUESTS
	server.POST("/events", createEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	//PUT REQUESTS for updating
	server.PUT("/events/:id", updateEvent)

	//DELETE requests for deleting
	server.DELETE("/events/:id", deleteEvent)
}
