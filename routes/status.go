package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStatus(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}
