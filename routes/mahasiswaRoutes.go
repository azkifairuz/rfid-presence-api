package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes(r *gin.Engine) {
	r.POST("/mahasiswa", controllers.CreateMahasiswa)
	r.GET("/mahasiswa", controllers.GetAllMahasiswa)
	r.GET("/mahasiswa/:id", controllers.GetMahasiswa)
}
