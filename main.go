package main

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	APIUrl    = "https://api.netatmo.com"                              //Netatmo API Endpoint
	UserAgent = "nerdatmo/v0.1 (https://github.com/acorsinl/nerdatmo)" // User-Agent for requests
)

func main() {
	// Load environment configuration
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error reading environment configuration: " + err.Error())
	}

	// Authenticate to Netatmo and grab station data
	netatmoAuth := authenticateToNetatmo()
	stationData := getStationData(netatmoAuth)
	log.Println(stationData)
}
