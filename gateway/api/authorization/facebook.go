package authorization

import (
	"encoding/json"
	"net/http"
)

type authorizationFacebook struct {
	method
	token    string
	userInfo *UserInfo
}

func NewAuthorizationFacebook(token string) Authorization {
	return &authorizationFacebook{
		token: token,
	}
}

func (auth *authorizationFacebook) SubmitAuth() error {
	auth.userInfo = &UserInfo{}
	req, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email&access_token="+auth.token, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(auth.userInfo)
	defer res.Body.Close()

	if err != nil {
		return err
	}
	return nil
}

func (auth *authorizationFacebook) GetData() *UserInfo     { return auth.userInfo }
func (auth *authorizationFacebook) GetServiceName() string { return SERVICE_MAME_FACEBOOK }
