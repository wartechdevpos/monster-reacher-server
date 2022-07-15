package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const oauth2_config_name = "oauth2"

type oauth2ConfigInfo struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectUrl  string   `json:"redirect_url"`
	Scopes       []string `json:"scopes"`
	Endpoint     struct {
		DomainUrl string `json:"domain_url"`
		AuthUrl   string `json:"auth_url"`
		TokenUrl  string `json:"token_url"`
		UserUrl   string `json:"user_url"`
	} `json:"endpoint"`
	Params map[string]string `json:"params"`
}

var cacheOAuth2Config map[string]*oauth2ConfigInfo = initOAuth2Config()

func GetOAuth2Config() map[string]*oauth2ConfigInfo {
	if cacheOAuth2Config == nil {
		panic("not fount file on " + GetConfigFilePath(oauth2_config_name, true))
	}
	return cacheOAuth2Config
}

func initOAuth2Config() map[string]*oauth2ConfigInfo {
	cfg := make(map[string]*oauth2ConfigInfo)
	jsonFile, err := os.Open(GetConfigFilePath(oauth2_config_name, false))
	if err != nil {
		jsonFile, err = os.Open(GetConfigFilePath(oauth2_config_name, true))
		if err != nil {
			panic(err)
		}
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal([]byte(byteValue), &cfg); err != nil {
		return nil
	}
	return cfg
}
