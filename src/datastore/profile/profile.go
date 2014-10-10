package profile

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
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

func New() *Profile {
	return &Profile{}
}

func (p *Profile) Save() {
	_, err := datastore.EntityCollection.Upsert(bson.M{"_id": e.Id}, c)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (p *Profile) Fetch(id bson.ObjectId) *[]Profile {
	var results []Profile
	err := datastore.EntityCollection.Find(bson.M{"userId": id}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return &results
}

func (p *Profile) All() []Profile {
	var results []Profile
	err := datastore.ProfileCollection.Find(nil).All(&results)
	if err != nil {
		log.Printf("No records found: %v\n", err)
	}
	return results
}
