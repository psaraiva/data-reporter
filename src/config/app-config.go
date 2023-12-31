package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Name     string
	Version  string
	Port     string
	Host     string
	DBengine string
}

func (c *AppConfig) Load() (AppConfig, error) {
	config := AppConfig{}
	data, err := os.ReadFile("./config/app.json")
	if err != nil {
		return config, err
	}

	if json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *AppConfig) LoadDbEngine() (string, error) {
	appConfig, err := c.Load()
	if err != nil {
		return "unknow", err
	}

	return appConfig.DBengine, nil
}
