package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.Engine) {
	r.POST("card/register", controllers.RegisterCard)
	r.GET("card/read/:uid", controllers.ReadCard)
}
