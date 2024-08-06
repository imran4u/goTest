package hander

import (
	"encoding/json"
	"fmt"
	model "gotest/src/main/model"
	"net/http"
)

func HandlePerson(w http.ResponseWriter, r *http.Request) {
	person := model.Person{
		Name:    "imran",
		Address: "Bangalore or Bihar",
	}

	farhan := model.Person{
		Name:    "Farhan Khan",
		Address: "Bihar Darbhanga",
	}

	query := r.URL.Query().Get("query")
	var result any

	switch query {
	case "farhan":
		result = farhan
	case "all":
		result = []model.Person{person, farhan}
	default:
		result = person
	}

	// out, err := json.MarshalIndent(person, "", " ")
	out, err := json.MarshalIndent(result, "", " ")

	if err != nil {
		fmt.Fprintln(w, "Error to parse person")

	}
	w.Write(out)

}
