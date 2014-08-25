package main

import (
	"log"
)

type (
	peopleRepo struct {
		dataBase
	}
	person struct {
		Id        string `json:"id" 			bson:"_id"`
		FirstName string `json:"firstName"		bson:"firstName"`
		Surname   string `json:"surname"		bson:"surname"`
	}
)

func createPeopleRepo() *peopleRepo {
	repo := new(peopleRepo)
	return repo
}

func (repo *peopleRepo) createPerson(p person) {
	db := repo.connect()
	collection := db.C("people")
	err := collection.Insert(p)
	if err != nil {
		log.Printf("Can't create person: %v\n", err)
	}
}
