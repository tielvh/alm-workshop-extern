package main

import (
	"encoding/json"
	"net/http"
)

type Workshop struct {
	Name         string   `json:"name"`
	Date         string   `json:"date"`
	Presentator  string   `json:"presentator"`
	Participants []string `json:"participants"`
}

var workshop = Workshop{
	Name:         "ALM Workshop",
	Date:         "14/02/2025",
	Presentator:  "AE Consultants",
	Participants: []string{"John Doe", "Mary Little Lamb", "Chuck Norris"},
}

func getWorkshopHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the struct to JSON and write it to the response
	json.NewEncoder(w).Encode(workshop)
}

func postWorkshopHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON data into a new Workshop struct
	var newWorkshop Workshop
	err := json.NewDecoder(r.Body).Decode(&newWorkshop)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON data"))
		return
	}

	// Update the workshop details
	workshop = newWorkshop

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workshop)
}

func WorkshopHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getWorkshopHandler(w, r)
	case "POST":
		postWorkshopHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
