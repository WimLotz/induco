package main

import (
	"labix.org/v2/mgo/bson"
	"log"
)

type (
	usersRepo struct{}

	user struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"_"`
		GoogleAuthId string        `bson:"googleAuthId,omitempty" json:"googleAuthId,omitempty"`
	}
)

func createUsersRepo() *usersRepo {
	repo := new(usersRepo)
	return repo
}

func (repo *usersRepo) createUser(u user) {
	err := usersCollection.Insert(u)
	if err != nil {
		log.Printf("Can't create user: %v\n", err)
	}
}

func (repo *usersRepo) fetchUserIdFromGooglePlusId(id string) bson.ObjectId {
	var u user
	err := usersCollection.Find(bson.M{"googleAuthId": id}).One(&u)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return u.Id
}
