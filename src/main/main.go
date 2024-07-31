package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error ", error)
		os.Exit(1)
	}

	fmt.Println("Success & code: ", res.StatusCode)
	// fmt.Println("Success & code: ", res)

	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))

}
