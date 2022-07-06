package authorization

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ENDPOINT_TWITTER_TOKEN = "https://api.twitter.com/oauth/access_token"
)

type authorizationTwitter struct {
	method
	token    string
	userInfo *UserInfo
}

// NewAuthorizationTwitter - toekn is oauth_token--oauth_verifier
func NewAuthorizationTwitter(token string) Authorization {
	return &authorizationTwitter{
		token: token,
	}
}

func ToTwitterUserInfo(data interface{}) *UserInfo {
	if s, ok := data.(*UserInfo); ok {
		return s
	}
	return nil
}

func (auth *authorizationTwitter) SubmitAuth() error {
	tokens := strings.Split(auth.token, "--")
	if len(tokens) < 1 {
		return errors.New("token not support")
	}
	auth.userInfo = &UserInfo{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?oauth_token=%s&oauth_verifier=%s", ENDPOINT_TWITTER_TOKEN, tokens[0], tokens[1]), nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	bodyStr := string(bodyBytes)

	if res.StatusCode != http.StatusOK {
		return errors.New(bodyStr)
	}
	if len(strings.Split(bodyStr, "&")) < 2 {
		return errors.New("response format not support")
	}
	auth.userInfo.ID = strings.Split(strings.Split(bodyStr, "&")[1], "=")[1]
	auth.userInfo.Name = strings.Split(strings.Split(bodyStr, "&")[2], "=")[2]

	return nil
}

func (auth *authorizationTwitter) GetData() *UserInfo     { return auth.userInfo }
func (auth *authorizationTwitter) GetServiceName() string { return SERVICE_MAME_TWITTER }
