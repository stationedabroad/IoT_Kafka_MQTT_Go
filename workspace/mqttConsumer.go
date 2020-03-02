package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	// "strings"

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
			// body := strings.NewReader(fmt.Sprintf(`-d '
			// 							{
			// 						        "Timestamp": %d,
			// 						        "Latitude": %f,
			// 						        "Longitude": %f
			// 							}'`, recvdMsg.Timestamp, recvdMsg.Latitude, recvdMsg.Longitude))
			msgByteSlice := []byte(fmt.Sprintf(`{
											        "timestamp": %d,
											        "location": {
											        				"lat": %f, 
											        				"lon": %f
											        			}
												}`, recvdMsg.Timestamp, recvdMsg.Latitude, recvdMsg.Longitude))
			msgBuf := bytes.NewBuffer(msgByteSlice)										

			esUrl := fmt.Sprintf("http://my_es:9200/mqtt/_doc/%d", recvdMsg.Timestamp)
			req, err := http.NewRequest("POST", esUrl, msgBuf)
			if err != nil {
				fmt.Printf("could not create new request to elasticsearch: %v\n", err)
			}
			req.Header.Set("Content-Type", "application/json")
			fmt.Println("Http request: ", req)
			client := &http.Client{}
			resp, err := client.Do(req)
			// resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("could not write to elasticsearch: %v\n", err)
			}
			var b bytes.Buffer
			resp.Write(&b)
			fmt.Printf("response status: %s\nBody: %s\n", resp.Status, b.String())
			defer resp.Body.Close()

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
