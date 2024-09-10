package main

type Car struct {
	TopSpeed   int      `yaml:"topspeed"` // top speed in miles per hour
	Name       string   `yaml:"name"`
	Cool       bool     `yaml:"cool"`
	Passengers []string `yaml:"passengers"`
}
