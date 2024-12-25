// routes/fakultasRoutes.go
package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func ProdiRoutes(r *gin.Engine) {
	r.POST("/prodi", controllers.CreateProdi)
	r.GET("/prodi", controllers.GetAllProdi)
	r.GET("/prodi/:id", controllers.GetProdi)
}
