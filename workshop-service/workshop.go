package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Workshop struct {
	Name         string   `json:"name"`
	Date         string   `json:"date"`
	Presentator  string   `json:"presentator"`
	Participants []string `json:"participants"`
	SweaterScore int8     `json:"sweaterscore"`
}

var defaultSweaterScore, _ = strconv.ParseInt(os.Getenv("DEFAULT_SWEATERSCORE"), 10, 8)

var workshop = Workshop{
	Name:         "ALM Workshop",
	Date:         "14/02/2025",
	Presentator:  "AE Consultants",
	Participants: []string{"John Doe", "Mary Little Lamb", "Chuck Norris"},
	SweaterScore: int8(defaultSweaterScore),
}

func getWorkshopHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the struct to JSON and write it to the response
	json.NewEncoder(w).Encode(workshop)
}

func isValidSweaterScore(score int8) bool {
	return score >= 1 && score <= 10
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

	if !isValidSweaterScore(newWorkshop.SweaterScore) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("SweaterScore must be between 1 and 10"))
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