package oauth2

import (
	"golang.org/x/oauth2"
)

const (
	ENDPOINT_FACEBOOK_AUTH  = "https://www.facebook.com/dialog/oauth"
	ENDPOINT_FACEBOOK_TOKEN = "https://graph.facebook.com/oauth/access_token"
)

func getOAut2FacebookConfig() (*oauth2.Config, string, string) {
	return &oauth2.Config{
		ClientID:     "415004290679134",
		ClientSecret: "ac29c62b9bb016d2c7703a3a7dc5130b",
		RedirectURL:  "https://insuanhouse.ddns.net:3000/auth/facebook",
		Scopes:       []string{"public_profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  ENDPOINT_FACEBOOK_AUTH,
			TokenURL: ENDPOINT_FACEBOOK_TOKEN,
		},
	}, "https://graph.facebook.com/v14.0/me", SERVICE_MAME_FACEBOOK
}
