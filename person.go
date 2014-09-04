package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	peopleRepo struct{}

	person struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId    bson.ObjectId `bson:"userId,omitempty" json:"_"`
		FirstName string        `bson:"firstName" json:"firstName"`
		Surname   string        `bson:"surname" json:"surname"`
		Email     string        `bson:"email" json:"emailAddress"`
		NeedWork  bool          `bson:"needWork" json:"needWork"`
		NeedHelp  bool          `bson:"needHelp" json:"needHelp"`
		WorkExp   string        `bson:"workExp" json:"workExp,omitempty"`
	}
)

func createPeopleRepo() *peopleRepo {
	repo := new(peopleRepo)
	return repo
}

func (p *person) save() {
	_, err := peopleCollection.Upsert(bson.M{"_id": p.Id}, p)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (repo *peopleRepo) fetchPersonProfiles(userId bson.ObjectId) *[]person {
	var results []person
	err := peopleCollection.Find(bson.M{"userId": userId}).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return &results
}
