package authorization

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2/google"
)

type authorizationGoogle struct {
	method
	token    string
	userInfo *UserInfo
}

func NewAuthorizationGoogle(token string) Authorization {
	return &authorizationGoogle{
		token: token,
	}
}

func ToGoogleUserInfo(data interface{}) *UserInfo {
	if s, ok := data.(*UserInfo); ok {
		return s
	}
	return nil
}

func (auth *authorizationGoogle) SubmitAuth() error {
	b := []byte(`{"web":{"client_id":"961959484291-j9lcuricbqtjn1btdhhgf3gn76o3a0n9.apps.googleusercontent.com","project_id":"monster-reacher","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["http://insuanhouse.ddns.net:3000/auth/google"]}}`)

	config, err := google.ConfigFromJSON(b,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile")
	if err != nil {
		return err
	}
	tok, err := config.Exchange(context.TODO(), auth.token)

	if err != nil {
		return err
	}
	client := config.Client(context.TODO(), tok)

	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	auth.userInfo = &UserInfo{}
	err = json.Unmarshal(bodyBytes, auth.userInfo)
	return err
}

func (auth *authorizationGoogle) GetData() *UserInfo     { return auth.userInfo }
func (auth *authorizationGoogle) GetServiceName() string { return SERVICE_MAME_GOOGLE }
