package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func PresenceClass(r *gin.Engine) {
	r.GET("presence/:class", controllers.GetPresenceByClass)
}
