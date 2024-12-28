// routes/fakultasRoutes.go
package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func ProdiRoutes(r *gin.Engine) {
	r.POST("api/prodi", controllers.CreateProdi)
	r.GET("api/prodi", controllers.GetAllProdi)
	r.GET("api/prodi/:id", controllers.GetProdi)
}
