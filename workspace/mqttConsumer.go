package main

import (
	"fmt"
	"encoding/json"
	"log"

	"github.com/stationedabroad/go-kafka-avro"
	// "github.com/dangkaka/go-kafka-avro"
	"github.com/stationedabroad/sarama-cluster"
)

var esMqttMsgStr = `curl -XPOST -H "Content-Type: application/json" 127.0.0.1:9200/mqtt/_doc/%d -d '
{
		"ts": 
        "longitude": %f,
        "latitude": %f,
}'`

// var kafkaServers = []string{"localhost:9092"}
// var schemaRegistryServers = []string{"http://localhost:8081"}
// var topic = "test"
var kafkaServers = []string{kafka1}
var schemaRegistryServers = []string{schemaRegServer}

const (
	kafka1 = "kafka-1:39092"
	schemaRegServer = "http://schema-registry:8081"
	ktopic = "mqtt_messages"
	// mqtt://<username>:<password>@farmer.cloudmqtt.com:31352
)

type mqttRecvdMessage struct {
	Battery int `json:"battery"`
	Longitude float64 `json:"longitude"`
	Accuracy int `json:"accuracy"`
	BarometricPressure float64 `json:"barometricPressure"`
	BatteryStatus int `json:"batteryStatus"`
	VerticalAccuracy int `json:"verticalAccuracy"`
	Latitude float64 `json:"latitude"`
	Trigger string `json:"trigger"`
	Connectivity string `json:"connectivity"`
	Timestamp uint64 `json:"timestamp"`
	Altitude float64 `json:"altitude"`
	TrackerId string `json:"trackerId"`
}

func main() {
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			recvdMsg := &mqttRecvdMessage{}
			if err := json.Unmarshal([]byte(msg.Value), recvdMsg); err != nil {
				log.Fatalf("Unmarshalling error: %v\n", err)
			}
			// fmt.Printf("Message type: %T\n", *recvdMsg)
			// data, _ := json.Marshal(recvdMsg)
			fmt.Printf("unmarshalled/mashalled message: (long: %f, lat: %f)\n", recvdMsg.Longitude, recvdMsg.Latitude)
			// fmt.Println(msg.Value)
			
		},
		OnError: func(err error) {
			fmt.Println("Consumer error", err)
		},
		OnNotification: func(notification *cluster.Notification) {
			fmt.Println(notification)
		},
	}

	consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, ktopic, "consumer-group", consumerCallbacks)
	if err != nil {
		fmt.Println(err)
	}
	consumer.Consume()
}
