package main

import (
	"log"
	"net/http"
)

const (
	StationsURL = "/stations"
)

type Station struct {
	Id      string   `json:"id"`
	Type    string   `json:"type"`
	Modules []Module `json:"modules"`
}

type Module struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func GetStations(w http.ResponseWriter, r *http.Request) {
	var station = Station{}
	netatmoAuth := authenticateToNetatmo()
	stationData := getStationData(netatmoAuth)
	for _, stationElement := range stationData.Body.Devices {
		station.Id = stationElement.Id
		station.Type = stationElement.Type
		for _, moduleElement := range stationElement.Modules {
			var module = Module{}
			module.Id = moduleElement.Id
			module.Type = moduleElement.Type
			station.Modules = append(station.Modules, module)
		}
	}

	log.Println(station)
}
