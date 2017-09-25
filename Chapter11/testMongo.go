package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"time"
)

type Record struct {
	Xvalue int
	Yvalue int
}

func main() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{"127.0.0.1:27017"},
		Timeout: 20 * time.Second,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Printf("DialWithInfo: %s\n", err)
		os.Exit(100)
	}
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("goDriver").C("someData")

	err = collection.Insert(&Record{1, 0})
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	err = collection.Insert(&Record{-1, 0})
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	var recs []Record
	err = collection.Find(bson.M{"yvalue": 0}).All(&recs)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for x, y := range recs {
		fmt.Println(x, y)
	}
	fmt.Println("Found:", len(recs), "results!")
}
