package user

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type (
	UsersRepo struct{}

	User struct {
		Id       bson.ObjectId `bson:"_id" json:"id"`
		Email    string        `bson:"email" json:"email"`
		Password string        `bson:"password" json:"password"`
	}
)

func New() *User {
	return &User{}
}

func (u *User) Save() {
	_, err := datastore.UsersCollection.Upsert(bson.M{"_id": u.Id}, u)
	if err != nil {
		log.Printf("Unable to save record: %v\n", err)
	}
}

func (u *User) Fetch(userId bson.ObjectId) User {
	var results User
	err := datastore.UsersCollection.Find(bson.M{"_id": userId}).All(&results)
	if err != nil {
		log.Printf("no records found: %v\n", err)
	}
	return results
}
