package main

type Car struct {
	TopSpeed   int      `yaml:"topspeed", json:"topspeed` // top speed in miles per hour
	Name       string   `yaml:"name", json:"name"`
	Cool       bool     `yaml:"cool", json:"cool"`
	Passengers []string `yaml:"passengers", json:"passengers"`
}
