package main

import ( 
	"fmt"
	"github.com/stationedabroad/go-kafka-avro"
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
)

var kafkaServers = []string{kafka1}
var schemaRegistryServers = []string{schemaRegServer}

func main() {
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("could not create producer: %s, error: %v", kafka1, err)
	}
	fmt.Println(fmt.Sprintf("%T", *producer))
	fmt.Println(schema)
}

func SendMessage(producer *kafka.AvroProducer, schema string) {
	message := `{
		"batteryStatus": "",
		"longitude": "",
		"accuracy": "",
		"barometricPressure": "",
		"bateryStatus": "",
		"verticalAccuracy": "",
		"latitude": "",
		"trigger": "",
		connectivity": "",
		"timestamp": "",
		"altitude": "",
		"trackerId": ""
	}`
	fmt.Println(message)	
}