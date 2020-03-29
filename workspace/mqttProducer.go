package main

import ( 
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/stationedabroad/go-kafka-avro"
	// "github.com/pkg/errors"
	"workspace/mqtt"
)

const (
	kafka1 = "kafka-1:39092"
	schemaRegServer = "http://schema-registry:8081"
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
	// Set the mqtt listener Go'ing
	mqttUri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		log.Fatalf("mqtt uri error: %v\n", err)
	}
	mqttClient := mqtt.NewMqttReceiver("sub", mqttUri)
	topicChannel := make(chan []byte)
	go mqttClient.Listen(topic, topicChannel)

	// Set the Kafka-receiver Go'ing
	schema := getSchema("mqttMessage.avsc")
	var mqttLocation mqtt.MqttLocation
	for {
		recvdMsg := <-topicChannel
		err := json.Unmarshal(recvdMsg, &mqttLocation)
		if err != nil {
			log.Fatalf("Unmarshalling error: %v\n", err)
		}
		SendMessage(schema, producer, mqttLocation)
	}
}

func SendMessage(schema string, producer *kafka.AvroProducer, loc mqtt.MqttLocation) error {
	var mqttTopic = "mqtt_messages"

	// b, err := json.Marshal(loc)
	// if err != nil {
	// 	return errors.Wrap(err, "marshalling json")
	// }
	message := fmt.Sprintf(`{
		"battery": %d,
		"longitude": %f,
		"accuracy": %d,
		"barometricPressure": %f,
		"batteryStatus": %d,
		"verticalAccuracy": %d,
		"latitude": %f,
		"trigger": "%s",
		"connectivity": "%s",
		"timestamp": %v,
		"altitude": %d,
		"trackerId": "%s"
	}`, loc.Batt, loc.Long, loc.Acc, loc.P, loc.BS, loc.Vac, loc.Latt, loc.T, loc.Conn, loc.Tst, loc.Alt, loc.Tid)

	key := time.Now().String()
	err := producer.Add(mqttTopic, schema, []byte(key), []byte(message))
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
	}
	return nil
	// uncomment below to check output in CLI
	// fmt.Printf("Message sent key: %v, msg: %s\n", key, message)	
}

func getSchema(path string) string {
	schema, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(schema)
}