package authorization

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

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
	b := []byte(os.Getenv("GOOGLE_CREDENTAILS_JSON"))

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
	log.Println(49)
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		return err
	}
	log.Println(54)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(59)
	auth.userInfo = &UserInfo{}
	log.Println(61, string(bodyBytes))
	err = json.Unmarshal(bodyBytes, auth.userInfo)
	return err
}

func (auth *authorizationGoogle) GetData() *UserInfo     { return auth.userInfo }
func (auth *authorizationGoogle) GetServiceName() string { return SERVICE_MAME_GOOGLE }
