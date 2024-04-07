package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacchunn/rest-api/db"
	"github.com/isaacchunn/rest-api/routes"
)

func main() {
	//Set up the gin engine with logger and recovery middleware attached
	//A http server.
	db.InitDB()
	server := gin.Default()
	//Register routes to our server
	routes.RegisterRoutes(server)
	//Run the server on port 8080 for localhost:8080
	server.Run(":8080")
}
