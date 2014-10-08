package entity

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	Id                 bson.ObjectId `bson:"_id" json:"id"`
	UserId             bson.ObjectId `bson:"userId" json:"userId"`
	IsCompany          bool          `bson:"isCompany" json:"isCompany"`
	NeedWork           bool          `bson:"needWork,omitempty" json:"needWork,omitempty"`
	LookingForWork     bool          `bson:"lookingForWork,omitempty" json:"lookingForWork,omitempty"`
	CompanyName        string        `bson:"companyName,omitempty" json:"companyName,omitempty"`
	FirstName          string        `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Surname            string        `bson:"surname,omitempty" json:"surname,omitempty"`
	PersonalBio        string        `bson:"personalBio,omitempty" json:"personalBio,omitempty"`
	CompanyBio         string        `bson:"companyBio,omitempty" json:"companyBio,omitempty"`
	EmailAddress       string        `bson:"emailAddress,omitempty" json:"emailAddress,omitempty"`
	MobileNumber       string        `bson:"mobileNumber,omitempty" json:"mobileNumber,omitempty"`
	WorkNumber         string        `bson:"workNumber,omitempty" json:"workNumber,omitempty"`
	HomeNumber         string        `bson:"homeNumber,omitempty" json:"homeNumber,omitempty"`
	WorkExperienceTags []string      `bson:"homeNumber,omitempty" json:"workExperienceTags,omitempty"`
	NeededExperience   []string      `bson:"neededExperience,omitempty" json:"neededExperience,omitempty"`
}

func New() *Entity {
	return &Entity{}
}

func (e *Entity) Save() {
	_, err := datastore.EntityCollection.Upsert(bson.M{"_id": e.Id}, c)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (e *Entity) Fetch(id bson.ObjectId) *[]Entity {
	var results []Entity
	err := datastore.EntityCollection.Find(bson.M{"userId": id}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return &results
}

func (e *Entity) All() []Entity {
	var results []Entity
	err := datastore.EntityCollection.Find(nil).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return results
}
