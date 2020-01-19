package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttReceiver Struct {
	client mqtt.Client
	messages int
}

type MqttListener interface {
	Connect(clientId string, uri *url.URL) (mqtt.Client, error)
	createClientOptions(clientId string, uri *url.URL) mqtt.ClientOptions
}

func NewMqttReceiver(instance string) *MqttReceiver {
	return &MqttReceiver{Client:}
}

func (m *mqttListener) Connect(clientId string, uri *url.URL) (mqtt.Client, error) {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
		fmt.Println("waiting ...")
	}
	if err := token.Error; err != nil {
		return nil, fmt.Sprintf("could not connect: %v\n", err)
	}
	return client, nil
}

func (m *mqttListener) createClientOptions(clientId string, uri *url.URL) mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("wss://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts	
}