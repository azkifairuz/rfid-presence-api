package initializers
import (
	"log"
    "github.com/joho/godotenv"
)
func LoadEnvVariables()  {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	  return
	}
	log.Fatal("connected to db")

}