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
)

type dataBase struct{}

func (db *dataBase) connect() {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	connectedDb := session.DB(dbName)
	peopleCollection = connectedDb.C("people")
	companiesCollection = connectedDb.C("companies")
}
