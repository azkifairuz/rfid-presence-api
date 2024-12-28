package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func DosenRoutes(r *gin.Engine) {
	r.POST("api/dosen", controllers.CreateDosen)
	r.GET("api/dosen", controllers.GetAllDosen)
	r.GET("api/dosen/:id", controllers.GetDosen)
}
