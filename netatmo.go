package main

import (
	"log"
	"os"

	"github.com/go-resty/resty"
)

type NetatmoAuth struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	ExpiresIn    int      `json:"expires_in"`
	ExpireIn     int      `json:"expire_in"`
}

func authenticate(client *resty.Client) string {
	client.SetFormData(map[string]string{
		"grant_type":       "password",
		"scope":            "read_station",
		"client_id":        os.Getenv("CLIENT_ID"),
		"client_secret":    os.Getenv("CLIENT_SECRET"),
		"netatmo_username": os.Getenv("NETATMO_USERNAME"),
		"netatmo_password": os.Getenv("NETATMO_PASSWORD")})

	response, err := client.R().SetResult(NetatmoAuth{}).Post("/oauth2/token")
	if err != nil {
		log.Fatal("Error retrieven access tokens: " + err.Error())
	}

	return response.Result().(*NetatmoAuth).AccessToken
}
