package oauth2

import (
	"golang.org/x/oauth2"
)

const (
	ENDPOINT_GOOGLE_AUTH  = "https://accounts.google.com/o/oauth2/auth"
	ENDPOINT_GOOGLE_TOKEN = "https://oauth2.googleapis.com/token"
)

func getOAut2GoogleConfig() (*oauth2.Config, string, string) {
	return &oauth2.Config{
		ClientID:     "961959484291-j9lcuricbqtjn1btdhhgf3gn76o3a0n9.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "https://insuanhouse.ddns.net:3000/auth/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  ENDPOINT_GOOGLE_AUTH,
			TokenURL: ENDPOINT_GOOGLE_TOKEN,
		},
	}, "https://www.googleapis.com/oauth2/v1/userinfo", SERVICE_MAME_GOOGLE
}
