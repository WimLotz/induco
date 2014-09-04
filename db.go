package main

import (
	"labix.org/v2/mgo"
	"log"
	"os"
)

var (
	uri                 = "mongodb://wim:test@kahana.mongohq.com:10082/induco"
	dbName              = "induco"
	peopleCollection    *mgo.Collection
	companiesCollection *mgo.Collection
	usersCollection     *mgo.Collection
)

type dataBase struct{}

func (db *dataBase) connect() {
	dbSession, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo:  %v\n", err)
		os.Exit(1)
	}
	dbSession.SetSafe(&mgo.Safe{})
	connectedDb := dbSession.DB(dbName)
	peopleCollection = connectedDb.C("people")
	companiesCollection = connectedDb.C("companies")
	usersCollection = connectedDb.C("users")
}
