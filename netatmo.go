package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// NetatmoAuth struct deals with Authentication Tokens
type NetatmoAuth struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	ExpiresIn    int      `json:"expires_in"`
	ExpireIn     int      `json:"expire_in"`
}

// NetatmoResponse deals with API requests
type NetatmoResponse struct {
	Body Body `json:"body"`
}

// Body holds API request bodies
type Body struct {
	Devices    []Devices `json:"devices"`
	User       User      `json:"user`
	Status     string    `json:"status"`
	TimeExec   float64   `json:"time_exec"`
	TimeServer float64   `json:"time_server"`
}

// Devices holds devices information
type Devices struct {
	Id              string        `json:"_id"`
	CipherId        string        `json:"cipher_id"`
	DateSetup       int           `json:"date_setup"`
	LastSetup       int           `json:"last_setup"`
	Type            string        `json:"type"`
	LastStatusStore int           `json:"last_status_store"`
	ModuleName      string        `json:"module_name"`
	Firmware        int           `json:"firmware"`
	LastUpgrade     int           `json:"last_upgrade"`
	WifiStatus      int           `json:"wifi_status"`
	Reachable       bool          `json:"reachable"`
	CO2Calibrating  bool          `json:"co2_calibrating`
	StationName     string        `json:"station_name"`
	DataType        []string      `json:"data_type"`
	Place           Place         `json:"place"`
	DashboardData   DashboardData `json:"dashboard_data"`
	Modules         []Modules     `json:"modules"`
}

// Place holds place information
type Place struct {
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Timezone string    `json:"timezone"`
	Location []float32 `json:"location"`
}

// DashboardData holds dashboard data information
type DashboardData struct {
	TimeUTC          int     `json:"time_utc"`
	Temperature      float32 `json:"Temperature"`
	CO2              int     `json:"CO2"`
	Humidity         int     `json:"Humidity"`
	Noise            int     `json:"Noise"`
	Pressure         float32 `json:"Pressure"`
	AbsolutePressure float32 `json:"AbsolutePressure"`
	MinTemp          float32 `json:"min_temp"`
	MaxTemp          float32 `json:"max_temp"`
	DateMinTemp      int     `json:"date_min_temp"`
	DateMaxTemp      int     `json:"date_max_temp"`
	PressureTrend    string  `json:"pressure_trend"`
}

// Modules holds modules information
type Modules struct {
	Id             string              `json:"_id"`
	Type           string              `json:"type"`
	ModuleName     string              `json:"module_name"`
	DataType       []string            `json:"data_type"`
	LastSetup      int                 `json:"last_setup"`
	Reachable      bool                `json:"reachable"`
	DashboardData  DashboardDataModule `json:"dashboard_data"`
	Firmware       int                 `json:"firmware"`
	LastMessage    int                 `json:"last_message"`
	LastSeen       int                 `json:"last_seen"`
	RfStatus       int                 `json:"rf_status"`
	BatteryVP      int                 `json:"battery_vp"`
	BatteryPercent int                 `json:"battery_percent"`
}

type DashboardDataModule struct {
	TimeUTC     int     `json:"time_utc"`
	Temperature float32 `json:"Temperature"`
	CO2         int     `json:"CO2,omitempty"`
	Humidity    int     `json:"Temperature"`
	MinTemp     float32 `json:"min_temp"`
	MaxTemp     float32 `json:"max_temp"`
	DateMinTemp int     `json:"date_min_temp"`
	DateMaxTemp int     `json:"date_max_temp"`
}

// User holds user data
type User struct {
	Mail           string         `json:"mail"`
	Administrative Administrative `json:"administrative"`
}

// Administrative holds administrative data
type Administrative struct {
	Lang         string `json:"lang"`
	RegLocale    string `json:"reg_locale"`
	Unit         int    `json:"unit"`
	WindUnit     int    `json:"windunit"`
	PressureUnit int    `json:"pressureunit"`
	FeelLikeAlgo int    `json:"feel_like_algo"`
}

func authenticateToNetatmo() *NetatmoAuth {
	var netatmoAuth = new(NetatmoAuth)

	// Set form values
	payload := url.Values{}
	payload.Set("grant_type", "password")
	payload.Set("scope", "read_station")
	payload.Set("client_id", os.Getenv("CLIENT_ID"))
	payload.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	payload.Set("username", os.Getenv("NETATMO_USERNAME"))
	payload.Set("password", os.Getenv("NETATMO_PASSWORD"))

	// Perform HTTP request
	req, _ := http.NewRequest("POST", APIUrl+"/oauth2/token", strings.NewReader(payload.Encode()))
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("HTTP request failed: " + err.Error())
	}

	// Parse response and unmarshal
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &netatmoAuth); err != nil {
		log.Println("Error unmarshalling: " + err.Error())
	}
	return netatmoAuth
}

func getStationData(netatmoAuth *NetatmoAuth) *NetatmoResponse {
	var netatmoResponse = new(NetatmoResponse)

	// Set form values
	payload := url.Values{}
	payload.Set("access_token", netatmoAuth.AccessToken)

	// Perform HTTP request
	req, _ := http.NewRequest("POST", APIUrl+"/api/getstationsdata", strings.NewReader(payload.Encode()))
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("HTTP request failed: " + err.Error())
	}

	// Parse response and unmarshal
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &netatmoResponse); err != nil {
		log.Println("Error unmarshalling: " + err.Error())
	}
	return netatmoResponse
}
