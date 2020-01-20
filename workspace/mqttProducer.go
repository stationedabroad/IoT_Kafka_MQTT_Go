package main

import ( 
	"fmt"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/stationedabroad/go-kafka-avro"
	"workspace/mqtt"
)

const (
	kafka1 = "kafka-1:9092"
	schemaRegServer = "schema-registry:8082"

	schema = `{
		"type": "record",
		"name": "MqttMessage",
		"fields": [
			{"name": "battery", "type": "int"},
			{"name": "longitude", "type": "float"},
			{"name": "accuracy", "type": "int"},
			{"name": "barometricPressure", "type": "float"},
			{"name": "batteryStatus", "type": "int"},
			{"name": "verticalAccuracy", "type": "int"},
			{"name": "latitude", "type": "float"},
			{"name": "trigger", "type": "string"},
			{"name": "connectivity", "type": "string"},
			{"name": "timestamp", "type": "long"},
			{"name": "altitiude", "type": "float"},
			{"name": "trackerId", "type": "string"}
		]
	}`

	topic = "owntracks/zlaaxmtf/A481FF15-8C60-4118-BE0A-9A0E6554A63C"
	// mqtt://<username>:<password>@farmer.cloudmqtt.com:31352
)

var kafkaServers = []string{kafka1}
var schemaRegistryServers = []string{schemaRegServer}

func main() {
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("could not create producer: %s, error: %v", kafka1, err)
	}
	fmt.Println(fmt.Sprintf("%T\n", *producer))

	// Set the mqtt listener Go'ing
	mqttUri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		log.Fatalf("mqtt uri error: %v\n", err)
	}
	mqttClient := mqtt.NewMqttReceiver("sub", mqttUri)
	topicChannel := make(chan []byte)
	go mqttClient.Listen(topic, topicChannel)

	// Set the Kafka-receiver Go'ing
	var mqttLocation mqtt.MqttLocation
	for {
		recvdMsg := <-topicChannel
		fmt.Println(string(recvdMsg))
		err := json.Unmarshal(recvdMsg, &mqttLocation)
		if err != nil {
			log.Fatalf("Unmarshalling error: %v\n", err)
		}
		SendMessage(producer, mqttLocation)
		// fmt.Printf("Received message:\n%v-%v", mqttLocation.Batt, mqttLocation.Long)
	}
}

func SendMessage(producer *kafka.AvroProducer, loc mqtt.MqttLocation) {
	message := fmt.Sprintf(`{
		"batteryStatus": %d,
		"longitude": %f,
		"accuracy": %d,
		"barometricPressure": %f,
		"bateryStatus": %d,
		"verticalAccuracy": %d,
		"latitude": %f,
		"trigger": %s,
		connectivity": %s,
		"timestamp": %v,
		"altitude": %d,
		"trackerId": %s
	}`, loc.Batt, loc.Long, loc.Acc, loc.P, loc.BS, loc.Vac, loc.Latt, loc.T, loc.Conn, loc.Tst, loc.Alt, loc.Tid)
	key := time.Now().String()
	fmt.Println(message)	
}