package main

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Number int `json:"number"`
	Word string
}


func main() {
	// myJsonMap := make(map[string]interface{})
	// jsonData := []byte(`{"hello":"world"}`)
	// err := json.Unmarshal(jsonData, &myJsonMap)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s\n", myJsonMap["hello"])

	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)
}