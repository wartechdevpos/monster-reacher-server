package oauth2

import (
	"golang.org/x/oauth2"
)

const (
	ENDPOINT_TWITTER_AUTH  = "https://api.twitter.com/oauth/authorize"
	ENDPOINT_TWITTER_TOKEN = "https://api.twitter.com/2/oauth2/token"
)

func getOAut2TwitterConfig() (*oauth2.Config, string, string, oauth2.AuthCodeOption) {
	return &oauth2.Config{
		ClientID:     "ckk3c1pxNVo5dGNnU0xhZ0Q1YWI6MTpjaQ",
		ClientSecret: "vEuPsuxp7EJWrv8jCNI0gWRboA5hzk_WbyH5RiK01ZsV4CQMDJ",
		RedirectURL:  "https://insuanhouse.ddns.net:3000/auth/twitter",
		Scopes:       []string{"users.read", "tweet.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  ENDPOINT_TWITTER_AUTH,
			TokenURL: ENDPOINT_TWITTER_TOKEN,
		},
	}, "https://api.twitter.com/2/users/me", SERVICE_MAME_TWITTER, oauth2.SetAuthURLParam("code_verifier", "challenge")
}
