package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	StationsURL = "/stations"
	ModulesUrl  = "/modules"
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
	currentTime := time.Now().UnixNano()
	if nerdatmo.DataTime-currentTime > DataExpiry {
		log.Println("Refreshing data")
		netatmoAuth := authenticateToNetatmo()
		nerdatmo.StationData = getStationData(netatmoAuth)
		nerdatmo.DataTime = currentTime
	}
	// @todo check local time and double check against cache creation
	var station = new(Station)
	for _, stationElement := range nerdatmo.StationData.Body.Devices {
		station.Id = stationElement.Id
		station.Type = stationElement.Type
		for _, moduleElement := range stationElement.Modules {
			var module = new(Module)
			module.Id = moduleElement.Id
			module.Type = moduleElement.Type
			station.Modules = append(station.Modules, *module)
		}
	}
	// @todo json response
	log.Println(station)
}

func GetModules(w http.ResponseWriter, r *http.Request) {
	var modules []Module
	stationId := mux.Vars(r)["stationId"]
	for _, stationElement := range nerdatmo.StationData.Body.Devices {
		if stationElement.Id == stationId {
			for _, moduleElement := range stationElement.Modules {
				var module = new(Module)
				module.Id = moduleElement.Id
				module.Type = moduleElement.Type
				modules = append(modules, *module)
			}
		}
	}
	log.Println(modules)
}

func GetModule(w http.ResponseWriter, r *http.Request) {
	stationId := mux.Vars(r)["stationId"]
	moduleId := mux.Vars(r)["moduleId"]

	log.Println("Requesting data for station " + stationId + " and module " + moduleId)
}
