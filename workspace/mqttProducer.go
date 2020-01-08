package main

import ( 
	"fmt"
	"github.com/stationedabroad/go-kafka-avro"
)

const (
	kafka1 = "kafka-1:9092"
	schemaRegServer = "schema-registry:8081"
)

var kafkaServers = []string{kafka1}
var schemaRegistryServers = []string{schemaRegServer}

func main() {
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("could not create producer: %s, error: %v", kafka1, err)
	}
	fmt.Println(fmt.Sprintf("%T", *producer))
}