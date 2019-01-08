package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

const (
	APIUrl     = "https://api.netatmo.com"                              // Netatmo API Endpoint
	UserAgent  = "nerdatmo/v0.1 (https://github.com/acorsinl/nerdatmo)" // User-Agent for requests
	ListenPort = "6666"
	Version    = "v0.1"
	DataExpiry = 1800 // How long is the cache valid
)

type Nerdatmo struct {
	StationData *NetatmoResponse
	DataTime    int64
}

var nerdatmo Nerdatmo

func main() {
	// Load environment configuration
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error reading environment configuration: " + err.Error())
	}

	// Authenticate to Netatmo and cache station data
	netatmoAuth := authenticateToNetatmo()
	stationData := getStationData(netatmoAuth)
	log.Println(stationData)

	// Preload all info for faster querying
	nerdatmo.StationData = stationData
	nerdatmo.DataTime = time.Now().UnixNano()

	r := mux.NewRouter()
	r.HandleFunc(StationsURL, GetStations).Methods("GET")
	r.HandleFunc(StationsURL+"/{stationId}"+ModulesUrl, GetModules).Methods("GET")
	r.HandleFunc(StationsURL+"/{stationId}"+ModulesUrl+"/{moduleId}", GetModule).Methods("GET")
	http.Handle("/", r)

	log.Println("Nerdatmo server " + Version + " running on port " + ListenPort)
	log.Fatal(http.ListenAndServe(":"+ListenPort, http.DefaultServeMux))
}
