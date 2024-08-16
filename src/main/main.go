// Golang program illustrates how
// to implement an interface
package main

import "fmt"

// Creating an interface
type vehicle interface {

	// Methods
	haveFourWheel() bool
	highSpeed() bool // if speed more than 100
}

type car struct {
	fuelType string // petrol or desial
}
type bike struct {
	color string
}

func (c car) haveFourWheel() bool {
	return true
}
func (c car) highSpeed() bool {
	return true
}

func (b bike) haveFourWheel() bool {
	return false
}
func (b bike) highSpeed() bool {
	return false
}

// Main Method
func main() {

	var v vehicle

	v = car{fuelType: "petrol"}

	fmt.Println("have four wheel :", v.haveFourWheel())
	fmt.Println("high speed:", v.highSpeed())
	fmt.Println(v)

	v = bike{color: "red"}

	fmt.Println("have four wheel :", v.haveFourWheel())
	fmt.Println("high speed:", v.highSpeed())
	fmt.Println(v)

}
