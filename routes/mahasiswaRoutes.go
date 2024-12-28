package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes(r *gin.Engine) {
	r.POST("api/mahasiswa", controllers.CreateMahasiswa)
	r.GET("api/mahasiswa", controllers.GetAllMahasiswa)
	r.GET("api/mahasiswa/:id", controllers.GetMahasiswa)
}
