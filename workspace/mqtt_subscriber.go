package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/url"
	"time"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	Topic = "owntracks/zlaaxmtf/A481FF15-8C60-4118-BE0A-9A0E6554A63C"
	// mqtt://zlaaxmtf:_rTbTI7V_Sxm@farmer.cloudmqtt.com:31352
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	fmt.Println(opts)
	client := mqtt.NewClient(opts)
	// fmt.Println("CLIENT: ", client)
	token := client.Connect()
	// fmt.Println("TOKEN: ", token)
	for !token.WaitTimeout(30 * time.Second) {
		fmt.Println("waiting ...")
	}
	if err := token.Error; err != nil {
		// fmt.Println("client connected 5 Error")
		// log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("wss://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string) {
	client := connect("sub", uri)
	fmt.Printf("client connected now listening on: [%s]\n", uri)
	fmt.Println("Topic is : ",topic)
	var f interface{}
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s]:[%s]\n\n", msg.Topic(), string(msg.Payload()))
		err := json.Unmarshal(msg.Payload(), &f)
		if err != nil {
			//
		}
		fmt.Println(f)
	})
}

func main() {
	// Make sure env variable MQTT_URL is set or use .env file
	mqtt_uri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		fmt.Println("error entered ...", mqtt_uri)
		log.Fatal(err)
	}1

	go listen(mqtt_uri, Topic)
	
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
	}
}