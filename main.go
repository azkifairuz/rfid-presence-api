package main

import (
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
  }

func main() {
	r := gin.Default()
	routes.FakultasRoutes(r)
	routes.ProdiRoutes(r)
	routes.DosenRoutes(r)
	routes.MahasiswaRoutes(r)
	routes.StafRoutes(r)
	routes.KelasRoutes(r)

	r.Run()
}
