package main

import (
	"time"

	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/mqtt"
	"github.com/azkifairuz/rfid-presence-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
  }

func main() {
	go func() {
		mqtt.InitMqtt() // Jalankan fungsi untuk memulai koneksi MQTT dan subscribe topik
	}()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Ganti dengan frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.FakultasRoutes(r)
	routes.ProdiRoutes(r)
	routes.DosenRoutes(r)
	routes.MahasiswaRoutes(r)
	routes.StafRoutes(r)
	routes.KelasRoutes(r)
	routes.JadwalRoutes(r)
	routes.CardRoutes(r)
	routes.PresenceClass(r)
	routes.AuthRoutes(r)

	r.Run()
}
