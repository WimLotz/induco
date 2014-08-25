package main

import (
	"labix.org/v2/mgo"
	"log"
	//"labix.org/v2/mgo/bson"
	"os"
)

var (
	uri    = "mongodb://wim:test@kahana.mongohq.com:10082/induco"
	dbName = "induco"
)

type dataBase struct{}

func (db *dataBase) connect() *mgo.Database {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	connectedDb := session.DB(dbName)
	return connectedDb
}
