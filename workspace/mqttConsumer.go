package main

import (
	"fmt"

	"github.com/stationedabroad/go-kafka-avro"
	// "github.com/dangkaka/go-kafka-avro"
	"github.com/stationedabroad/sarama-cluster"
)

var esMqttMsgStr = `curl -XPOST -H "Content-Type: application/json" 127.0.0.1:9200/mqtt/_doc/%d -d '
{
        "longitude": %f,
        "latitude": %f,
}'`

// var kafkaServers = []string{"localhost:9092"}
// var schemaRegistryServers = []string{"http://localhost:8081"}
// var topic = "test"
const (
	kafka1 = "kafka-1:39092"
	schemaRegServer = "http://schema-registry:8081"
	ktopic = "mqtt_messages"
	// mqtt://<username>:<password>@farmer.cloudmqtt.com:31352
)

var kafkaServers = []string{kafka1}
var schemaRegistryServers = []string{schemaRegServer}

func main() {
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			fmt.Printf("Message type: %T\n", msg)
			fmt.Println(msg)
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
