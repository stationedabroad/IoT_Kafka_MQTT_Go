package main

import (
	"fmt"
	"io/ioutil"

	// "github.com/elodina/go-avro"
)

func main() {
	schema, err := ioutil.ReadFile("mqttMessage.avsc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(schema))
	// schema := avro.LoadSchemas("./")
	// fmt.Println(schema)
	// fmt.Printf("%T\n", schema)
	// fmt.Printf("Schema name: %s\n", schema["sulman.go.com.mqttmessage"].GetName())
	// fmt.Printf("JSON representation: %s\n", schema["sulman.go.com.mqttmessage"].String())
	// fmt.Printf("Schema Type: %d\n", schema["sulman.go.com.mqttmessage"].Type())

}


	// schema = `{
	// 	"namespace": "sulman.go.com",
	// 	"type": "record",
	// 	"name": "mqttmessage",
	// 	"fields": [
	// 		{"name": "battery", "type": "int"},
	// 		{"name": "longitude", "type": "float"},
	// 		{"name": "accuracy", "type": "int"},
	// 		{"name": "barometricPressure", "type": "float"},
	// 		{"name": "batteryStatus", "type": "int"},
	// 		{"name": "verticalAccuracy", "type": "int"},
	// 		{"name": "latitude", "type": "float"},
	// 		{"name": "trigger", "type": "string"},
	// 		{"name": "connectivity", "type": "string"},
	// 		{"name": "timestamp", "type": "long"},
	// 		{"name": "altitude", "type": "float"},
	// 		{"name": "trackerId", "type": "string"}
	// 	]
	// }`