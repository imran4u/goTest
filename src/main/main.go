package main

import "fmt"

func main() {
	defer funOne()
	defer funTwo()

	defer func(end string) {
		fmt.Println(end)
	}("The end")

	secondLine := "2nd line"
	fmt.Println("Hello world brother")
	fmt.Println(secondLine)

	// Hello world brother
	// 2nd line
	// The end
	// funTwo
	// funOne

}

func funOne() {
	fmt.Println("funOne")
}

func funTwo() {
	fmt.Println("funTwo")
}
