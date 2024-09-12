package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	file, err := os.ReadFile("yaml/hello.yaml")
	if err != nil {
		log.Fatal("file not found", err)
	}

	// Create an empty Car to be are target of unmarshalling
	var car Car

	// Unmarshal our input YAML file into empty Car (var c)
	if err := yaml.Unmarshal(file, &car); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", car)
	fmt.Println("\n ------------json formate ---------------")
	// convert to json
	jsonBytes, err := json.Marshal(car)
	if err != nil {
		log.Fatal("Error converting to json", err)
	}
	fmt.Println(string(jsonBytes))
}
