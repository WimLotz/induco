package person

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type (
	PeopleRepo struct{}

	Person struct {
		Id          bson.ObjectId `bson:"_id" json:"id"`
		UserId      bson.ObjectId `bson:"userId" json:"_"`
		FirstName   string        `bson:"firstName" json:"firstName"`
		Surname     string        `bson:"surname" json:"surname"`
		Email       string        `bson:"email" json:"emailAddress"`
		NeedWork    bool          `bson:"needWork" json:"needWork"`
		NeedHelp    bool          `bson:"needHelp" json:"needHelp"`
		WorkExpTags []string      `bson:"workExpTags" json:"workExpTags"`
		Bio         string        `bson:"bio" json:"bio"`
	}
)

func New() *Person {
	return &Person{}
}

func CreatePeopleRepo() *PeopleRepo {
	return &PeopleRepo{}
}

func (p *Person) Save() {
	_, err := datastore.PeopleCollection.Upsert(bson.M{"_id": p.Id}, p)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (repo *PeopleRepo) FetchPersonProfiles(userId bson.ObjectId) *[]Person {
	var results []Person
	err := datastore.PeopleCollection.Find(bson.M{"userId": userId}).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return &results
}
