package hander

import (
	"encoding/json"
	"fmt"
	db "gotest/src/main/db"
	model "gotest/src/main/model"
	"net/http"
	"strconv"
)

func HandlePerson(w http.ResponseWriter, r *http.Request) {

	var result any
	switch r.Method {
	case "GET":
		query, error := strconv.Atoi(r.URL.Query().Get("id"))

		if error == nil {
			result = db.GetAll()
		} else {
			result = db.GetPersonById(query)
		}

	case "POST":

		var p model.Person
		// parse body in model
		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			result = "enter correct name and address"

		} else {

			person := model.Person{Name: p.Name, Address: p.Address, Id: db.GetSize() + 1}
			db.PostPerson(person)
			result = db.GetAll()
		}

	default:

		result = "Sorry, only GET and POST methods are supported."
	}

	out, err := json.MarshalIndent(result, "", " ")

	if err != nil {
		fmt.Fprintln(w, "500 something went wrong")
	}
	w.Write(out)

}
