package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Welcome to the Go API Server! May the code be with you!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/about" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "This is the About page. Learn more about our awesome API!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	log.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", mux)
	if err!= nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}