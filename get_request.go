package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	http.HandleFunc("/person", getPerson)
	http.ListenAndServe(":8080", nil)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	// Create a new Person object with some sample data
	person := Person{FirstName: "John", LastName: "Doe"}

	// Marshal the Person object into a JSON byte slice
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
