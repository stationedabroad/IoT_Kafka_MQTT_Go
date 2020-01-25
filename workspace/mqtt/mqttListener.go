package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqttp "github.com/eclipse/paho.mqtt.golang"
)

type MqttLocation struct {
	Batt int 	 `json:"batt"`
	Long float64 `json:"lon"`
	Acc  int 	 `json:"acc"`
	P 	 float64 `json:"p"`
	BS 	 int 	 `json:"bs"`
	Vac  int 	 `json:"vac"`
	Latt float64 `json:"lat"`
	T  	 string  `json:"t"`
	Conn string  `json:"conn"`
	Tst  uint64  `json:"tst"`
	Alt  int 	 `json:"alt"`
	Type string  `json:"_type"`
	Tid  string  `json:"tid"`
}

type MqttReceiver struct {
	client mqttp.Client
	recvdMsgCount int
}

func NewMqttReceiver(clientId string, uri *url.URL) *MqttReceiver {
	fmt.Println("entered NewMqttRec ...", clientId, uri)
	mqttClient := connect(clientId, uri)
	if mqttClient == nil {
		log.Fatal("error in New mqtt receiever	")
	}
	log.Printf("client connected now listening on: [%s]\n", uri)
	return &MqttReceiver{client: mqttClient, recvdMsgCount: 0}
}

func (m *MqttReceiver) Listen(topic string, recvch chan<- []byte) {
	m.client.Subscribe(topic, 0, func(client mqttp.Client, msg mqttp.Message) {
		recvch <- msg.Payload()
		m.recvdMsgCount++
		fmt.Printf("message count %d\n", m.recvdMsgCount)
	})
}

func connect(clientId string, uri *url.URL) mqttp.Client {
	opts := createClientOptions(clientId, uri)
	// fmt.Println(opts)
	client := mqttp.NewClient(opts)
	token := client.Connect()
	// fmt.Println("CLIENT: ", client)
	for !token.WaitTimeout(3 * time.Second) {
		fmt.Println("waiting ...")
	}
	if err := token.Error; err != nil {
		// fmt.Println("Error in Connect ...")
		// return nil
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqttp.ClientOptions {
	opts := mqttp.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("wss://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts	
}