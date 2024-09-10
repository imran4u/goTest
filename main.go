package main

import (
	"fmt"
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
	fmt.Println(car)
}
