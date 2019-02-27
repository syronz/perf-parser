package mongo

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"github.com/syronz/perf-parser/models"
	//"gopkg.in/mgo.v2/bson"
)

type mongoDB struct {
	session		*mgo.Session
	collection	*mgo.Collection
}

type People struct {
	Name string
	Phone string
	Friends map[string]int
}

func Insert(d model.Data) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("perf_parse").C("data")


	err = c.Insert(d)
	if err != nil {
		log.Fatal(err)
	}
}
