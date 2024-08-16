package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	address string
}

func (p *person) doubleAgeAndPrint() {
	p.age = p.age * 2
	fmt.Println(*p)
}

func main() {
	ali := person{
		name:    "ali",
		age:     22,
		address: "Tehran",
	}
	ali.doubleAgeAndPrint()

	fmt.Println(ali)
	fmt.Println(ali.name)
}
