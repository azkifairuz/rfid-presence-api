package helper

import (
	"github.com/gin-gonic/gin"
)

// Helper function for sending JSON responses
func ResponseDefault(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
		"statusCode": statusCode,
	})
}
