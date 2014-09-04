package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	companiesRepo struct{}

	company struct {
		Id          bson.ObjectId `bson:"_id" json:"id"`
		UserId      bson.ObjectId `bson:"userId" json:"_"`
		Name        string        `bson:"name" json:"name"`
		Email       string        `bson:"email" json:"email"`
		TelNumber   string        `bson:"telNumber" json:"telNumber"`
		Information string        `bson:"information" json:"information"`
	}
)

func createCompaniesRepo() *companiesRepo {
	repo := new(companiesRepo)
	return repo
}

func (c *company) save() {
	_, err := companiesCollection.Upsert(bson.M{"_id": c.Id}, c)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (repo *companiesRepo) fetchCompanyProfiles(id bson.ObjectId) *[]company {
	var results []company
	err := companiesCollection.Find(bson.M{"userId": id}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return &results
}
