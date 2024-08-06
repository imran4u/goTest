package main

import (
	handler "gotest/src/main/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HandleHomePage)
	log.Println("Starting server ....")
	http.ListenAndServe(":8080", nil)
}
