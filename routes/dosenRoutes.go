package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func DosenRoutes(r *gin.Engine) {
	r.POST("/dosen", controllers.CreateDosen)
	r.GET("/dosen", controllers.GetAllDosen)
	r.GET("/dosen/:id", controllers.GetDosen)
}
