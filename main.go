package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

const (
	APIUrl     = "https://api.netatmo.com"                              //Netatmo API Endpoint
	UserAgent  = "nerdatmo/v0.1 (https://github.com/acorsinl/nerdatmo)" // User-Agent for requests
	ListenPort = "6666"
	Version    = "v0.1"
)

func main() {
	// Load environment configuration
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error reading environment configuration: " + err.Error())
	}

	/* // Authenticate to Netatmo and grab station data
	netatmoAuth := authenticateToNetatmo()
	stationData := getStationData(netatmoAuth)
	log.Println(stationData) */

	r := mux.NewRouter()
	r.HandleFunc(StationsURL, GetStations).Methods("GET")
	http.Handle("/", r)

	log.Println("Nerdatmo server " + Version + " running on port " + ListenPort)
	log.Fatal(http.ListenAndServe(":"+ListenPort, http.DefaultServeMux))
}
