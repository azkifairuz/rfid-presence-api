package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func JadwalRoutes(r *gin.Engine) {
	r.POST("api/jadwal", controllers.CreateJadwal)
	r.GET("api/jadwal", controllers.GetAllJadwal)
	r.GET("api/jadwal/:id", controllers.GetJadwal)
}
