package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//Setup handler for incoming get request
	//GET POST PUT PATCH DELETE
	server.GET("/status", getStatus)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //events/1 /events/5

	server.POST("/events", createEvent)

}
