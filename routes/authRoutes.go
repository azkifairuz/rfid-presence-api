package routes

import (
	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("auth/login", controllers.Login)
	r.POST("auth/change_pw", controllers.ChangePassword)
}