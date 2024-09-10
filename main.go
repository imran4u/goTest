package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	car := Car{
		TopSpeed: 10,
		Name:     "Chevy",
		Cool:     true,
		Passengers: []string{
			"imran",
			"ali",
			"farooq",
			"waleed",
		},
	}

	out, err := yaml.Marshal(car) // convert struct to yaml
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out)) // print yaml (out)
}
