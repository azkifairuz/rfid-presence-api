// routes/fakultasRoutes.go
package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func FakultasRoutes(r *gin.Engine) {
	r.POST("/fakultas", controllers.CreateFakultas)
	r.GET("/fakultas", controllers.GetAllFakultas)
	r.GET("/fakultas/:id", controllers.GetFakultas)
}
