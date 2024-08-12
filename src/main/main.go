package main

import "fmt"

func main() {
	// Map:
	// 1. Create
	// 2. Get
	// 3. Set
	// 4. Delete
	// 5. Iterate
	// 6. Length
	// 7. Copy
	// 8. Merge

	mapObj := make(map[string]string)
	mapObj["gr"] = "Grape fruit"
	mapObj["app"] = "Apple fruit"
	mapObj["or"] = "Orange fruit"
	mapObj["ki"] = "Kiwi fruit"

	// Iterate
	for key, value := range mapObj {
		fmt.Println(key, ":", value)
	}

	fmt.Println(mapObj)
	fmt.Println(mapObj["gr"])
	fmt.Println("-------------------------------------------------")
	// switch case

	var value string = "1"
	switch value {
	case "1":
		fmt.Println("One")
		fallthrough
	case "2":
		fmt.Println("Two")
	case "3":
		fmt.Println("Three")

	default:
		fmt.Println("Default")
	}

}
