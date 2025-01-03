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
		&models.Kelas{},
		&models.Mahasiswa{},
		&models.Staf{},
		&models.Jadwal{},
		&models.Account{},
		&models.MhsCard{},
		&models.Presensi{},
	);err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
		fmt.Printf("Error during AutoMigrate: %v", err)
	}
	fmt.Println("migrate succes")

}