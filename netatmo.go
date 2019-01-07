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

type NetatmoAuth struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	ExpiresIn    int      `json:"expires_in"`
	ExpireIn     int      `json:"expire_in"`
}

func authenticateToNetatmo() *NetatmoAuth {
	var netatmoAuth = new(NetatmoAuth)

	payload := url.Values{}
	payload.Set("grant_type", "password")
	payload.Set("scope", "read_station")
	payload.Set("client_id", os.Getenv("CLIENT_ID"))
	payload.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	payload.Set("username", os.Getenv("NETATMO_USERNAME"))
	payload.Set("password", os.Getenv("NETATMO_PASSWORD"))

	req, _ := http.NewRequest("POST", APIUrl+"/oauth2/token", strings.NewReader(payload.Encode()))
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &netatmoAuth); err != nil {
		log.Println("Error unmarshalling: " + err.Error())
	}
	return netatmoAuth
}
