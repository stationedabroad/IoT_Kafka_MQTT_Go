package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
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
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s]:[%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	// Make sure env variable MQTT_URL is set or use .env file
	mqtt_uri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := mqtt_uri.Path[1:len(mqtt_uri.Path)] + "/#"

	go listen(mqtt_uri, topic)
	
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
	}
}