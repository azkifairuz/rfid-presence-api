package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func KelasRoutes(r *gin.Engine) {
	r.POST("/kelas", controllers.CreateKelas)
	r.GET("/kelas", controllers.GetAllKelas)
	r.GET("/kelas/:id", controllers.GetKelas)
}
