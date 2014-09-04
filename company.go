package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	companiesRepo struct{}

	company struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"_"`
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
	_, err := companiesCollection.UpsertId(c.Id, c)
	if err != nil {
		log.Printf("unable to save record: %v\n", err)
	}
}

func (repo *companiesRepo) fetchCompanyProfile(id bson.ObjectId) *company {
	var c company
	err := companiesCollection.Find(bson.M{"_id": id}).One(&c)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return &c
}
