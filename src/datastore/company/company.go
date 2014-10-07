package company

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type (
	CompaniesRepo struct{}

	Company struct {
		Id          bson.ObjectId `bson:"_id" json:"id"`
		UserId      bson.ObjectId `bson:"userId" json:"_"`
		Name        string        `bson:"name" json:"name"`
		Email       string        `bson:"email" json:"email"`
		TelNumber   string        `bson:"telNumber" json:"telNumber"`
		Information string        `bson:"information" json:"information"`
	}
)

func New() *Company {
	return &Company{}
}

func CreateCompaniesRepo() *CompaniesRepo {
	return &CompaniesRepo{}
}

func (c *Company) Save() {
	_, err := datastore.CompaniesCollection.Upsert(bson.M{"_id": c.Id}, c)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (repo *CompaniesRepo) FetchCompanyProfiles(id bson.ObjectId) *[]Company {
	var results []Company
	err := datastore.CompaniesCollection.Find(bson.M{"userId": id}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return &results
}

func (repo *CompaniesRepo) All() []Company {
	var results []Company
	err := datastore.CompaniesCollection.Find(nil).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return results
}
