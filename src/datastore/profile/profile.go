package profile

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Profile struct {
	Id             bson.ObjectId `bson:"_id" json:"id"`
	UserId         bson.ObjectId `bson:"userId" json:"userId"`
	IsCompany      bool          `bson:"isCompany" json:"isCompany"`
	NeedWork       bool          `bson:"needWork,omitempty" json:"needWork,omitempty"`
	LookingForWork bool          `bson:"lookingForWork,omitempty" json:"lookingForWork,omitempty"`
	CompanyName    string        `bson:"companyName,omitempty" json:"companyName,omitempty"`
	FirstName      string        `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Surname        string        `bson:"surname,omitempty" json:"surname,omitempty"`
	Bio            string        `bson:"bio,omitempty" json:"bio,omitempty"`
	EmailAddress   string        `bson:"emailAddress,omitempty" json:"emailAddress,omitempty"`
	MobileNumber   string        `bson:"mobileNumber,omitempty" json:"mobileNumber,omitempty"`
	WorkNumber     string        `bson:"workNumber,omitempty" json:"workNumber,omitempty"`
	HomeNumber     string        `bson:"homeNumber,omitempty" json:"homeNumber,omitempty"`
	WorkExpTags    []string      `bson:"workExpTags,omitempty" json:"workExpTags,omitempty"`
	NeededExpTags  []string      `bson:"neededExpTags,omitempty" json:"neededExpTags,omitempty"`
}

func New() *Profile {
	return &Profile{}
}

func (p *Profile) Save() {
	_, err := datastore.ProfileCollection.Upsert(bson.M{"_id": p.Id}, p)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (p *Profile) Fetch(userId bson.ObjectId) []Profile {
	var results []Profile
	err := datastore.ProfileCollection.Find(bson.M{"userId": userId}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return results
}

func (p *Profile) All() []Profile {
	var results []Profile
	err := datastore.ProfileCollection.Find(nil).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return results
}
