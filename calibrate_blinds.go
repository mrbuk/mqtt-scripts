package main

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var blinds = []string{
	"sashka/blind-garage",
	"kitchen/blind-door",
	"kitchen/blind-small",
	"kitchen/blind-large",
	"kitchen/blind-street",
	"livingroom/blind-left",
}

func main() {

	broker := os.Getenv("MQTT_BROKER")

	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	for _, s := range blinds {
		topic := fmt.Sprintf("zwave/%s/112/0/29/set", s)
		log.Printf("calibrate for: %s\n", s)
		token := client.Publish(topic, 0, false, "1")
		if token.Wait() && token.Error() != nil {
			log.Printf("error with blind %s: %s\n", s, token.Error())
		}
	}

	client.Disconnect(3000)
}
