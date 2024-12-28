package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func StafRoutes(r *gin.Engine) {
	r.POST("api/staf", controllers.CreateStaf)
	r.GET("api/staf", controllers.GetAllStaf)
	r.GET("api/staf/:id", controllers.GetStaf)
}
