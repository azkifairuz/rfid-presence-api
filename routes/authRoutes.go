package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("api/auth/login", controllers.Login)
	r.POST("api/auth/change_pw", controllers.ChangePassword)
}