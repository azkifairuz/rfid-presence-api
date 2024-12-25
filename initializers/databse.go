package initializers
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
  )
  var DB *gorm.DB
func ConnectToDb()  {
	var err error
	dsn := "host=localhost user=azki password=azkidb dbname=rfid_presence_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})	

	if err != nil {
		log.Fatal("failed connect to db")

	}
}