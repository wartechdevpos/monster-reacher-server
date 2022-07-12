package oauth2

import (
	"golang.org/x/oauth2"
)

const SERVICES_NAME = "wartech"

const (
	ENDPOINT_WARTECH_AUTH  = "http://127.0.0.1.net:20521/authorize"
	ENDPOINT_WARTECH_TOKEN = "http://127.0.0.1.net:20521/token"
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
	}, "http://127.0.0.1.net:20521/user", SERVICE_NAME_WARTECH
}
