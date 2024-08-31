package main

import "fmt"

func main() {
	secondLine := "2nd line"
	fmt.Println("Hello world brother")
	fmt.Println(secondLine)

	fmt.Println("\nrecursive output:")
	fmt.Println(recursive(-50))
}

func recursive(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n + recursive(n-1)
}
