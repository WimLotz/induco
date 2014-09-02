package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	peopleRepo struct{}

	person struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"_"`
		GoogleAuthId string        `bson:"googleAuthId,omitempty" json:"googleAuthId,omitempty"`
		FirstName    string        `bson:"firstName" json:"firstName"`
		Surname      string        `bson:"surname" json:"surname"`
		Email        string        `bson:"email" json:"emailAddress"`
		NeedWork     bool          `bson:"needWork" json:"needWork"`
		NeedHelp     bool          `bson:"needHelp" json:"needHelp"`
		WorkExp      string        `bson:"workExp" json:"workExp,omitempty"`
	}
)

func createPeopleRepo() *peopleRepo {
	repo := new(peopleRepo)
	return repo
}

func (repo *peopleRepo) createPerson(p person) {
	err := peopleCollection.Insert(p)
	if err != nil {
		log.Printf("Can't create person: %v\n", err)
	}
}

func (repo *peopleRepo) updatePerson(p person) {
	change := bson.M{"$set": bson.M{"firstName": p.FirstName, "surname": p.Surname, "email": p.Email, "needWork": p.NeedWork, "needHelp": p.NeedHelp, "workExp": p.WorkExp}}
	err := peopleCollection.UpdateId(p.Id, change)
	if err != nil {
		log.Printf("unable to update record: %v\n", err)
	}
}

func (repo *peopleRepo) fetchProfile(id bson.ObjectId) *person {
	var p person
	err := peopleCollection.Find(bson.M{"_id": id}).One(&p)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return &p
}

func (repo *peopleRepo) fetchObjIdOnGooglePlusId(id string) bson.ObjectId {
	var p person
	err := peopleCollection.Find(bson.M{"googleAuthId": id}).One(&p)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return p.Id
}
