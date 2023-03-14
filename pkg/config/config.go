package config

import (
	"Leviosa/pkg/log"
	"encoding/json"
	"os"
)

type config struct {
	ProxyUrl string `json:"proxyUrl"`
}

const configPath = "config.json"

var Config *config

func init() {
	Config = &config{}
	_, err := os.Stat(configPath)
	if err != nil {
		log.Logger.Warning("No config file detected! Creating default config!")
		bts, err := json.MarshalIndent(Config, "", "	")
		if err != nil {
			log.Logger.Error(err.Error())
		}
		err = os.WriteFile(configPath, bts, 0644)
		if err != nil {
			log.Logger.Error(err.Error())
		}
		return
	}
	f, err := os.ReadFile(configPath)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	err = json.Unmarshal(f, &Config)
	if err != nil {
		log.Logger.Error(err.Error())
	}
}
