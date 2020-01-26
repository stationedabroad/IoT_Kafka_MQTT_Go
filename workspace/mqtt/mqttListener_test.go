package mqtt

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	// "mqtt"
)

func TestReceiverCreation(t *testing.T) {
	// Remember to set MQTT env variable with mqtt server details e.g. mqtt://<username>:<password>@farmer.cloudmqtt.com:port
	uri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		t.Errorf("parse uri error: %v", err)
	}
	mqttClient := NewMqttReceiver("sub", uri)
	mqttType := fmt.Sprintf("%T", mqttClient)
	if mqttType != "*mqtt.MqttReceiver" {
		t.Errorf("incorrect mqtt receiver type, expected: %v, received: %v", mqttType, uri)
	}

	if mqttClient.recvdMsgCount != 0 {
		t.Errorf("received message count != 0 at receiver creation, count is: %d", mqttClient.recvdMsgCount)
	}
}

func TestListener(t *testing.T) {
	// This is difficult to test as requires sending messages from an MQTT IoT server: creae mqtt producer ... ?
	t.Skip()
}