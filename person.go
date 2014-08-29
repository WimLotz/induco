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
		Id           bson.ObjectId `bson:"_id,omitempty" json:"_"`
		GoogleAuthId string        `bson:"googleAuthId,omitempty" json:"googleAuthId,omitempty"`
		FirstName    string        `bson:"firstName" json:"firstName"`
		Surname      string        `bson:"surname" json:"surname"`
		Email        string        `bson:"email" json:"emailAddress"`
		NeedWork     bool          `bson:"needWork" json:"needWork"`
		NeedHelp     bool          `bson:"needHelp" json:"needHelp"`
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

func (repo *peopleRepo) updatePerson(p person) {
	db := repo.connect()
	collection := db.C("people")
	change := bson.M{"$set": bson.M{"firstName": p.FirstName, "surname": p.Surname, "email": p.Email, "needWork": p.NeedWork, "needHelp": p.NeedHelp}}
	err := collection.UpdateId(p.Id, change)
	if err != nil {
		log.Printf("unable to update record: %v\n", err)
	}
}

func (repo *peopleRepo) fetchObjIdOnGooglePlusId(id string) bson.ObjectId {
	db := repo.connect()
	collection := db.C("people")
	var p person
	err := collection.Find(bson.M{"googleAuthId": id}).One(&p)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return p.Id
}
