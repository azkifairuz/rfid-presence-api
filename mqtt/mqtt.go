package mqtt

import (
	"log"
	"time"

	"github.com/azkifairuz/rfid-presence-api/controllers"
	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)
var mqttClient mqtt.Client

func initMqtt(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")

	mqttClient = mqtt.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v",token.Error())
	}
	log.Printf("connected to mqtt broker")
	
}