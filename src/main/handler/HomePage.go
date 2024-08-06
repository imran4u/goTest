package hander

import (
	"fmt"
	"net/http"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	result := "Hello World"
	fmt.Fprint(w, result)
}
