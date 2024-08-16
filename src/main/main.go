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

	// assigns value to struct members
	rect := Rectangle{7, 4, "rectangle"}
	circ := Circle{5, "circle"}
	// call calculate() with struct variable rect
	calculate(rect)
	calculate(circ)

}

//2nd Interface

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
	name            string
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

type Circle struct {
	radius float32
	name   string
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println(s)
	fmt.Println("Area:", s.area())
}
