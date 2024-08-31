package main

import (
	"fmt"
	"os"
)

func init() {
	os.Setenv("NAME", "Imran")
	os.Setenv("NAME_TITLE", "ali")
}

func main() {
	fmt.Println("Hello, World!")

	name := os.Getenv("NAME")
	title := os.Getenv("NAME_TITLE")

	fmt.Printf("Hello, %s %s!\n", name, title)
}
