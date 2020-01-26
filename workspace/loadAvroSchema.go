package main

import (
	"fmt"

	"github.com/elodina/go-avro"
)

func main() {
	schema := avro.LoadSchemas("./")
	fmt.Println(schema)
	fmt.Printf("%T\n", schema)
	fmt.Printf("Schema name: %s\n", schema["sulman.go.com.mqttmessage"].GetName())
	fmt.Printf("JSON representation: %s\n", schema["sulman.go.com.mqttmessage"].String())
	fmt.Printf("Schema Type: %d\n", schema["sulman.go.com.mqttmessage"].Type())

}