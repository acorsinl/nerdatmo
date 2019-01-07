package main

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	APIUrl = "https://api.netatmo.com" //Netatmo API Endpoint
)

func main() {
	// Load environment configuration
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error reading environment configuration: " + err.Error())
	}
}
