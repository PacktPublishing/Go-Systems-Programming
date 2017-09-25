package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "You do not give an argument!"
	} else {
		myString = arguments[1]
	}

	buf := []byte(myString)
	io.WriteString(os.Stdout, string(buf))
	io.WriteString(os.Stdout, "\n")
}
