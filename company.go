package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	companiesRepo struct{}

	company struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"_"`
		GoogleAuthId string        `bson:"googleAuthId,omitempty" json:"googleAuthId,omitempty"`
		Name         string        `bson:"name" json:"name"`
		Email        string        `bson:"email" json:"email"`
		TelNumber    string        `bson:"telNumber" json:"telNumber"`
		Information  string        `bson:"information" json:"information"`
	}
)

func createCompaniesRepo() *companiesRepo {
	repo := new(companiesRepo)
	return repo
}

func (repo *companiesRepo) update(c company) {
	//change := bson.M{"$set": bson.M{"firstName": p.FirstName, "surname": p.Surname, "email": p.Email, "needWork": p.NeedWork, "needHelp": p.NeedHelp, "workExp": p.WorkExp}}
	//err := companiesCollection.UpdateId(p.Id, change)
	//if err != nil {
	//	log.Printf("unable to update record: %v\n", err)
	//}
}

func (repo *companiesRepo) fetchProfile(id bson.ObjectId) *company {
	var c company
	err := peopleCollection.Find(bson.M{"_id": id}).One(&c)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return &c
}
