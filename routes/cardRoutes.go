package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.POST("api/card/register", controllers.RegisterCard)
	r.GET("api/card/read/:uid", controllers.ReadCard)
}
