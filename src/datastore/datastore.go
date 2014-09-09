package datastore

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var (
	uri                 = "mongodb://wim:test@kahana.mongohq.com:10082/induco"
	dbName              = "induco"
	PeopleCollection    *mgo.Collection
	CompaniesCollection *mgo.Collection
	UsersCollection     *mgo.Collection
)

type DataBase struct{}

func New() (db *DataBase) {
	return &DataBase{}
}

func (db *DataBase) Connect() {
	dbSession, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo:  %v\n", err)
		os.Exit(1)
	}
	dbSession.SetSafe(&mgo.Safe{})
	connectedDb := dbSession.DB(dbName)
	PeopleCollection = connectedDb.C("people")
	CompaniesCollection = connectedDb.C("companies")
	UsersCollection = connectedDb.C("users")
}
