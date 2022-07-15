package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const service_config_name = "service"

type serviceConfig struct {
	Services  map[string]serviceConfigServices `json:"services"`
	Databases map[string]serviceConfigDatabase `json:"databases"`
}

type serviceConfigServices struct {
	Hosts []string `json:"hosts"`
	Ports []int    `json:"ports"`
}

type serviceConfigDatabase struct {
	Host string `json:"host"`
}

var cacheServiceConfig *serviceConfig = initServiceConfig()

func initServiceConfig() *serviceConfig {
	cfg := &serviceConfig{}
	jsonFile, err := os.Open(GetConfigFilePath(service_config_name, false))
	if err != nil {
		jsonFile, err = os.Open(GetConfigFilePath(service_config_name, true))
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

func GetServiceConfig() *serviceConfig {
	if cacheServiceConfig == nil {
		panic("not fount file on " + GetConfigFilePath(service_config_name, true))
	}
	return cacheServiceConfig
}
