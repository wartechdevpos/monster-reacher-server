package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/oauth2"
)

const (
	SERVICE_MAME_GOOGLE   = "google"
	SERVICE_MAME_FACEBOOK = "facebook"
	SERVICE_MAME_TWITTER  = "twitter"
	SERVICE_MAME_APPLE    = "apple"
)

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type oAuth2 struct {
	oConfig      *oauth2.Config
	userUrl      string
	opts         []oauth2.AuthCodeOption
	userInfo     *UserInfo
	providerName string
}

func (authPtr *oAuth2) SubmitAuth(tokenCode string) error {
	if authPtr.oConfig == nil {
		return errors.New("please call init and set oConfig and userUrl")
	}
	tok, err := authPtr.oConfig.Exchange(context.TODO(), tokenCode, authPtr.opts...)

	if err != nil {
		return err
	}
	client := authPtr.oConfig.Client(context.TODO(), tok)

	res, err := client.Get(authPtr.userUrl)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	authPtr.userInfo = &UserInfo{}
	json.Unmarshal(bodyBytes, authPtr.userInfo)

	// for twiiter
	if authPtr.userInfo.ID == "" {
		type DatauserInfo struct {
			Data UserInfo
		}
		datauserInfo := &DatauserInfo{}
		json.Unmarshal(bodyBytes, datauserInfo)
		authPtr.userInfo = &datauserInfo.Data
	}

	return nil
}
func (authPtr *oAuth2) GetData() *UserInfo     { return authPtr.userInfo }
func (authPtr *oAuth2) GetServiceName() string { return authPtr.providerName }

func initOAuth2(oConfig *oauth2.Config, userUrl string, providerName string, opts ...oauth2.AuthCodeOption) *oAuth2 {
	return &oAuth2{
		oConfig:      oConfig,
		userUrl:      userUrl,
		opts:         opts,
		userInfo:     nil,
		providerName: providerName,
	}
}

func NewOAuth2Provider(provider string) *oAuth2 {
	switch strings.ToLower(provider) {
	case SERVICE_MAME_GOOGLE:
		return initOAuth2(getOAut2GoogleConfig())
	case SERVICE_MAME_FACEBOOK:
		return initOAuth2(getOAut2FacebookConfig())
	case SERVICE_MAME_TWITTER:
		return initOAuth2(getOAut2TwitterConfig())
	case SERVICE_MAME_APPLE:
		return initOAuth2(getOAut2AppleConfig())
	}

	return nil
}
