package main

import (
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"log"
	"fmt"
)


func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	if err := initializers.DB.AutoMigrate(
		&models.Fakultas{},
		&models.Prodi{},
		&models.Dosen{},
		&models.Mahasiswa{},
		&models.Staf{},
		&models.Kelas{},
		&models.Jadwal{},
		&models.Account{},
	);err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
		fmt.Printf("Error during AutoMigrate: %v", err)
	}
	fmt.Println("migrate succes")

}