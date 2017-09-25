package main

import (
	"fmt"
	"net/http"
	"time"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the /about page at %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func cv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the /CV page at %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	title := currentTime
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1><h2 align=\"center\">%s</h2>", Body, title)
	fmt.Printf("Served: %s for %s\n", r.URL.Path, r.Host)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Welcome to my home page!\n")
	} else {
		fmt.Fprintf(w, "Unknown page: %s from %s\n", r.URL.Path, r.Host)
	}
	fmt.Printf("Served: %s for %s\n", r.URL.Path, r.Host)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/about", about)
	m.HandleFunc("/CV", cv)
	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", home)

	http.ListenAndServe(":8001", m)
}
