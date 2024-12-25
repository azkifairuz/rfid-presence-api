package main

import (
	"github.com/azkifairuz/rfid-presence-api/controlers"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
  }

func main() {
	r := gin.Default()
	r.POST("fakultas/", controlers.CreateFakultas)
	r.GET("fakultas/", controlers.GetAllFakultas)
	r.GET("fakultas/:id", controlers.GetFakultas)

	r.Run() // listen and serve on 0.0.0.0:8080
}
