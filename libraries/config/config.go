package config

import "os"

const (
	docker_PATH_PROJECT = "/app/monster-reacher-server"
	local_PATH_PROJECT  = "D:/WTProject/devpos/monster-reacher-server"
)

var ENV = "docker"

func GetConfigFilePath(cfgName string, forceDefault bool) string {
	if os.Getenv("WARTECH_SERVER_ENV") != "" {
		ENV = os.Getenv("WARTECH_SERVER_ENV")
	}
	if ENV == "docker" {
		return docker_PATH_PROJECT + "/config/" + cfgName + ".json"
	}

	if forceDefault {
		return local_PATH_PROJECT + "/config/" + cfgName + ".json"
	}
	return local_PATH_PROJECT + "/config/" + cfgName + "_local.json"
}
