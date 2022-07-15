package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const name_config_name = "name"

type nameConfig struct {
	ServiceName struct {
		Gateway           string `json:"gateway"`
		ServicesDiscovery string `json:"services_discovery"`
	} `json:"service_name"`
	MicroServiceName struct {
		Wartech        string `json:"wartech"`
		Authentication string `json:"authentication"`
		Profile        string `json:"profile"`
		Character      string `json:"character"`
	} `json:"microservice_name"`
}

var cacheNameConfig *nameConfig = initNameConfig()

func GetNameConfig() *nameConfig {
	if cacheNameConfig == nil {
		panic("not fount file on " + GetConfigFilePath(name_config_name, true))
	}
	return cacheNameConfig
}

func initNameConfig() *nameConfig {
	cfg := nameConfig{}
	jsonFile, err := os.Open(GetConfigFilePath(name_config_name, false))
	if err != nil {
		jsonFile, err = os.Open(GetConfigFilePath(name_config_name, true))
		if err != nil {
			panic(err)
		}
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal([]byte(byteValue), &cfg); err != nil {
		return nil
	}
	return &cfg
}
