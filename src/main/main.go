package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	fmt.Printf("Hello \n")
	fmt.Print(a)
	fmt.Printf("\n---------------------\n")

	a[2] = 6
	fmt.Print(a)

	fmt.Print("\n------\n")
	a = append(a, 8)
	fmt.Print(a)
	fmt.Print("\n------\n")

	//----------------

	arr1 := [6]int{10, 11, 12, 13, 14, 15} //Array
	myslice := arr1[1:5]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

}
