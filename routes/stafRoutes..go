package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func StafRoutes(r *gin.Engine) {
	r.POST("/staf", controllers.CreateStaf)
	r.GET("/staf", controllers.GetAllStaf)
	r.GET("/staf/:id", controllers.GetStaf)
}
