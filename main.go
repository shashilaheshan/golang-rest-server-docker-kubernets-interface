package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Name: "The Go Programming Language"},
	{ID: "2", Name: "Clean Code"},
	{ID: "3", Name: "You Don't Know JS"},
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}
func main() {

	http.HandleFunc("/books", getBooks)

	// Start the HTTP server
	port := 8084
	fmt.Println("running the server on %s", port)
	// Listen on port and handle errors
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatalf("Error occured while connecting %s", err)
	}
}
