package main

import "fmt"

func main() {
	// 5, 8 alishan : wrong

	// var a int = 5
	a := 5
	fmt.Printf("value of a %d\n", a)
	changeValue(a)
	fmt.Printf("after changeValue: value of a %d\n", a)

	changeByRef(&a)
	fmt.Printf("After changeBy ref: value of a %d\n", a) //9
	changeByRef(&a)
	fmt.Printf("After changeBy ref: value of a %d\n", a) //13
}

func changeValue(a int) {
	a = a + 3
	fmt.Printf("Inside changeValue: value of a %d\n", a)
}

func changeByRef(a *int) {
	*a = *a + 4
}
