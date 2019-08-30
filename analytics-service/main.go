package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	router.HandleFunc("/identify", Identify).Methods("POST")
	router.HandleFunc("/track", Track).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// HealthCheck -- To provide a response for the healthcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive")
}

// Identify -- To handle identify calls from the client
func Identify(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "")
	}
	w.WriteHeader(http.StatusCreated)
}

// Track -- To handle track calls from the client
func Track(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "")
	}
	w.WriteHeader(http.StatusCreated)
}
