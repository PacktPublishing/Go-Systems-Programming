package main

import (
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"net/http"
	"os"
	"time"
)

var DatabaseName string
var collectionName string

type Document struct {
	P1 int
	P2 int
	P3 int
	P4 int
	P5 int
}

func content(w http.ResponseWriter, r *http.Request) {
	var Data []Document
	myT := template.Must(template.ParseGlob("mongoDB.gohtml"))

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{"127.0.0.1:27017"},
		Timeout: 20 * time.Second,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Printf("DialWithInfo: %s\n", err)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DatabaseName).C(collectionName)

	err = c.Find(nil).All(&Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found:", len(Data), "results!")
	myT.ExecuteTemplate(w, "mongoDB.gohtml", Data)
}

func main() {
	arguments := os.Args

	if len(arguments) <= 2 {
		fmt.Println("Please provide a Database and a Collection!")
		os.Exit(100)
	} else {
		DatabaseName = arguments[1]
		collectionName = arguments[2]
	}

	http.HandleFunc("/", content)
	http.ListenAndServe(":8001", nil)
}
