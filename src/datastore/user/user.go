package user

import (
	"datastore"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type (
	UsersRepo struct{}

	User struct {
		Id           bson.ObjectId `bson:"_id" json:"_"`
		GoogleAuthId string        `bson:"googleAuthId" json:"googleAuthId"`
	}
)

func New(id bson.ObjectId, googleAuthId string) *User {
	return &User{
		Id:           id,
		GoogleAuthId: googleAuthId,
	}
}

func CreateUsersRepo() *UsersRepo {
	return &UsersRepo{}
}

func (repo *UsersRepo) CreateUser(u *User) {
	err := datastore.UsersCollection.Insert(u)
	if err != nil {
		log.Printf("Can't create user: %v\n", err)
	}
}

func (repo *UsersRepo) FetchUserIdFromGooglePlusId(id string) bson.ObjectId {
	var u User
	err := datastore.UsersCollection.Find(bson.M{"googleAuthId": id}).One(&u)
	if err != nil {
		log.Printf("no record found: %v\n", err)
	}
	return u.Id
}
