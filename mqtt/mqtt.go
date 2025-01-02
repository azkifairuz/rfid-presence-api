package mqtt

import (
	"encoding/json"
	"log"
	"time"

	// "time"

	// "github.com/azkifairuz/rfid-presence-api/controllers"
	// "github.com/azkifairuz/rfid-presence-api/helper"
	// "github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)
var mqttClient mqtt.Client

func InitMqtt(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")

	mqttClient = mqtt.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v",token.Error())
	}
	log.Printf("connected to mqtt broker")
	
	subscribeToTopic("rfid-system/register", RegisterCardFromMQTT)
	subscribeToTopic("rfid-system/read", ReadCardFromMQTT)
}

func subscribeToTopic(topic string, handler mqtt.MessageHandler)  {
	token := mqttClient.Subscribe(topic,1 ,handler)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v",token.Error())
	}
	log.Printf("connected to mqtt broker")

}

func RegisterCardFromMQTT(client mqtt.Client, msg mqtt.Message) {
	var body struct {
		CardUid string `json:"carduid"`
		Nim     string `json:"nim"`
	}

	if err := helper.ParseJSON(msg.Payload(), &body); err != nil {
		log.Println("Failed to parse register card message:", err)
		return
	}

	var mahasiswa models.Mahasiswa
	if err := initializers.DB.Where("nim = ?", body.Nim).First(&mahasiswa).Error; err != nil {
		log.Println("Mahasiswa not found")
		return
	}

	var existingCard models.MhsCard
	if err := initializers.DB.Where("mahasiswa_id = ?", mahasiswa.ID).First(&existingCard).Error; err == nil {
		log.Println("Card already registered for this mahasiswa")
		return
	}

	card := models.MhsCard{CardUid: body.CardUid, MahasiswaID: mahasiswa.ID}
	if err := initializers.DB.Create(&card).Error; err != nil {
		log.Println("Failed to register card:", err)
		return
	}

	log.Println("Card successfully registered for mahasiswa:", mahasiswa.Name)
}

func ReadCardFromMQTT(client mqtt.Client, msg mqtt.Message) {
	var body struct {
		CardUid string `json:"carduid"`
	}
	if err := helper.ParseJSON(msg.Payload(), &body); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	cardUid := body.CardUid
	var card models.CardResponse
	result := initializers.DB.Table("mhs_cards").
		Select("mhs_cards.card_uid as uid, mahasiswas.nim as mahasiswa_nim, mahasiswas.id as mahasiswa_id, mahasiswas.name as mahasiswa_name, kelas.id AS kelas_id, kelas.name AS mahasiswa_class").
		Joins("JOIN mahasiswas ON mhs_cards.mahasiswa_id = mahasiswas.id").
		Joins("JOIN kelas ON mahasiswas.kelas_id = kelas.id").
		Where("mhs_cards.card_uid = ?", cardUid).
		Scan(&card)

	if result.RowsAffected == 0 {
		client.Publish("rfid-system/response", 0, false, `{"status": "error", "message": "data not found"}`)
		return
	}

	var existingPresensi models.Presensi
	currentDate := time.Now().Format("2006-01-02")
	checkResult := initializers.DB.Where("mahasiswa_id = ? AND DATE(date) = ?", card.MahasiswaId, currentDate).
		First(&existingPresensi)

	if checkResult.RowsAffected > 0 {
		client.Publish("rfid-system/response", 0, false, `{"status": "error", "message": "already checked in for today"}`)
		return
	}

	newPresensi := models.Presensi{
		MahasiswaID: card.MahasiswaId,
		KelasID:     card.KelasId,
		Date:        time.Now(),
	}

	if err := initializers.DB.Create(&newPresensi).Error; err != nil {
		client.Publish("rfid-system/response", 0, false, `{"status": "error", "message": "failed to save presensi"}`)
		return
	}

	response := struct {
		Status  string              `json:"status"`
		Message string              `json:"message"`
		Data    models.CardResponse `json:"data"`
	}{
		Status:  "success",
		Message: "success read card and presence record",
		Data:    card,
	}

	responseJSON, _ := json.Marshal(response)
	client.Publish("rfid-system/response", 0, false, responseJSON)
	log.Println("Presence record created successfully.")
}
