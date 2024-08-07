package db

import (
	model "gotest/src/main/model"
)

type singleton []model.Person

var instance singleton

func getInstance() singleton {
	if instance == nil {
		// instance = make(singleton, 5)
		instance = singleton{}
	}

	return instance
}

func GetAll() singleton {
	return getInstance()

}

func GetPersonById(id int) model.Person {

	for _, person := range getInstance() {
		if person.Id == id {
			return person
		}
	}
	return model.Person{}
}

func PostPerson(person model.Person) singleton {
	instance = append(getInstance(), person)
	return getInstance()
}

func GetSize() int {
	return len(getInstance())
}
