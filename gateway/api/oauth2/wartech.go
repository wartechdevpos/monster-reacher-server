package oauth2

import (
	"fmt"

	"golang.org/x/oauth2"
	"wartech-studio.com/monster-reacher/libraries/config"
)

const SERVICES_NAME = "wartech"

var listenHost = fmt.Sprintf("%s:%d",
	config.WartechConfig().Services[SERVICES_NAME].Hosts[0],
	config.WartechConfig().Services[SERVICES_NAME].Ports[0])

var (
	ENDPOINT_WARTECH_AUTH  = listenHost + "/authorize"
	ENDPOINT_WARTECH_TOKEN = listenHost + "/token"
)

func getOAut2WartechConfig() (*oauth2.Config, string, string) {
	return &oauth2.Config{
		ClientID:     "85462020023651",
		ClientSecret: "ac29c66a3bb016d2c632a3a7dc5130b",
		RedirectURL:  "https://insuanhouse.ddns.net:3000/auth/wartech",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  ENDPOINT_WARTECH_AUTH,
			TokenURL: ENDPOINT_WARTECH_TOKEN,
		},
	}, listenHost + "/user", SERVICE_NAME_WARTECH
}
