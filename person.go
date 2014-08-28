package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	peopleRepo struct {
		dataBase
	}
	person struct {
		Id           bson.ObjectId `json:"id" 				bson:"_id"`
		ValidationId string        `json:"validationId"		bson:"validationId"`
		FirstName    string        `json:"firstName"		bson:"firstName"`
		Surname      string        `json:"surname"			bson:"surname"`
		Email        string        `json:"emailAddress"		bson:"email"`
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
