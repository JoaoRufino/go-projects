package common

import mgo "gopkg.in/mgo.v2"

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

func connectDB() {
	session, err := mgo.Dial("172.17.0.2")
	if err != nil {
		panic(err)
	}

	DB = &mongo{session.DB("taskdb").C("tasks")}

}
