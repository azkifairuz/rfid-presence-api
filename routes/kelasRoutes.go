package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func KelasRoutes(r *gin.Engine) {
	r.POST("api/kelas", controllers.CreateKelas)
	r.GET("api/kelas", controllers.GetAllKelas)
	r.GET("api/kelas/:id", controllers.GetKelas)
}
