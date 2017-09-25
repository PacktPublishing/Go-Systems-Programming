package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename string, key interface{}) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	encodeJSON := json.NewEncoder(out)
	err = encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	out.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		os.Exit(100)
	}

	filename := arguments[1]
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		}}

	saveToJSON(filename, myRecord)
}
