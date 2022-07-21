package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"golang.org/x/oauth2"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
)

const (
	SERVICE_MAME_GOOGLE   = "google"
	SERVICE_MAME_FACEBOOK = "facebook"
	SERVICE_MAME_TWITTER  = "twitter"
	SERVICE_MAME_APPLE    = "apple"
	SERVICE_NAME_WARTECH  = "wartech"
)

type oAuth2 struct {
	oConfig      *oauth2.Config
	userUrl      string
	opts         []oauth2.AuthCodeOption
	userInfo     *data_schema.AuthenticationData
	providerName string
}

func (authPtr *oAuth2) SubmitAuth(ctx context.Context, tokenCode string) error {
	if authPtr.oConfig == nil {
		return errors.New("please call init and set oConfig and userUrl")
	}
	tok, err := authPtr.oConfig.Exchange(ctx, tokenCode, authPtr.opts...)

	if err != nil {
		return err
	}
	client := authPtr.oConfig.Client(ctx, tok)

	res, err := client.Get(authPtr.userUrl)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	authPtr.userInfo = &data_schema.AuthenticationData{}
	json.Unmarshal(bodyBytes, authPtr.userInfo)

	// for twiiter
	if authPtr.userInfo.GetId() == "" {
		type DatauserInfo struct {
			Data data_schema.AuthenticationData
		}
		datauserInfo := &DatauserInfo{}
		json.Unmarshal(bodyBytes, datauserInfo)
		authPtr.userInfo = &datauserInfo.Data
	}

	return nil
}
func (authPtr *oAuth2) GetData() *data_schema.AuthenticationData { return authPtr.userInfo }
func (authPtr *oAuth2) GetServiceName() string                   { return authPtr.providerName }

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
	oauth2Config := config.GetOAuth2Config()[provider]
	if oauth2Config == nil {
		return nil
	}
	oConfig := &oauth2.Config{
		ClientID:     oauth2Config.ClientId,
		ClientSecret: oauth2Config.ClientSecret,
		RedirectURL:  oauth2Config.RedirectUrl,
		Scopes:       oauth2Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  oauth2Config.Endpoint.AuthUrl,
			TokenURL: oauth2Config.Endpoint.TokenUrl,
		},
	}
	opts := []oauth2.AuthCodeOption{}

	for k, v := range oauth2Config.Params {
		opts = append(opts, oauth2.SetAuthURLParam(k, v))
	}
	return initOAuth2(oConfig, oauth2Config.Endpoint.UserUrl, provider, opts...)
}
