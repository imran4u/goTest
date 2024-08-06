package main

import (
	handler "gotest/src/main/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HandleHomePage)
	http.HandleFunc("/person", handler.HandlePerson)
	log.Println("Starting server ....")
	http.ListenAndServe(":8080", nil)
}
