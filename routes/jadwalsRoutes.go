package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func JadwalRoutes(r *gin.Engine) {
	r.POST("/jadwal", controllers.CreateJadwal)
	r.GET("/jadwal", controllers.GetAllJadwal)
	r.GET("/jadwal/:id", controllers.GetJadwal)
}
