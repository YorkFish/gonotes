package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var MgoConnect *mgo.Collection

func MD() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Println(">>> can not connect mongodb")
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	MgoConnect = session.DB("gintest").C("users")
}
