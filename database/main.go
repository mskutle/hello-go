package database

import (
	"fmt"
	"github.com/globalsign/mgo"
)

var dbSession *mgo.Session

func init() {
	session, err := mgo.Dial("mongodb://root:root@mongo:27017/hello-go?authSource=admin")

	if err != nil {
		panic(err)
	}
	dbSession = session
	fmt.Println("Dial success!")
}

func GetDatabase() *mgo.Database {
	return dbSession.DB("")
}

func GetDatabaseCloned() *mgo.Database {
	return dbSession.Clone().DB("")
}