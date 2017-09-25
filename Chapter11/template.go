package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Entry struct {
	WebSite string
	WebName string
	Quality string
}

var filename string

func dynamicContent(w http.ResponseWriter, r *http.Request) {
	var Data []Entry
	var f *os.File
	if filename == "" {
		f = os.Stdin
	} else {
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	myT := template.Must(template.ParseGlob("template.gohtml"))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		parts := strings.Fields(scanner.Text())
		if len(parts) == 3 {
			temp := Entry{WebSite: parts[0], WebName: parts[1], Quality: parts[2]}
			Data = append(Data, temp)
		}
	}

	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT.ExecuteTemplate(w, "template.gohtml", Data)
}

func staticPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("static.gohtml"))
	myT.ExecuteTemplate(w, "static.gohtml", nil)
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		filename = ""
	} else {
		filename = arguments[1]
	}

	http.HandleFunc("/static", staticPage)
	http.HandleFunc("/dynamic", dynamicContent)
	http.ListenAndServe(":8001", nil)
}
