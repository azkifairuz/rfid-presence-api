// routes/fakultasRoutes.go
package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func FakultasRoutes(r *gin.Engine) {
	r.POST("api/fakultas", controllers.CreateFakultas)
	r.GET("api/fakultas", controllers.GetAllFakultas)
	r.GET("api/fakultas/:id", controllers.GetFakultas)
}
